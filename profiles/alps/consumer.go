package alps

import (
	"encoding/json"
	"net/url"
)

// Profile represents the "alps" root element as defined by the ALPS profile
// specification (2.2.1)
type Profile interface {
	// Version returns the ALPS version of this profile
	Version() Version
	// Doc returns the Doc element, if any, associated with this profile.
	// Returns nil if no Doc is associated.
	Doc() Doc
	// Ext returns the Ext element, if any, associated with this profile.
	// Returns nil if no Ext is associated.
	Ext() Ext
	// Link returns the Link element, if any, associated with this profile.
	// Return snil if no Ext is associated.
	Link() Link
	// Descriptor returns the descriptors with the provided Id or name provided.
	// A url in string format may also be provided. If no identifiers are
	// provided then Descriptor will return all the Descriptors associated with
	// this profile.
	Descriptor(identifiers ...string) <-chan Descriptor
}

// Doc represents the "doc" element as defined by the ALPS profile spec (2.2.2)
type Doc interface {
	// Href returns the Href of the documentation
	Href() url.URL
	// Format returns the specified format for the documentation. The valid
	// values are: text, html, asciidoc, and markdown.
	Format() string
	// Value returns the value of the documentation
	Value() string

	json.Unmarshaler
}

// Descriptor represents the "descriptor" element as defined by the ALPS
// profile specification (2.2.3)
type Descriptor interface {
	// Id returns the profile-wide unique identifier for this Descriptor.
	ID() string
	// Href returns the href, if any, of a Descriptor related to this
	// Descirptor. The returned url.URL must contain a fragment the references
	// the related Descritor.
	Href() url.URL
	// Name returns the common, non-unique value associated with this
	// Descriptor. For more information on this value see section 2.2.7.1 of
	// the ALPS speficiation.
	Name() string
	// Doc returns the Doc element, if any, associated with this profile.
	// Returns nil if no Doc is associated.
	Doc() Doc
	// Ext returns the Ext element, if any, associated with this profile.
	// Returns nil if no Ext is associated.
	Ext() Ext
	// Type returns the "type" property of this Descriptor as documented in
	// section 2.2.12. The valid returns are: semantic, safe, idempotent, and
	// unsafe. If no "type" is specific on the profile, semantic is always
	// returned.
	Type() Control
	// Descriptor returns the descriptors with the provided Id or name provided.
	// A url in string format may also be provided. If no identifiers are
	// provided then Descriptor will return all the Descriptors associated with
	// this profile.
	Descriptor(identifiers ...string) <-chan Descriptor
	// Link returns the Link element, if any, associated with this profile.
	// Return snil if no Ext is associated.
	Link() Link
	// Rt returns the string representation of the resource type that will be
	// returned from this Descriptor.
	Rt() string

	json.Unmarshaler
}

// Ext represents the "ext" element as defined by the ALPS profile
// specification (2.2.4)
type Ext interface {
	// Id returns the identifier for this ext element.
	ID() string
	// Href returns the url.URL which points to a document explaining the
	// use of this ext element.
	Href() url.URL
	// Value returns the string representation of the value of this ext
	// element.
	Value() string

	json.Unmarshaler
}

// Link represents the "link" element as defined by the ALPS profile
// specification (2.2.8)
type Link interface {
	// Href returns the url of an external document whose relationship can be
	// described by the value of the Rel() method.
	Href() url.URL
	// Rel returns a string representing a registered relation type or
	// an extension relation type. For more information see section (2.2.10) of
	// the ALPS profile specification.
	Rel() string

	json.Unmarshaler
}
