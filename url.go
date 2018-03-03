// Package surly wraps std library's url to make it marshalable and unmarshalable to various formats
package surly

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"net/url"
	"strings"
)

// URL is a wrapper for the url.URL type that allows json unmarshaling
type URL string

// New creates a new URL from a url.URL.
func New(rawurl string) (URL, error) {
	_, err := url.Parse(rawurl)
	if err != nil {
		return "", err
	}
	return URL(rawurl), nil
}

// Parsed returns url.URL representation of the URL string
func (u URL) Parsed() *url.URL {
	parsed, _ := url.Parse(string(u))
	return parsed
}

// UnmarshalJSON implements the json.Unmarshaler interface
func (u URL) UnmarshalJSON(b []byte) error {
	var err error
	val := string(bytes.TrimSpace(bytes.Trim(b, `"`)))
	u, err = New(val)
	return err
}

// MarshalJSON implements the json.Marshaler interface
func (u URL) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, u)), nil
}

// UnmarshalXML immplements the xml.Unmarshaler interface
func (u URL) UnmarshalXML(d *xml.Decoder, start xml.StartElement) (err error) {
	var v string
	d.DecodeElement(&v, &start)
	val := strings.TrimSpace(v)
	u, err = New(val)
	return err
}

// ResolveReference resolves a URI reference by delegating to the underlying url.URL's implementation
func (u URL) ResolveReference(other URL) URL {
	return URL(u.Parsed().ResolveReference(other.Parsed()).String())
}
