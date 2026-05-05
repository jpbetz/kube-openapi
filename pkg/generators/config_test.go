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

package generators

import (
	"os"
	"path/filepath"
	"testing"

	"k8s.io/gengo/v2/generator"
	"k8s.io/gengo/v2/types"
	"k8s.io/kube-openapi/cmd/openapi-gen/args"
)

func TestIsReadOnlyPkg(t *testing.T) {
	tests := []struct {
		name         string
		pkgPath      string
		readOnlyPkgs []string
		want         bool
	}{
		{
			name:         "nil readonly pkgs matches nothing",
			pkgPath:      "k8s.io/apimachinery/pkg/runtime",
			readOnlyPkgs: nil,
			want:         false,
		},
		{
			name:         "empty readonly pkgs matches nothing",
			pkgPath:      "k8s.io/apimachinery/pkg/runtime",
			readOnlyPkgs: []string{},
			want:         false,
		},
		{
			name:         "exact match",
			pkgPath:      "k8s.io/apimachinery/pkg/apis/meta/v1",
			readOnlyPkgs: []string{"k8s.io/apimachinery/pkg/apis/meta/v1"},
			want:         true,
		},
		{
			name:         "no prefix match",
			pkgPath:      "k8s.io/apimachinery/pkg/apis/meta/v1",
			readOnlyPkgs: []string{"k8s.io/apimachinery"},
			want:         false,
		},
		{
			name:         "non-readonly package",
			pkgPath:      "k8s.io/sample-apiserver/pkg/apis/wardle/v1beta1",
			readOnlyPkgs: []string{"k8s.io/apimachinery/pkg/apis/meta/v1"},
			want:         false,
		},
		{
			name:         "multiple readonly pkgs",
			pkgPath:      "k8s.io/apimachinery/pkg/runtime",
			readOnlyPkgs: []string{"k8s.io/apimachinery/pkg/apis/meta/v1", "k8s.io/apimachinery/pkg/runtime"},
			want:         true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isReadOnlyPkg(tt.pkgPath, tt.readOnlyPkgs)
			if got != tt.want {
				t.Errorf("isReadOnlyPkg(%q, %v) = %v, want %v", tt.pkgPath, tt.readOnlyPkgs, got, tt.want)
			}
		})
	}
}

// versionDirWithAPIGroup creates a temp dir <root>/<group>/<version> with
// an apigroup.yaml at <root>/<group> declaring the given modelPackage.
// Returns the version dir.
func versionDirWithAPIGroup(t *testing.T, group, modelPkg, version string) string {
	t.Helper()
	groupDir := filepath.Join(t.TempDir(), group)
	versionDir := filepath.Join(groupDir, version)
	if err := os.MkdirAll(versionDir, 0o755); err != nil {
		t.Fatalf("mkdir: %v", err)
	}
	manifest := "apiVersion: apidefinitions.config.k8s.io/v1alpha1\n" +
		"kind: APIGroup\n" +
		"metadata: {name: " + group + "}\n" +
		"spec:\n" +
		"  modelPackage: " + modelPkg + "\n" +
		"  versions: [{name: " + version + "}]\n"
	if err := os.WriteFile(filepath.Join(groupDir, "apigroup.yaml"), []byte(manifest), 0o644); err != nil {
		t.Fatalf("write: %v", err)
	}
	return versionDir
}

func TestGetModelNameTargets_ReadOnlyPkgs(t *testing.T) {
	localPkg := "k8s.io/sample-apiserver/pkg/apis/wardle/v1beta1"
	depPkg := "k8s.io/apimachinery/pkg/apis/meta/v1"

	// Build a minimal generator.Context with two input packages —
	// one "local", one "dependency" — each with a public struct.
	// resolvePackageModelPackage walks pkg.Dir's parent for an
	// apigroup.yaml, so we materialize one per package.
	universe := types.Universe{
		localPkg: {
			Path: localPkg,
			Dir:  versionDirWithAPIGroup(t, "wardle", "io.k8s.sample-apiserver.pkg.apis.wardle", "v1beta1"),
			Name: "v1beta1",
			Types: map[string]*types.Type{
				"Flunder": {
					Name: types.Name{Package: localPkg, Name: "Flunder"},
					Kind: types.Struct,
				},
			},
		},
		depPkg: {
			Path: depPkg,
			Dir:  versionDirWithAPIGroup(t, "meta", "io.k8s.apimachinery.pkg.apis.meta", "v1"),
			Name: "v1",
			Types: map[string]*types.Type{
				"ObjectMeta": {
					Name: types.Name{Package: depPkg, Name: "ObjectMeta"},
					Kind: types.Struct,
				},
			},
		},
	}

	ctx := &generator.Context{
		Inputs:   []string{localPkg, depPkg},
		Universe: universe,
	}

	tests := []struct {
		name           string
		readOnlyPkgs   []string
		wantTargetPkgs []string
	}{
		{
			name:           "no readonly pkgs generates for all packages",
			readOnlyPkgs:   nil,
			wantTargetPkgs: []string{localPkg, depPkg},
		},
		{
			name:           "readonly dep excludes dependency",
			readOnlyPkgs:   []string{depPkg},
			wantTargetPkgs: []string{localPkg},
		},
		{
			name:           "readonly local excludes local",
			readOnlyPkgs:   []string{localPkg},
			wantTargetPkgs: []string{depPkg},
		},
		{
			name:           "both readonly excludes all",
			readOnlyPkgs:   []string{localPkg, depPkg},
			wantTargetPkgs: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &args.Args{
				OutputModelNameFile: "zz_generated.model_name.go",
				ReadOnlyPkgs:        tt.readOnlyPkgs,
			}

			targets := GetModelNameTargets(ctx, a, nil)

			gotPkgs := make([]string, 0, len(targets))
			for _, tgt := range targets {
				st, ok := tgt.(*generator.SimpleTarget)
				if !ok {
					t.Fatalf("unexpected target type %T", tgt)
				}
				gotPkgs = append(gotPkgs, st.PkgPath)
			}

			if len(gotPkgs) != len(tt.wantTargetPkgs) {
				t.Fatalf("got %d targets %v, want %d targets %v", len(gotPkgs), gotPkgs, len(tt.wantTargetPkgs), tt.wantTargetPkgs)
			}

			for i, want := range tt.wantTargetPkgs {
				if gotPkgs[i] != want {
					t.Errorf("target[%d] = %q, want %q", i, gotPkgs[i], want)
				}
			}
		})
	}
}
