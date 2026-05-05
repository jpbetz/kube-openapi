/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package apidefinitions

// Metadata is the kubernetes-style metadata block used by codegen manifests.
type Metadata struct {
	Name string `json:"name"`
}

// APIGroup declares an external versioned API group.
type APIGroup struct {
	APIVersion string       `json:"apiVersion"`
	Kind       string       `json:"kind"`
	Metadata   Metadata     `json:"metadata"`
	Spec       APIGroupSpec `json:"spec"`
}

// APIGroupSpec provides a specification of an APIGroup.
type APIGroupSpec struct {
	// ModelPackage is the OpenAPI model package prefix shared by every
	// version in this group. The per-version model package is
	// "<ModelPackage>.<version>" (e.g. "io.k8s.api.apps" + "v1" yields
	// "io.k8s.api.apps.v1").
	ModelPackage string `json:"modelPackage,omitempty"`

	// Versions enumerates the versions of the group that openapi-gen
	// should activate for. A version under the group's directory that
	// is not listed here (e.g. a deprecated version retained for
	// conversions only) is left alone.
	Versions []Version `json:"versions,omitempty"`
}

// Version describes a single version inside APIGroupSpec.Versions.
type Version struct {
	// Name is the version, e.g. "v1" or "v1beta2".
	Name string `json:"name"`
}

// ModelPackageFor returns the OpenAPI model package for the given version.
func (g *APIGroup) ModelPackageFor(version string) string {
	if g.Spec.ModelPackage == "" {
		return ""
	}
	return g.Spec.ModelPackage + "." + version
}

// HasVersion reports whether version is listed in spec.versions.
func (g *APIGroup) HasVersion(version string) bool {
	for _, v := range g.Spec.Versions {
		if v.Name == version {
			return true
		}
	}
	return false
}
