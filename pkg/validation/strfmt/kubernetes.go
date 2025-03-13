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

// Package strfmt contains custom string formats used in Kubernetes.
package strfmt

import (
	"encoding/json"
	"net/netip"
	"strings"

	"github.com/blang/semver/v4"

	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/apimachinery/pkg/util/validation"
)

// DNS1035Label represents a DNS label that conforms to RFC 1035.
// This format is required for Kubernetes Service names.
//
// Requirements:
// - Contain at most 63 characters
// - Contain only lowercase alphanumeric characters or '-'
// - Start with an alphabetic character
// - End with an alphanumeric character
//
// Examples:
// - Valid: "my-name", "abc-123"
// - Invalid: "123-abc" (starts with number), "my_name" (underscore), "MY-NAME" (uppercase)
type DNS1035Label string

// MarshalText turns this instance into text
func (d DNS1035Label) MarshalText() ([]byte, error) {
	return []byte(string(d)), nil
}

// UnmarshalText hydrates this instance from text
func (d *DNS1035Label) UnmarshalText(data []byte) error {
	*(d) = DNS1035Label(string(data))
	return nil
}

// String converts this value to a string
func (d DNS1035Label) String() string {
	return string(d)
}

// MarshalJSON returns the DNS1035Label as JSON
func (d DNS1035Label) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(d))
}

// UnmarshalJSON sets the DNS1035Label from JSON
func (d *DNS1035Label) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}
	return d.UnmarshalText([]byte(str))
}

// DNS1035LabelPrefix represents a DNS label prefix that conforms to RFC 1035.
// This is similar to DNS1035Label but allows a hyphen at the end.
//
// Requirements:
// - Contain at most 63 characters
// - Contain only lowercase alphanumeric characters or '-'
// - Start with an alphabetic character
//
// Examples:
// - Valid: "my-name-", "abc-"
// - Invalid: "123-" (starts with number), "my_name-" (underscore), "MY-NAME-" (uppercase)
type DNS1035LabelPrefix string

// MarshalText turns this instance into text
func (d DNS1035LabelPrefix) MarshalText() ([]byte, error) {
	return []byte(string(d)), nil
}

// UnmarshalText hydrates this instance from text
func (d *DNS1035LabelPrefix) UnmarshalText(data []byte) error {
	*(d) = DNS1035LabelPrefix(string(data))
	return nil
}

// String converts this value to a string
func (d DNS1035LabelPrefix) String() string {
	return string(d)
}

// MarshalJSON returns the DNS1035LabelPrefix as JSON
func (d DNS1035LabelPrefix) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(d))
}

// UnmarshalJSON sets the DNS1035LabelPrefix from JSON
func (d *DNS1035LabelPrefix) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}
	return d.UnmarshalText([]byte(str))
}

// DNS1123Label represents a DNS label that conforms to RFC 1123.
// This format is required for most Kubernetes resource names.
//
// Requirements:
// - Contain at most 63 characters
// - Contain only lowercase alphanumeric characters or '-'
// - Start with an alphanumeric character
// - End with an alphanumeric character
//
// Examples:
// - Valid: "my-name", "123-abc"
// - Invalid: "-abc" (starts with hyphen), "my_name" (underscore), "MY-NAME" (uppercase)
type DNS1123Label string

// MarshalText turns this instance into text
func (d DNS1123Label) MarshalText() ([]byte, error) {
	return []byte(string(d)), nil
}

// UnmarshalText hydrates this instance from text
func (d *DNS1123Label) UnmarshalText(data []byte) error {
	*(d) = DNS1123Label(string(data))
	return nil
}

// String converts this value to a string
func (d DNS1123Label) String() string {
	return string(d)
}

// MarshalJSON returns the DNS1123Label as JSON
func (d DNS1123Label) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(d))
}

// UnmarshalJSON sets the DNS1123Label from JSON
func (d *DNS1123Label) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}
	return d.UnmarshalText([]byte(str))
}

