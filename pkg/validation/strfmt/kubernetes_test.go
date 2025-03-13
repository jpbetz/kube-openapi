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

package strfmt

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDNS1035Label(t *testing.T) {
	validLabels := []string{
		"a",
		"a-b",
		"a-b-c",
		"abc123",
		"abc-123",
	}

	invalidLabels := []string{
		"",                      // empty
		"-abc",                  // starts with hyphen
		"abc-",                  // ends with hyphen
		"ABC",                   // uppercase
		"a.b",                   // contains dot
		"a_b",                   // contains underscore
		"123",                   // starts with number
		strings.Repeat("a", 64), // too long
	}

	for _, label := range validLabels {
		t.Run(label, func(t *testing.T) {
			assert.True(t, isDNS1035Label(label), "Expected %q to be a valid DNS1035 label", label)
		})
	}

	for _, label := range invalidLabels {
		t.Run(label, func(t *testing.T) {
			assert.False(t, isDNS1035Label(label), "Expected %q to be an invalid DNS1035 label", label)
		})
	}
}

func TestDNS1035LabelPrefix(t *testing.T) {
	// TODO: Add test cases for DNS1035LabelPrefix validation
	// Test cases should include:
	// - Valid prefixes
	// - Invalid prefixes
	// - Edge cases
}

func TestDNS1123Label(t *testing.T) {
	validLabels := []string{
		"a",
		"a-b",
		"a-b-c",
		"abc123",
		"abc-123",
		"123-abc", // Unlike DNS1035, can start with number
	}

	invalidLabels := []string{
		"",                      // empty
		"-abc",                  // starts with hyphen
		"abc-",                  // ends with hyphen
		"ABC",                   // uppercase
		"a.b",                   // contains dot
		"a_b",                   // contains underscore
		strings.Repeat("a", 64), // too long
	}

	for _, label := range validLabels {
		t.Run(label, func(t *testing.T) {
			assert.True(t, isDNS1123Label(label), "Expected %q to be a valid DNS1123 label", label)
		})
	}

	for _, label := range invalidLabels {
		t.Run(label, func(t *testing.T) {
			assert.False(t, isDNS1123Label(label), "Expected %q to be an invalid DNS1123 label", label)
		})
	}
}

func TestDNS1123LabelPrefix(t *testing.T) {
	// TODO: Add test cases for DNS1123LabelPrefix validation
	// Test cases should include:
	// - Valid prefixes
	// - Invalid prefixes
	// - Edge cases
}

func TestDNS1123Subdomain(t *testing.T) {
	validSubdomains := []string{
		"a",
		"a.b",
		"a.b.c",
		"abc.123",
		"abc-123.def-456",
		"123.abc", // Unlike DNS1035, can start with number
	}

	invalidSubdomains := []string{
		"",                              // empty
		"-abc.def",                      // segment starts with hyphen
		"abc-.def",                      // segment ends with hyphen
		"ABC.def",                       // uppercase
		"a_b.def",                       // contains underscore
		strings.Repeat("a.", 127) + "a", // total length > 253
	}

	for _, subdomain := range validSubdomains {
		t.Run(subdomain, func(t *testing.T) {
			assert.True(t, isDNS1123Subdomain(subdomain), "Expected %q to be a valid DNS1123 subdomain", subdomain)
		})
	}

	for _, subdomain := range invalidSubdomains {
		t.Run(subdomain, func(t *testing.T) {
			assert.False(t, isDNS1123Subdomain(subdomain), "Expected %q to be an invalid DNS1123 subdomain", subdomain)
		})
	}
}

func TestDNS1123SubdomainPrefix(t *testing.T) {
	// TODO: Add test cases for DNS1123SubdomainPrefix validation
	// Test cases should include:
	// - Valid prefixes
	// - Invalid prefixes
	// - Edge cases
}

func TestQualifiedName(t *testing.T) {
	validNames := []string{
		"example.com/name",
		"example.com/name_with_underscore",
		"example.com/name-with-dash",
	}

	invalidNames := []string{
		"",                                 // empty
		"/name",                            // no namespace
		"example.com/",                     // no name
		"example.com//",                    // empty segment
		"-example.com/name",                // invalid DNS subdomain
		"example.com/-name",                // segment starts with dash
		"example.com/name-",                // segment ends with dash
		strings.Repeat("a", 254) + "/name", // too long
	}

	for _, name := range validNames {
		t.Run(name, func(t *testing.T) {
			assert.True(t, isQualifiedName(name), "Expected %q to be a valid qualified name", name)
		})
	}

	for _, name := range invalidNames {
		t.Run(name, func(t *testing.T) {
			assert.False(t, isQualifiedName(name), "Expected %q to be an invalid qualified name", name)
		})
	}
}

