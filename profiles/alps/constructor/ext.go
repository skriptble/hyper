package constructor

import "net/url"

type ext struct {
	ID    string `json:"id,omitempty"`
	Href  string `json:"href,omitempty"`
	Value string `json:"value,omitempty"`
}

// NewExt creates an "ext" element that can be added to a profile or descriptor.
func NewExt(id, value string, href *url.URL) Option {
	return func(i interface{}) error {
		e := new(ext)
		e.ID = id
		e.Value = value

		if href != nil {
			e.Href = href.String()
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
