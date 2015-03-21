package constructor

import (
	"net/url"

	"github.com/skriptble/hyper/profiles/alps"
)

type descriptor struct {
	ID          string       `json:"id,omitempty"`
	Href        string       `json:"href,omitempty"`
	Name        string       `json:"name,omitempty"`
	HyperType   alps.Control `json:"type,omitempty"`
	Descriptors []descriptor `json:"descriptor,omitempty"`
	Rt          string       `json:"rt,omitempty"`
	Doc         *doc         `json:"doc,omitempty"`
	Ext         *ext         `json:"ext,omitempty"`
	Link        *link        `json:"link,omitempty"`
	index       map[string]struct{}
}

// NewDescriptor creates a new "descriptor" element that can be added to a
// profile or a descriptor.
func NewDescriptor(id string, href *url.URL, opts ...Option) (Option, error) {
	if id == "" && href == nil {
		return nil, ErrMissingID
	}

	d := new(descriptor)
	d.ID = id
	if href != nil {
		d.Href = href.String()
	}
	d.index = make(map[string]struct{})
	d.index[d.ID] = struct{}{}

	for _, opt := range opts {
		err := opt(d)
		if err != nil {
			return nil, err
		}
	}

	return func(i interface{}) error {
		innerDescriptor := d
		switch t := i.(type) {
		case *profile:
			// Compare the child descriptor's index to its potential parent's
			// index
			for idx := range innerDescriptor.index {
				_, conflict := t.index[idx]
				if conflict {
					return ErrIDConflict
				}
				t.index[idx] = struct{}{}
			}
			t.Descriptors = append(t.Descriptors, *innerDescriptor)
		case *descriptor:
			// Compare the child descriptor's index to its potential parent's
			// index
			for idx := range innerDescriptor.index {
				_, conflict := t.index[idx]
				if conflict {
					return ErrIDConflict
				}
				t.index[idx] = struct{}{}
			}
			t.Descriptors = append(t.Descriptors, *innerDescriptor)
		default:
			return ErrTypeUnknown
		}
		return nil
	}, nil
}

// NewName creates a new "name" element that can be added to a descriptor.
func NewName(name string) Option {
	return func(i interface{}) error {
		d, ok := i.(*descriptor)
		if !ok {
			return ErrTypeUnknown
		}
		d.Name = name
		return nil
	}
}

// NewType creates a new "type" element that can be added to a descriptor.
func NewType(control alps.Control) Option {
	return func(i interface{}) error {
		d, ok := i.(*descriptor)
		if !ok {
			return ErrTypeUnknown
		}
		d.HyperType = control
		return nil
	}
}

// NewRt creates a new "rt" element that can be added to a descriptor.
func NewRt(rt string) Option {
	return func(i interface{}) error {
		d, ok := i.(*descriptor)
		if !ok {
			return ErrTypeUnknown
		}
		d.Rt = rt
		return nil
	}
}
