package constructor

import "net/url"

type ext struct {
	ID    string
	Href  url.URL
	Value string
}

// NewExt creates an "ext" element that can be added to a profile or descriptor.
func NewExt(id, value string, href url.URL) Option {
	return func(i interface{}) error {
		e := ext{
			ID:    id,
			Href:  href,
			Value: value,
		}
		switch t := i.(type) {
		case *profile:
			t.Ext = e
		case *descriptor:
			t.Ext = e
		default:
			return ErrTypeUnknown
		}

		return nil
	}
}
