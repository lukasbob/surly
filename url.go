// Package surly provides a URL string that is to marshalable and unmarshalable to various formats
package surly

import (
	"net/url"
	"strings"
)

// URL wraps a string that is guaranteed to be a valid URL
type URL struct {
	u string
}

// MustParse creates a new URL from a valid URL string, or panics on error
func MustParse(rawurl string) URL {
	u, err := New(rawurl)
	if err != nil {
		panic(err)
	}
	return u
}

// New creates a new URL from a string
func New(rawurl string) (u URL, err error) {
	if _, err = url.Parse(rawurl); err == nil {
		u.u = rawurl
	}
	return
}

// Parsed returns url.URL representation of the URL string
func (u URL) Parsed() *url.URL {
	parsed, _ := url.Parse(u.u)
	return parsed
}

//MarshalText implements the encoding.TextMarshaler interface
func (u URL) MarshalText() ([]byte, error) {
	return []byte(u.String()), nil
}

// UnmarshalText implements the encoding.TextUnarshaler interface
func (u *URL) UnmarshalText(text []byte) (err error) {
	*u, err = New(strings.TrimSpace(string(text)))
	return err
}

// ResolveReference resolves a URI reference by delegating to the underlying url.URL's implementation
func (u URL) ResolveReference(other URL) URL {
	return URL{u.Parsed().ResolveReference(other.Parsed()).String()}
}

func (u URL) String() string {
	return u.u
}
