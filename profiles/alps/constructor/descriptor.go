package constructor

import (
	"fmt"
	"net/url"

	"github.com/skriptble/hyper/profiles/alps"
)

type descriptor struct {
	ID          string
	Href        url.URL
	Name        string
	HyperType   alps.Control
	Descriptors []descriptor
	Rt          string
	Doc         doc
	Ext         ext
	Link        link
	index       map[string]struct{}
}

// NewDescriptor creates a new "descriptor" element that can be added to a
// profile or a descriptor.
func NewDescriptor(id string, href url.URL, opts ...Option) (Option, error) {
	if id == "" && href == (url.URL{}) {
		return nil, ErrMissingID
	}

	d := new(descriptor)
	d.ID = id
	d.Href = href
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
			fmt.Println(t)
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
