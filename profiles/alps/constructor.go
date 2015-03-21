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

type ConstructorDoc interface {
	Doc
	json.Marshaler
}

type ConstructorDescriptor interface {
	Descriptor
	json.Marshaler
}

type ConstructorExt interface {
	Ext
	json.Marshaler
}

type ConstructorLink interface {
	Link
	json.Marshaler
}
