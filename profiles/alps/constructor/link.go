package constructor

import "net/url"

type link struct {
	href url.URL
	rel  string
}

// NewLink creates an "link" element that can be added to a profile or
// descriptor.
func NewLink(href url.URL, rel string) Option {
	return func(i interface{}) error {
		l := link{
			href: href,
			rel:  rel,
		}
		switch t := i.(type) {
		case *profile:
			t.Link = l
		case *descriptor:
			t.Link = l
		default:
			return ErrTypeUnknown
		}
		return nil
	}
}
