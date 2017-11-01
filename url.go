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
type URL struct {
	*url.URL
}

func MustParse(rawurl string) *URL {
	u, err := Parse(rawurl)
	if err != nil {
		panic(err)
	}
	return u
}

// Parse creates a new URL, delegating to url.Parse for parsing.
func Parse(rawurl string) (*URL, error) {
	u, err := url.Parse(rawurl)
	if err == nil {
		return New(u), err
	}
	return nil, err
}

// New creates a new URL
func New(u *url.URL) *URL {
	return &URL{URL: u}
}

// UnmarshalJSON implements the json.Unmarshaler interface
func (u *URL) UnmarshalJSON(b []byte) error {
	parsed, err := url.Parse(string(bytes.TrimSpace(bytes.Trim(b, `"`))))
	if err == nil {
		u.URL = parsed
	}
	return err
}

// MarshalJSON implements the json.Marshaler interface
func (u *URL) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, u.String())), nil
}

// UnmarshalXML immplements the xml.Unmarshaler interface
func (u *URL) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v string
	d.DecodeElement(&v, &start)
	parsed, err := url.Parse(strings.TrimSpace(v))
	if err == nil {
		u.URL = parsed
	}
	return err
}

// ResolveReference resolves a URI reference by delegating to the underlying url.URL's implementation
func (u *URL) ResolveReference(other *URL) *URL {
	return New(u.URL.ResolveReference(other.URL))
}