// DNS1123LabelPrefix represents a DNS label prefix that conforms to RFC 1123.
// This is similar to DNS1123Label but allows a hyphen at the end.
//
// Requirements:
// - Contain at most 63 characters
// - Contain only lowercase alphanumeric characters or '-'
// - Start with an alphanumeric character
//
// Examples:
// - Valid: "my-name-", "123-abc-"
// - Invalid: "-abc-" (starts with hyphen), "my_name-" (underscore), "MY-NAME-" (uppercase)
type DNS1123LabelPrefix string

// MarshalText turns this instance into text
func (d DNS1123LabelPrefix) MarshalText() ([]byte, error) {
	return []byte(string(d)), nil
}

// UnmarshalText hydrates this instance from text
func (d *DNS1123LabelPrefix) UnmarshalText(data []byte) error {
	*(d) = DNS1123LabelPrefix(string(data))
	return nil
}

// String converts this value to a string
func (d DNS1123LabelPrefix) String() string {
	return string(d)
}

// MarshalJSON returns the DNS1123LabelPrefix as JSON
func (d DNS1123LabelPrefix) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(d))
}

// UnmarshalJSON sets the DNS1123LabelPrefix from JSON
func (d *DNS1123LabelPrefix) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}
	return d.UnmarshalText([]byte(str))
}

// DNS1123Subdomain represents a DNS subdomain name that conforms to RFC 1123.
// This format is used for Kubernetes resource names that need to be referenced by DNS.
//
// Requirements:
// - Contain at most 253 characters
// - Contain only lowercase alphanumeric characters, '-' or '.'
// - Start with an alphanumeric character
// - End with an alphanumeric character
// - Each segment (label) must:
//   - Contain at most 63 characters
//   - Start with an alphanumeric character
//   - End with an alphanumeric character
//
// Examples:
// - Valid: "example.com", "my-name.namespace.svc"
// - Invalid: ".example" (starts with dot), "example." (ends with dot), "my_name.example" (underscore)
type DNS1123Subdomain string

// MarshalText turns this instance into text
func (d DNS1123Subdomain) MarshalText() ([]byte, error) {
	return []byte(string(d)), nil
}

// UnmarshalText hydrates this instance from text
func (d *DNS1123Subdomain) UnmarshalText(data []byte) error {
	*(d) = DNS1123Subdomain(string(data))
	return nil
}

// String converts this value to a string
func (d DNS1123Subdomain) String() string {
	return string(d)
}

// MarshalJSON returns the DNS1123Subdomain as JSON
func (d DNS1123Subdomain) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(d))
}

// UnmarshalJSON sets the DNS1123Subdomain from JSON
func (d *DNS1123Subdomain) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}
	return d.UnmarshalText([]byte(str))
}

// DNS1123SubdomainPrefix represents a DNS subdomain prefix that conforms to RFC 1123.
// This is similar to DNS1123Subdomain but allows a hyphen at the end of the last segment.
//
// Requirements:
// - Contain at most 253 characters
// - Contain only lowercase alphanumeric characters, '-' or '.'
// - Start with an alphanumeric character
// - Each segment (label) must:
//   - Contain at most 63 characters
//   - Start with an alphanumeric character
//   - End with an alphanumeric character (except last segment)
//   - Contain only lowercase alphanumeric characters or '-'
//
// Examples:
// - Valid: "example.com-", "my-name.namespace.svc-"
// - Invalid: ".example-" (starts with dot), "my_name.example-" (underscore)
type DNS1123SubdomainPrefix string

// MarshalText turns this instance into text
func (d DNS1123SubdomainPrefix) MarshalText() ([]byte, error) {
	return []byte(string(d)), nil
}

// UnmarshalText hydrates this instance from text
func (d *DNS1123SubdomainPrefix) UnmarshalText(data []byte) error {
	*(d) = DNS1123SubdomainPrefix(string(data))
	return nil
}

// String converts this value to a string
func (d DNS1123SubdomainPrefix) String() string {
	return string(d)
}

