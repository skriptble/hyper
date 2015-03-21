package alps

import (
	"encoding/json"
	"net/url"
)

type Constructor interface {
	// Version sets the version number for the profile
	Version(Version)
	// Base sets the base URL for the profile
	Base(url.URL)
	// Doc adds a documentation link to the profile
	Doc(ConstructorDoc)
	// AddDescriptor adds the given descriptor to the profile
	AddDescriptor(ConstructorDescriptor)
	// AddExt adds the given ext to the profile
	AddExt(ConstructorExt)
	// AddLink adds the given link to the profile
	AddLink(ConstructorLink)
}

type Doc interface {
	// Href returns the Href of the documentation
	Href() url.URL
	Format() string
	Value() string

	json.Unmarshaler
}

type ConstructorDoc interface {
	Doc
	json.Marshaler
}

type Descriptor interface {
	json.Unmarshaler
}

type ConstructorDescriptor interface {
	Descriptor
	json.Marshaler
}

type Ext interface {
	Id() string
	Href() url.URL
	Value() string
	json.Unmarshaler
}

type ConstructorExt interface {
	Ext
	json.Marshaler
}

type Link interface {
	Href() url.URL
	Rel() string

	json.Unmarshaler
}

type ConstructorLink interface {
	Link
	json.Marshaler
}
