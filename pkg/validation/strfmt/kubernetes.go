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
	"net"
	"strings"

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
	errs := validation.IsDNS1035Label(str)
	return len(errs) == 0
}

func isDNS1035LabelPrefix(str string) bool {
	// A prefix must follow the same rules as a full label, except it can end with a hyphen
	if str == "" {
		return false
	}
	if len(str) > 63 {
		return false
	}
	// Must start with a letter (not a number)
	if !(str[0] >= 'a' && str[0] <= 'z') {
		return false
	}
	for i := 0; i < len(str); i++ {
		c := str[i]
		if !(c >= 'a' && c <= 'z' || c >= '0' && c <= '9' || c == '-') {
			return false
		}
	}
	return true
}

func isDNS1123Label(str string) bool {
	errs := validation.IsDNS1123Label(str)
	return len(errs) == 0
}

func isDNS1123LabelPrefix(str string) bool {
	// A prefix must follow the same rules as a full label, except it can end with a hyphen
	if str == "" {
		return false
	}
	if len(str) > 63 {
		return false
	}
	if str[0] == '-' {
		return false
	}
	for i := 0; i < len(str); i++ {
		c := str[i]
		if !(c >= 'a' && c <= 'z' || c >= '0' && c <= '9' || c == '-') {
			return false
		}
	}
	return true
}

func isDNS1123Subdomain(str string) bool {
	// First check the total length
	if len(str) > 253 {
		return false
	}

	// Then check each segment
	segments := strings.Split(str, ".")
	for _, segment := range segments {
		if len(segment) > 63 {
			return false
		}
	}

	// Finally use the standard validation
	errs := validation.IsDNS1123Subdomain(str)
	return len(errs) == 0
}

func isDNS1123SubdomainPrefix(str string) bool {
	// A prefix must follow the same rules as a full subdomain, except the last segment can end with a hyphen
	if str == "" {
		return false
	}
	if len(str) > 253 {
		return false
	}
	segments := strings.Split(str, ".")
	for i, segment := range segments {
		if segment == "" {
			return false
		}
		if len(segment) > 63 {
			return false
		}
		if segment[0] == '-' {
			return false
		}
		// Only validate the ending hyphen for non-final segments
		if i < len(segments)-1 && segment[len(segment)-1] == '-' {
			return false
		}
		for _, c := range segment {
			if !(c >= 'a' && c <= 'z' || c >= '0' && c <= '9' || c == '-') {
				return false
			}
		}
	}
	return true
}

func isQualifiedName(str string) bool {
	// Must contain at least one '/'
	parts := strings.Split(str, "/")
	if len(parts) < 2 {
		return false
	}

	// The part before the first '/' must be a valid DNS1123 subdomain
	if !isDNS1123Subdomain(parts[0]) {
		return false
	}

	// The remaining parts must be valid path segments
	for _, part := range parts[1:] {
		if part == "" {
			return false
		}
		// Path segments can't start or end with a hyphen
		if strings.HasPrefix(part, "-") || strings.HasSuffix(part, "-") {
			return false
		}
		// Path segments can contain alphanumeric characters, '-', '_', and '.'
		for _, c := range part {
			if !(c >= 'a' && c <= 'z' || c >= '0' && c <= '9' || c == '-' || c == '_' || c == '.') {
				return false
			}
		}
	}

	return true
}

func isQuantity(str string) bool {
	// First try to parse the quantity
	_, err := resource.ParseQuantity(str)
	if err != nil {
		return false
	}

	// Then validate the format
	// Remove any leading sign
	if str != "" && (str[0] == '+' || str[0] == '-') {
		str = str[1:]
	}

	// Split into number and suffix
	var number, suffix string
	for i, c := range str {
		if !(c >= '0' && c <= '9' || c == '.') {
			number = str[:i]
			suffix = str[i:]
			break
		}
	}
	if suffix == "" {
		number = str
	}

	// Validate number format (should be a valid decimal)
	dots := 0
	for _, c := range number {
		if c == '.' {
			dots++
		} else if !(c >= '0' && c <= '9') {
			return false
		}
	}
	if dots > 1 {
		return false
	}

	// Validate suffix
	validSuffixes := map[string]bool{
		"":   true,
		"m":  true,
		"u":  true,
		"µ":  true,
		"n":  true,
		"Ki": true,
		"Mi": true,
		"Gi": true,
		"Ti": true,
		"Pi": true,
		"Ei": true,
	}

	return validSuffixes[suffix]
}

func isIP(str string) bool {
	// Use net.ParseIP to validate both IPv4 and IPv6 addresses
	ip := net.ParseIP(str)
	return ip != nil
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