// MarshalJSON returns the DNS1123SubdomainPrefix as JSON
func (d DNS1123SubdomainPrefix) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(d))
}

// UnmarshalJSON sets the DNS1123SubdomainPrefix from JSON
func (d *DNS1123SubdomainPrefix) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}
	return d.UnmarshalText([]byte(str))
}

// QualifiedName represents a qualified name string.
// This format is used for Kubernetes resource names that include a namespace or domain.
//
// Requirements:
// - Must be a valid DNS1123 subdomain prefix
// - Must contain at least one '/' to separate namespace/domain from name
// - The part before the first '/' must be a valid DNS1123 subdomain
// - The part after the first '/' must be a valid path segment
//
// Examples:
// - Valid: "example.com/resource", "k8s.io/api/core/v1"
// - Invalid: "example" (no namespace), "/resource" (no namespace), "example.com/" (no name)
type QualifiedName string

// MarshalText turns this instance into text
func (q QualifiedName) MarshalText() ([]byte, error) {
	return []byte(string(q)), nil
}

// UnmarshalText hydrates this instance from text
func (q *QualifiedName) UnmarshalText(data []byte) error {
	*(q) = QualifiedName(string(data))
	return nil
}

// String converts this value to a string
func (q QualifiedName) String() string {
	return string(q)
}

// MarshalJSON returns the QualifiedName as JSON
func (q QualifiedName) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(q))
}

// UnmarshalJSON sets the QualifiedName from JSON
func (q *QualifiedName) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}
	return q.UnmarshalText([]byte(str))
}

// Quantity represents a Kubernetes resource quantity.
// This format is used to specify resource requirements and limits.
//
// The format is based on Kubernetes quantity strings:
// - A number (integer or decimal)
// - Optionally followed by one of these suffixes:
//   - E, P, T, G, M, k (powers of 1000)
//   - Ei, Pi, Ti, Gi, Mi, Ki (powers of 1024)
//   - m (milli), u or µ (micro), n (nano)
//
// Examples:
// - Valid: "100m", "1Gi", "500Mi", "1.5", "0.1n"
// - Invalid: "1G" (invalid suffix), "1.1.1" (invalid number)
type Quantity string

// MarshalText turns this instance into text
func (q Quantity) MarshalText() ([]byte, error) {
	return []byte(string(q)), nil
}

// UnmarshalText hydrates this instance from text
func (q *Quantity) UnmarshalText(data []byte) error {
	*(q) = Quantity(string(data))
	return nil
}

// String converts this value to a string
func (q Quantity) String() string {
	return string(q)
}

// MarshalJSON returns the Quantity as JSON
func (q Quantity) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(q))
}

// UnmarshalJSON sets the Quantity from JSON
func (q *Quantity) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}
	return q.UnmarshalText([]byte(str))
}

// IP represents an IP address (either IPv4 or IPv6).
//
// Requirements:
// - Must be a valid IPv4 or IPv6 address
//
// Examples:
// - Valid IPv4: "192.168.1.1", "10.0.0.0"
// - Valid IPv6: "2001:db8::", "fe80::1"
// - Invalid: "256.1.2.3", "1.2.3", "2001:xyz::"
type IP string

// MarshalText turns this instance into text
func (ip IP) MarshalText() ([]byte, error) {
	return []byte(string(ip)), nil
}

// UnmarshalText hydrates this instance from text
func (ip *IP) UnmarshalText(data []byte) error {
	*(ip) = IP(string(data))
	return nil
}

// String converts this value to a string
func (ip IP) String() string {
	return string(ip)
}

// MarshalJSON returns the IP as JSON
func (ip IP) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(ip))
}

// UnmarshalJSON sets the IP from JSON
func (ip *IP) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}
	return ip.UnmarshalText([]byte(str))
}

// Validation functions for Kubernetes formats
func isDNS1035Label(str string) bool {
	return len(validation.IsDNS1035Label(str)) == 0
}

