package alps

import (
	"encoding/json"
	"net/url"
)

// Constructor is used to create a profile document.
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

// ConstructorDoc is a Doc element that can be used in the construction of an
// ALPS profile.
type ConstructorDoc interface {
	Doc
	json.Marshaler
}

// ConstructorDescriptor is a Descriptor element that can be used in the
// construction of an ALPS profile.
type ConstructorDescriptor interface {
	Descriptor
	json.Marshaler
}

// ConstructorExt is an Ext element that can be used in the construction of an
// ALPS profile.
type ConstructorExt interface {
	Ext
	json.Marshaler
}

// ConstructorLink is a Link element that can be used in the construction of an
// ALPS profile.
type ConstructorLink interface {
	Link
	json.Marshaler
}
