// Copyright 2015 go-swagger maintainers
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package validate

import (
	"reflect"
	"regexp"
	"strings"

	"k8s.io/kube-openapi/pkg/validation/errors"
	"k8s.io/kube-openapi/pkg/validation/spec"
	"k8s.io/kube-openapi/pkg/validation/strfmt"
)

type objectValidator struct {
	In                   string
	MaxProperties        *int64
	MinProperties        *int64
	Required             []string
	Properties           map[string]spec.Schema
	AdditionalProperties *spec.SchemaOrBool
	PatternProperties    map[string]spec.Schema
	Root                 interface{}
	KnownFormats         strfmt.Registry

	propValidators            map[string]*SchemaValidator
	additionalPropsValidators *SchemaValidator
	patternPropValidators     map[string]*SchemaValidator
}

func newObjectValidator(in string, maxProperties, minProperties *int64, required []string, properties map[string]spec.Schema,
	additionalProperties *spec.SchemaOrBool, patternProperties map[string]spec.Schema, root interface{}, knownFormats strfmt.Registry, options ...Option) valueValidator {
	propValidators := make(map[string]*SchemaValidator, len(properties))
	for k, schema := range properties {
		sch := schema
		propValidators[k] = NewSchemaValidator(&sch, root, knownFormats, options...)
	}
	var additionalPropsValidator *SchemaValidator
	if additionalProperties != nil && additionalProperties.Schema != nil {
		additionalPropsValidator = NewSchemaValidator(additionalProperties.Schema, root, knownFormats, options...)
	}
	patternPropsValidators := make(map[string]*SchemaValidator, len(patternProperties))
	for k, schema := range patternProperties {
		sch := schema
		patternPropsValidators[k] = NewSchemaValidator(&sch, root, knownFormats, options...)
	}
	return &objectValidator{
		In:                        in,
		MaxProperties:             maxProperties,
		MinProperties:             minProperties,
		Required:                  required,
		Properties:                properties,
		AdditionalProperties:      additionalProperties,
		PatternProperties:         patternProperties,
		Root:                      root,
		KnownFormats:              knownFormats,
		propValidators:            propValidators,
		additionalPropsValidators: additionalPropsValidator,
		patternPropValidators:     patternPropsValidators,
	}
}

func (o *objectValidator) Applies(path string, source interface{}, kind reflect.Kind) bool {
	// TODO: this should also work for structs
	// there is a problem in the type validator where it will be unhappy about null values
	// so that requires more testing
	r := reflect.TypeOf(source) == specSchemaType && (kind == reflect.Map || kind == reflect.Struct)
	debugLog("object validator for %q applies %t for %T (kind: %v)\n", path, r, source, kind)
	return r
}

func (o *objectValidator) isProperties(path string) bool {
	p := strings.Split(path, ".")
	return len(p) > 1 && p[len(p)-1] == jsonProperties && p[len(p)-2] != jsonProperties
}

func (o *objectValidator) isDefault(path string) bool {
	p := strings.Split(path, ".")
	return len(p) > 1 && p[len(p)-1] == jsonDefault && p[len(p)-2] != jsonDefault
}

func (o *objectValidator) isExample(path string) bool {
	p := strings.Split(path, ".")
	return len(p) > 1 && (p[len(p)-1] == swaggerExample || p[len(p)-1] == swaggerExamples) && p[len(p)-2] != swaggerExample
}

func (o *objectValidator) Validate(path string, data interface{}) *Result {
	val := data.(map[string]interface{})
	// TODO: guard against nil data
	numKeys := int64(len(val))

	if o.MinProperties != nil && numKeys < *o.MinProperties {
		return errorHelp.sErr(errors.TooFewProperties(path, o.In, *o.MinProperties))
	}
	if o.MaxProperties != nil && numKeys > *o.MaxProperties {
		return errorHelp.sErr(errors.TooManyProperties(path, o.In, *o.MaxProperties))
	}

	res := new(Result)

	// check validity of field names
	if o.AdditionalProperties != nil && !o.AdditionalProperties.Allows {
		// Case: additionalProperties: false
		for k := range val {
			_, regularProperty := o.Properties[k]
			matched := false

			for pk := range o.PatternProperties {
				if matches, _ := regexp.MatchString(pk, k); matches {
					matched = true
					break
				}
			}

			if !regularProperty && !matched {
				// Special properties "$schema" and "id" are ignored
				res.AddErrors(errors.PropertyNotAllowed(path, o.In, k))
			}
		}
	} else {
		// Cases: no additionalProperties (implying: true), or additionalProperties: true, or additionalProperties: { <<schema>> }
		for key, value := range val {
			_, regularProperty := o.Properties[key]

			// Validates property against "patternProperties" if applicable
			// BUG(fredbi): succeededOnce is always false

			// NOTE: how about regular properties which do not match patternProperties?
			matched, succeededOnce, _ := o.validatePatternProperty(path, key, value, res)

			if !(regularProperty || matched || succeededOnce) {

				// Cases: properties which are not regular properties and have not been matched by the PatternProperties validator
				if o.additionalPropsValidators != nil {
					// AdditionalProperties as Schema
					res.Merge(o.additionalPropsValidators.Validate(path+"."+key, value))
				} else if regularProperty && !(matched || succeededOnce) {
					// TODO: this is dead code since regularProperty=false here
					res.AddErrors(errors.FailedAllPatternProperties(path, o.In, key))
				}
			}
		}
		// Valid cases: additionalProperties: true or undefined
	}

	createdFromDefaults := map[string]bool{}

	// Property types:
	// - regular Property
	for pName, validator := range o.propValidators {
		rName := pName
		if path != "" {
			rName = path + "." + pName
		}

		// Recursively validates each property against its schema
		if v, ok := val[pName]; ok {
			r := validator.Validate(rName, v)
			res.Merge(r)
		}
	}

	// Check required properties
	if len(o.Required) > 0 {
		for _, k := range o.Required {
			if _, ok := val[k]; !ok && !createdFromDefaults[k] {
				res.AddErrors(errors.Required(path+"."+k, o.In))
				continue
			}
		}
	}

	// Check patternProperties
	// TODO: it looks like we have done that twice in many cases
	for key, value := range val {
		_, regularProperty := o.Properties[key]
		matched, _ /*succeededOnce*/, patterns := o.validatePatternProperty(path, key, value, res)
		if !regularProperty && (matched /*|| succeededOnce*/) {
			for _, pName := range patterns {
				if validator, ok := o.patternPropValidators[pName]; ok {
					res.Merge(validator.Validate(path+"."+key, value))
				}
			}
		}
	}
	return res
}

// TODO: succeededOnce is not used anywhere
func (o *objectValidator) validatePatternProperty(path string, key string, value interface{}, result *Result) (bool, bool, []string) {
	matched := false
	succeededOnce := false
	var patterns []string

	for k, validator := range o.patternPropValidators {
		if match, _ := regexp.MatchString(k, key); match {
			patterns = append(patterns, k)
			matched = true
			res := validator.Validate(path+"."+key, value)
			result.Merge(res)
		}
	}

	// BUG(fredbi): can't get to here. Should remove dead code (commented out).

	//if succeededOnce {
	//	result.Inc()
	//}

	return matched, succeededOnce, patterns
}
