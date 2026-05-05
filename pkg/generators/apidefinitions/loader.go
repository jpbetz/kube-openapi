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

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"sigs.k8s.io/yaml"
)

const (
	// We define an apiVersion and kinds for defining APIs
	// in the source tree like we do for everything else.
	schemeGroupVersion = "apidefinitions.config.k8s.io/v1alpha1"
	kindAPIGroup       = "APIGroup"

	// We have a naming convention for the files used to define
	// APIs in the source tree.
	apiGroupFile = "apigroup.yaml"
)

// LoadAPIGroup reads an apigroup.yaml file, returning nil if absent.
func LoadAPIGroup(dir string) (*APIGroup, error) {
	data, err := readManifest(dir, apiGroupFile)
	if err != nil || data == nil {
		return nil, err
	}
	g := &APIGroup{}
	if err := yaml.Unmarshal(data, g); err != nil {
		return nil, fmt.Errorf("%s: %w", filepath.Join(dir, apiGroupFile), err)
	}
	if err := validateTypeMeta(g.APIVersion, g.Kind, kindAPIGroup); err != nil {
		return nil, fmt.Errorf("%s: %w", filepath.Join(dir, apiGroupFile), err)
	}
	return g, nil
}

func readManifest(dir, filename string) ([]byte, error) {
	data, err := os.ReadFile(filepath.Join(dir, filename))
	if errors.Is(err, fs.ErrNotExist) {
		return nil, nil
	}
	return data, err
}

func validateTypeMeta(actualAPIVersion, actualKind, expectedKind string) error {
	if actualAPIVersion != schemeGroupVersion {
		return fmt.Errorf("expected apiVersion %s but got %s", schemeGroupVersion, actualAPIVersion)
	}
	if actualKind != expectedKind {
		return fmt.Errorf("expected kind %s but got %s", expectedKind, actualKind)
	}
	return nil
}
