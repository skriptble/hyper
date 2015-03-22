package constructor

import (
	"encoding/json"
	"errors"
	"net/url"

	"github.com/skriptble/hyper/profiles/alps"
)

// IdentifierConflict is returned whenever there is a conflict between
// identifiers in a descriptor when they are added to a descriptor or a profile.
var ErrIDConflict = errors.New("constructor: there are two descriptors with the same identifier")

// ConstructorTypeUnknown is returned when an option is passed a type that it
// doesn't know how to handle.
var ErrTypeUnknown = errors.New("constructor: the argument to this option call is of an unknown type")

// DescriptorMissingID is returned when neither an ID nor Href are given as
// arguments for NewDescriptor.
var ErrMissingID = errors.New("constructor: descriptor is missing ID or Href")

// Constructor is used to create a profile document.
type Constructor interface {
}

// Option is a configuration option that can be passed into certain New
// functions. This allows the same Option type to be used across all the types
// that require configuring. If an unknown type is passed in as the argument,
// the function should return a ConstructorTypeUnknown error.
type Option func(interface{}) error

// Profile is created from a constructor. It can be used to connect different
// profiles.
type Profile struct {
	profile *profile
}

// MarshalJSON interface implementation
func (p Profile) MarshalJSON() ([]byte, error) {
	pfile := struct {
		*profile `json:"alps"`
	}{
		profile: p.profile,
	}
	return json.Marshal(pfile)
}

type profile struct {
	Version     alps.Version `json:"version"`
	Doc         *doc         `json:"doc,omitempty"`
	Descriptors []descriptor `json:"descriptor"`
	Ext         *ext         `json:"ext,omitempty"`
	Link        *link        `json:"link,omitempty"`
	base        url.URL
	index       map[string]struct{}
}

// NewProfile returns a Profile that can be marshalled to JSON. The profile will
// be configured according to the options given as arguments.
func NewProfile(opts ...Option) (Profile, error) {
	p := new(profile)
	p.Version = alps.V1
	p.index = make(map[string]struct{})
	for _, opt := range opts {
		err := opt(p)
		if err != nil {
			return Profile{}, err
		}
	}

	return Profile{p}, nil
}

// NewVersion creates a new "version" element that can be added to a profile.
func NewVersion(version alps.Version) Option {
	return func(i interface{}) error {
		p, ok := i.(*profile)
		if !ok {
			return ErrTypeUnknown
		}
		p.Version = version
		return nil
	}
}

type doc struct {
	Href   string `json:"href,omitempty"`
	Format string `json:"format,omitempty"`
	Value  string `json:"value,omitempty"`
}

// NewDoc creates a "doc" element that can be added to a profile or descriptor.
func NewDoc(href *url.URL, format, value string) Option {
	d := new(doc)
	if href != nil {
		d.Href = href.String()
	}
	d.Format = format
	d.Value = value
	return func(i interface{}) error {
		switch t := i.(type) {
		case *profile:
			t.Doc = d
		case *descriptor:
			t.Doc = d
		default:
			return ErrTypeUnknown
		}
		return nil
	}
}