func isDNS1035LabelPrefix(str string) bool {
	return isDNS1035Label(maskTrailingDash(str))
}

func isDNS1123Label(str string) bool {
	return len(validation.IsDNS1123Label(str)) == 0
}

func isDNS1123LabelPrefix(str string) bool {
	return isDNS1123Label(maskTrailingDash(str))
}

func isDNS1123Subdomain(str string) bool {
	return len(validation.IsDNS1123Subdomain(str)) == 0
}

func isDNS1123SubdomainPrefix(str string) bool {
	return isDNS1123Subdomain(maskTrailingDash(str))
}

func isQualifiedName(str string) bool {
	return len(validation.IsQualifiedName(str)) == 0
}

func isQuantity(str string) bool {
	_, err := resource.ParseQuantity(str)
	return err == nil
}

func isIP(str string) bool {
	addr, err := netip.ParseAddr(str)
	if err != nil {
		return false
	}

	if addr.Zone() != "" {
		return false
	}

	if addr.Is4In6() {
		return false
	}

	return true
}

// Semver represents a semantic version string that follows the semver.org specification.
//
// swagger:strfmt semver
type Semver string

// MarshalText turns this instance into text
func (s Semver) MarshalText() ([]byte, error) {
	return []byte(s), nil
}

// UnmarshalText hydrates this instance from text
func (s *Semver) UnmarshalText(data []byte) error {
	*(s) = Semver(data)
	return nil
}

// String converts this value to a string
func (s Semver) String() string {
	return string(s)
}

// MarshalJSON returns the Semver as JSON
func (s Semver) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(s))
}

// UnmarshalJSON sets the Semver from JSON
func (s *Semver) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}
	return s.UnmarshalText([]byte(str))
}

// DeepCopyInto copies the receiver into out. out must be non-nil.
func (s *Semver) DeepCopyInto(out *Semver) {
	*out = *s
}

// DeepCopy creates a deep copy of Semver
func (s *Semver) DeepCopy() *Semver {
	if s == nil {
		return nil
	}
	out := new(Semver)
	s.DeepCopyInto(out)
	return out
}

func isSemver(str string) bool {
	_, err := semver.Parse(str)
	return err == nil
}

func init() {
	// Register Kubernetes-specific formats
	dns1035Label := DNS1035Label("")
	Default.Add("dns1035-label", &dns1035Label, isDNS1035Label)

	dns1035LabelPrefix := DNS1035LabelPrefix("")
	Default.Add("dns1035-label-prefix", &dns1035LabelPrefix, isDNS1035LabelPrefix)

	dns1123Label := DNS1123Label("")
	Default.Add("dns1123-label", &dns1123Label, isDNS1123Label)

	dns1123LabelPrefix := DNS1123LabelPrefix("")
	Default.Add("dns1123-label-prefix", &dns1123LabelPrefix, isDNS1123LabelPrefix)

	dns1123Subdomain := DNS1123Subdomain("")
	Default.Add("dns1123-subdomain", &dns1123Subdomain, isDNS1123Subdomain)

	dns1123SubdomainPrefix := DNS1123SubdomainPrefix("")
	Default.Add("dns1123-subdomain-prefix", &dns1123SubdomainPrefix, isDNS1123SubdomainPrefix)

	qualifiedName := QualifiedName("")
	Default.Add("qualified-name", &qualifiedName, isQualifiedName)

	quantity := Quantity("")
	Default.Add("quantity", &quantity, isQuantity)

	ip := IP("")
	Default.Add("ip", &ip, isIP)

	semver := Semver("")
	Default.Add("semver", &semver, isSemver)
}

// DeepCopyInto copies the receiver into out. out must be non-nil.
func (d *DNS1035Label) DeepCopyInto(out *DNS1035Label) {
	*out = *d
}

// DeepCopy creates a deep copy of DNS1035Label
func (d *DNS1035Label) DeepCopy() *DNS1035Label {
	if d == nil {
		return nil
	}
	out := new(DNS1035Label)
	d.DeepCopyInto(out)
	return out
}

