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
	"fmt"
	"reflect"

	"k8s.io/kube-openapi/pkg/validation/spec"
	"k8s.io/kube-openapi/pkg/validation/strfmt"
)

type schemaSliceValidator struct {
	In              string
	MaxItems        *int64
	MinItems        *int64
	UniqueItems     bool
	AdditionalItems *spec.SchemaOrBool
	Items           *spec.SchemaOrArray
	Root            interface{}
	KnownFormats    strfmt.Registry

	itemSchemaValidator      *SchemaValidator
	itemSchemasValidators    []*SchemaValidator
	additionalItemsValidator *SchemaValidator
}

func newSchemaSliceValidator(in string, maxItems, minItems *int64, uniqueItems bool, additionalItems *spec.SchemaOrBool,
	items *spec.SchemaOrArray, root interface{}, knownFormats strfmt.Registry, options ...Option) valueValidator {
	var itemSchemaValidator *SchemaValidator
	if items != nil && items.Schema != nil {
		itemSchemaValidator = NewSchemaValidator(items.Schema, root, knownFormats, options...)
	}
	var itemSchemasValidators []*SchemaValidator
	if items != nil && len(items.Schemas) > 0 {
		itemsSize := len(items.Schemas)
		itemSchemasValidators = make([]*SchemaValidator, itemsSize)
		for i := 0; i < itemsSize; i++ {
			itemSchemasValidators[i] = NewSchemaValidator(&items.Schemas[i], root, knownFormats, options...)
		}
	}
	var additionalItemsValidator *SchemaValidator
	if additionalItems != nil && additionalItems.Schema != nil {
		additionalItemsValidator = NewSchemaValidator(additionalItems.Schema, root, knownFormats, options...)
	}
	return &schemaSliceValidator{
		In:                       in,
		MaxItems:                 maxItems,
		MinItems:                 minItems,
		UniqueItems:              uniqueItems,
		AdditionalItems:          additionalItems,
		Items:                    items,
		Root:                     root,
		KnownFormats:             knownFormats,
		itemSchemaValidator:      itemSchemaValidator,
		itemSchemasValidators:    itemSchemasValidators,
		additionalItemsValidator: additionalItemsValidator,
	}
}

func (s *schemaSliceValidator) Applies(path string, source interface{}, kind reflect.Kind) bool {
	_, ok := source.(*spec.Schema)
	r := ok && kind == reflect.Slice
	return r
}

func (s *schemaSliceValidator) Validate(path string, data interface{}) *Result {
	result := new(Result)
	if data == nil {
		return result
	}
	val := reflect.ValueOf(data)
	size := val.Len()

	if s.itemSchemaValidator != nil {
		for i := 0; i < size; i++ {
			value := val.Index(i)
			result.Merge(s.itemSchemaValidator.Validate(fmt.Sprintf("%s.%d", path, i), value.Interface()))
		}
	}

	itemsSize := 0
	if  len(s.itemSchemasValidators) > 0 {
		itemsSize = len(s.Items.Schemas)
		for i := 0; i < itemsSize; i++ {
			validator := s.itemSchemasValidators[i]
			if val.Len() <= i {
				break
			}
			result.Merge(validator.Validate(fmt.Sprintf("%s.%d", path, i), val.Index(i).Interface()))
		}
	}
	if s.AdditionalItems != nil && itemsSize < size {
		if s.Items != nil && len(s.Items.Schemas) > 0 && !s.AdditionalItems.Allows {
			result.AddErrors(arrayDoesNotAllowAdditionalItemsMsg())
		}
		if s.additionalItemsValidator != nil {
			for i := itemsSize; i < size-itemsSize+1; i++ {
				result.Merge(s.additionalItemsValidator.Validate(fmt.Sprintf("%s.%d", path, i), val.Index(i).Interface()))
			}
		}
	}

	if s.MinItems != nil {
		if err := MinItems(path, s.In, int64(size), *s.MinItems); err != nil {
			result.AddErrors(err)
		}
	}
	if s.MaxItems != nil {
		if err := MaxItems(path, s.In, int64(size), *s.MaxItems); err != nil {
			result.AddErrors(err)
		}
	}
	if s.UniqueItems {
		if err := UniqueItems(path, s.In, val.Interface()); err != nil {
			result.AddErrors(err)
		}
	}
	result.Inc()
	return result
}