func TestQuantity(t *testing.T) {
	validQuantities := []string{
		"0",
		"100m",
		"1",
		"1.5",
		"1Gi",
		"100M",
		"0.01M",
		"100Mi",
		"100Ki",
		"100m",
		"100n",
		"-1",
		"-100m",
	}

	invalidQuantities := []string{
		"",        // empty
		"invalid", // not a number
		"1.1.1",   // invalid number
		"1.0.0Gi", // invalid number
	}

	for _, q := range validQuantities {
		t.Run(q, func(t *testing.T) {
			assert.True(t, isQuantity(q), "Expected %q to be a valid quantity", q)
		})
	}

	for _, q := range invalidQuantities {
		t.Run(q, func(t *testing.T) {
			assert.False(t, isQuantity(q), "Expected %q to be an invalid quantity", q)
		})
	}
}

func TestIP(t *testing.T) {
	validIPs := []string{
		"192.168.1.1",
		"10.0.0.0",
		"172.16.0.1",
		"2001:db8::",
		"fe80::1",
		"::1",
	}

	invalidIPs := []string{
		"",             // empty
		"256.1.2.3",    // invalid IPv4
		"1.2.3",        // incomplete IPv4
		"2001:xyz::",   // invalid IPv6
		"2001::db8::1", // invalid IPv6 (multiple ::)
		"not-an-ip",    // not an IP at all
	}

	for _, ip := range validIPs {
		t.Run(ip, func(t *testing.T) {
			assert.True(t, isIP(ip), "Expected %q to be a valid IP", ip)
		})
	}

	for _, ip := range invalidIPs {
		t.Run(ip, func(t *testing.T) {
			assert.False(t, isIP(ip), "Expected %q to be an invalid IP", ip)
		})
	}
}

func TestSemver(t *testing.T) {
	validSemvers := []string{
		"1.0.0",
		"2.3.4",
		"1.0.0-alpha",
		"1.0.0-alpha.1",
		"1.0.0-0.3.7",
		"1.0.0-x.7.z.92",
		"1.0.0-beta+exp.sha.5114f85",
		"1.0.0+20130313144700",
		"1.0.0-beta+exp.sha.5114f85",
		"11.200.300-alpha+meta",
	}

	invalidSemvers := []string{
		"",              // empty
		"1",             // missing minor and patch
		"1.0",           // missing patch
		"1.a.2",         // non-numeric version parts
		"1.0.0beta",     // prerelease without hyphen
		"v1.0.0",        // with v prefix
		"1.0.0-",        // empty prerelease
		"1.0.0+",        // empty build metadata
		"1.0.0-+",       // empty prerelease and build metadata
		"1.0.0-alpha_1", // invalid character in prerelease
		"1.0.0+alpha_1", // invalid character in build metadata
		"-1.0.0",        // negative major version
		"1.-2.0",        // negative minor version
		"1.0.-3",        // negative patch version
	}

	for _, v := range validSemvers {
		t.Run(v, func(t *testing.T) {
			assert.True(t, isSemver(v), "Expected %q to be a valid semver", v)
		})
	}

	for _, v := range invalidSemvers {
		t.Run(v, func(t *testing.T) {
			assert.False(t, isSemver(v), "Expected %q to be an invalid semver", v)
		})
	}
}

func TestSemverJSON(t *testing.T) {
	semver := Semver("1.0.0-alpha+001")

	// Test marshaling
	data, err := json.Marshal(semver)
	assert.NoError(t, err)
	assert.Equal(t, `"1.0.0-alpha+001"`, string(data))

	// Test unmarshaling
	var s Semver
	err = json.Unmarshal(data, &s)
	assert.NoError(t, err)
	assert.Equal(t, semver, s)
}

// Test JSON marshaling/unmarshaling
func TestKubernetesFormatJSON(t *testing.T) {
	t.Run("DNS1035Label", func(t *testing.T) {
		label := DNS1035Label("valid-label")
		data, err := json.Marshal(label)
		assert.NoError(t, err)
		assert.Equal(t, `"valid-label"`, string(data))

		var decoded DNS1035Label
		err = json.Unmarshal(data, &decoded)
		assert.NoError(t, err)
		assert.Equal(t, label, decoded)
	})

	t.Run("DNS1035LabelPrefix", func(t *testing.T) {
		// TODO: Add JSON marshal/unmarshal tests
	})

	t.Run("DNS1123Label", func(t *testing.T) {
		label := DNS1123Label("valid-label")
		data, err := json.Marshal(label)
		assert.NoError(t, err)
		assert.Equal(t, `"valid-label"`, string(data))

		var decoded DNS1123Label
		err = json.Unmarshal(data, &decoded)
		assert.NoError(t, err)
		assert.Equal(t, label, decoded)
	})

	t.Run("DNS1123LabelPrefix", func(t *testing.T) {
		// TODO: Add JSON marshal/unmarshal tests
	})

	t.Run("DNS1123Subdomain", func(t *testing.T) {
		subdomain := DNS1123Subdomain("valid.subdomain")
		data, err := json.Marshal(subdomain)
		assert.NoError(t, err)
		assert.Equal(t, `"valid.subdomain"`, string(data))

		var decoded DNS1123Subdomain
		err = json.Unmarshal(data, &decoded)
		assert.NoError(t, err)
		assert.Equal(t, subdomain, decoded)
	})

	t.Run("DNS1123SubdomainPrefix", func(t *testing.T) {
		// TODO: Add JSON marshal/unmarshal tests
	})

	t.Run("QualifiedName", func(t *testing.T) {
		name := QualifiedName("example.com/name")
		data, err := json.Marshal(name)
		assert.NoError(t, err)
		assert.Equal(t, `"example.com/name"`, string(data))

		var decoded QualifiedName
		err = json.Unmarshal(data, &decoded)
		assert.NoError(t, err)
		assert.Equal(t, name, decoded)
	})

	t.Run("Quantity", func(t *testing.T) {
		q := Quantity("100Mi")
		data, err := json.Marshal(q)
		assert.NoError(t, err)
		assert.Equal(t, `"100Mi"`, string(data))

		var decoded Quantity
		err = json.Unmarshal(data, &decoded)
		assert.NoError(t, err)
		assert.Equal(t, q, decoded)
	})

	t.Run("IP", func(t *testing.T) {
		ip := IP("192.168.1.1")
		data, err := json.Marshal(ip)
		assert.NoError(t, err)
		assert.Equal(t, `"192.168.1.1"`, string(data))

		var decoded IP
		err = json.Unmarshal(data, &decoded)
		assert.NoError(t, err)
		assert.Equal(t, ip, decoded)
	})
}

