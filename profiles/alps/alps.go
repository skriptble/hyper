package alps

import "net/url"

type Format int

// Control represents an ALPS type of hypermedia control
type Control int
type Version string

const (
	_ Format = iota
	TEXT
	HTML
	ASCIIDOC
	MARKDOWN
)

const (
	_ Control = iota
	SEMANTIC
	SAFE
	IDEMPOTENT
	UNSAFE
)

const (
	V_1_0 Version = "1.0"
)

type profile struct {
	version     Version
	base        url.URL
	doc         doc
	descriptors map[string]descriptor
	ext         ext
	link        link
}

type doc struct {
	href   url.URL
	format string
	value  string
}

type descriptor struct {
	id          string
	href        url.URL
	doc         doc
	ext         ext
	name        string
	hypertype   Control
	descriptors map[string]descriptor
	link        link
	rt          string
}

type ext struct {
	id    string
	href  url.URL
	value string
}

type link struct {
	href url.URL
	rel  string
}

func (f Format) String() (format string) {
	switch f {
	case 1:
		format = "text"
	case 2:
		format = "html"
	case 3:
		format = "asciidoc"
	case 4:
		format = "markdown"
	default:
		format = ""
	}

	return
}

func (c Control) String() (format string) {
	switch c {
	case 1:
		format = "sematic"
	case 2:
		format = "safe"
	case 3:
		format = "idempotent"
	case 4:
		format = "unsafe"
	default:
		format = "semantic"
	}

	return format
}