// DeepCopyInto copies the receiver into out. out must be non-nil.
func (d *DNS1035LabelPrefix) DeepCopyInto(out *DNS1035LabelPrefix) {
	*out = *d
}

// DeepCopy creates a deep copy of DNS1035LabelPrefix
func (d *DNS1035LabelPrefix) DeepCopy() *DNS1035LabelPrefix {
	if d == nil {
		return nil
	}
	out := new(DNS1035LabelPrefix)
	d.DeepCopyInto(out)
	return out
}

// DeepCopyInto copies the receiver into out. out must be non-nil.
func (d *DNS1123Label) DeepCopyInto(out *DNS1123Label) {
	*out = *d
}

// DeepCopy creates a deep copy of DNS1123Label
func (d *DNS1123Label) DeepCopy() *DNS1123Label {
	if d == nil {
		return nil
	}
	out := new(DNS1123Label)
	d.DeepCopyInto(out)
	return out
}

// DeepCopyInto copies the receiver into out. out must be non-nil.
func (d *DNS1123LabelPrefix) DeepCopyInto(out *DNS1123LabelPrefix) {
	*out = *d
}

// DeepCopy creates a deep copy of DNS1123LabelPrefix
func (d *DNS1123LabelPrefix) DeepCopy() *DNS1123LabelPrefix {
	if d == nil {
		return nil
	}
	out := new(DNS1123LabelPrefix)
	d.DeepCopyInto(out)
	return out
}

// DeepCopyInto copies the receiver into out. out must be non-nil.
func (d *DNS1123Subdomain) DeepCopyInto(out *DNS1123Subdomain) {
	*out = *d
}

// DeepCopy creates a deep copy of DNS1123Subdomain
func (d *DNS1123Subdomain) DeepCopy() *DNS1123Subdomain {
	if d == nil {
		return nil
	}
	out := new(DNS1123Subdomain)
	d.DeepCopyInto(out)
	return out
}

// DeepCopyInto copies the receiver into out. out must be non-nil.
func (d *DNS1123SubdomainPrefix) DeepCopyInto(out *DNS1123SubdomainPrefix) {
	*out = *d
}

// DeepCopy creates a deep copy of DNS1123SubdomainPrefix
func (d *DNS1123SubdomainPrefix) DeepCopy() *DNS1123SubdomainPrefix {
	if d == nil {
		return nil
	}
	out := new(DNS1123SubdomainPrefix)
	d.DeepCopyInto(out)
	return out
}

// DeepCopyInto copies the receiver into out. out must be non-nil.
func (q *QualifiedName) DeepCopyInto(out *QualifiedName) {
	*out = *q
}

// DeepCopy creates a deep copy of QualifiedName
func (q *QualifiedName) DeepCopy() *QualifiedName {
	if q == nil {
		return nil
	}
	out := new(QualifiedName)
	q.DeepCopyInto(out)
	return out
}

// DeepCopyInto copies the receiver into out. out must be non-nil.
func (q *Quantity) DeepCopyInto(out *Quantity) {
	*out = *q
}

// DeepCopy creates a deep copy of Quantity
func (q *Quantity) DeepCopy() *Quantity {
	if q == nil {
		return nil
	}
	out := new(Quantity)
	q.DeepCopyInto(out)
	return out
}

// DeepCopyInto copies the receiver into out. out must be non-nil.
func (ip *IP) DeepCopyInto(out *IP) {
	*out = *ip
}

// DeepCopy creates a deep copy of IP
func (ip *IP) DeepCopy() *IP {
	if ip == nil {
		return nil
	}
	out := new(IP)
	ip.DeepCopyInto(out)
	return out
}

func maskTrailingDash(name string) string {
	if len(name) > 1 && strings.HasSuffix(name, "-") {
		return name[:len(name)-2] + "a"
	}
	return name
}