// Test string conversion
func TestKubernetesFormatString(t *testing.T) {
	t.Run("DNS1035Label", func(t *testing.T) {
		label := DNS1035Label("valid-label")
		assert.Equal(t, "valid-label", label.String())
	})

	t.Run("DNS1035LabelPrefix", func(t *testing.T) {
		// TODO: Add string conversion tests
	})

	t.Run("DNS1123Label", func(t *testing.T) {
		label := DNS1123Label("valid-label")
		assert.Equal(t, "valid-label", label.String())
	})

	t.Run("DNS1123LabelPrefix", func(t *testing.T) {
		// TODO: Add string conversion tests
	})

	t.Run("DNS1123Subdomain", func(t *testing.T) {
		subdomain := DNS1123Subdomain("valid.subdomain")
		assert.Equal(t, "valid.subdomain", subdomain.String())
	})

	t.Run("DNS1123SubdomainPrefix", func(t *testing.T) {
		// TODO: Add string conversion tests
	})

	t.Run("QualifiedName", func(t *testing.T) {
		name := QualifiedName("example.com/name")
		assert.Equal(t, "example.com/name", name.String())
	})

	t.Run("Quantity", func(t *testing.T) {
		q := Quantity("100Mi")
		assert.Equal(t, "100Mi", q.String())
	})

	t.Run("IP", func(t *testing.T) {
		ip := IP("192.168.1.1")
		assert.Equal(t, "192.168.1.1", ip.String())
	})
}

// Test prefix validation
func TestPrefixValidation(t *testing.T) {
	t.Run("DNS1035LabelPrefix", func(t *testing.T) {
		validPrefixes := []string{
			"a",
			"a-b",
			"valid",
		}

		invalidPrefixes := []string{
			"",     // empty
			"-abc", // starts with hyphen
			"ABC",  // uppercase
			"123",  // starts with number
		}

		for _, prefix := range validPrefixes {
			assert.True(t, isDNS1035LabelPrefix(prefix), "Expected %q to be a valid DNS1035 label prefix", prefix)
		}

		for _, prefix := range invalidPrefixes {
			assert.False(t, isDNS1035LabelPrefix(prefix), "Expected %q to be an invalid DNS1035 label prefix", prefix)
		}
	})

	t.Run("DNS1123LabelPrefix", func(t *testing.T) {
		validPrefixes := []string{
			"a",
			"a-b",
			"valid",
			"123", // can start with number
		}

		invalidPrefixes := []string{
			"",     // empty
			"-abc", // starts with hyphen
			"ABC",  // uppercase
		}

		for _, prefix := range validPrefixes {
			assert.True(t, isDNS1123LabelPrefix(prefix), "Expected %q to be a valid DNS1123 label prefix", prefix)
		}

		for _, prefix := range invalidPrefixes {
			assert.False(t, isDNS1123LabelPrefix(prefix), "Expected %q to be an invalid DNS1123 label prefix", prefix)
		}
	})

	t.Run("DNS1123SubdomainPrefix", func(t *testing.T) {
		validPrefixes := []string{
			"a",
			"a.b",
			"valid.prefix",
			"123.456", // can start with number
		}

		invalidPrefixes := []string{
			"",         // empty
			"-abc.def", // segment starts with hyphen
			"ABC.def",  // uppercase
		}

		for _, prefix := range validPrefixes {
			assert.True(t, isDNS1123SubdomainPrefix(prefix), "Expected %q to be a valid DNS1123 subdomain prefix", prefix)
		}

		for _, prefix := range invalidPrefixes {
			assert.False(t, isDNS1123SubdomainPrefix(prefix), "Expected %q to be an invalid DNS1123 subdomain prefix", prefix)
		}
	})
}
