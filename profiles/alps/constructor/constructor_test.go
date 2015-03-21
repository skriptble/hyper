package constructor

import (
	"encoding/json"
	"net/url"
	"reflect"
	"testing"

	"github.com/skriptble/hyper/profiles/alps"
)

func TestProfile(t *testing.T) {
	// Should be able to create a new version
	versionOpt := NewVersion(alps.Version("2.0"))
	wantVersion := alps.Version("2.0")

	// Should be able to attach a version to a profile
	prof, err := NewProfile(versionOpt)
	if err != nil {
		t.Errorf("Unexpected error from NewProfile: %v", err)
	}
	gotVersion := prof.profile.Version
	if wantVersion != gotVersion {
		t.Error("Should be able to attach a version to a profile")
		t.Errorf("Wanted %+v, got %+v", wantVersion, gotVersion)
	}

	// Should not be able to attach a version to an unknonw type
	err = func(opt Option) error {
		return opt(profile{})
	}(versionOpt)
	if err != ErrTypeUnknown {
		t.Error("Should not be able to attach a version to an unknown type")
		t.Errorf("Wanted %v, got %v", ErrTypeUnknown, err)
	}

	// Should be able to marshal profile into JSON
	prof, err = NewProfile()
	if err != nil {
		t.Errorf("Unexpected error from NewProfile: %v", err)
	}
	got, err := json.Marshal(prof)
	if err != nil {
		t.Errorf("Unexpected error from json.Marshal: %v", err)
	}
	want := []byte(`{"alps":{"version":"1.0","descriptor":null}}`)
	if !reflect.DeepEqual(got, want) {
		t.Error("Should be able to marshal profile into JSON")
		t.Errorf("Wanted %v, got %v", want, got)
	}
}

func TestDoc(t *testing.T) {
	// Should be able to create a new doc
	href, err := url.Parse("http://example.net/foo")
	if err != nil {
		t.Errorf("Unexpected Error during url.Parse: %v", err)
	}
	format := "foo"
	value := "wat?wat?foo!foo!"

	docOpt := NewDoc(href, format, value)

	want := &doc{
		Href:   href.String(),
		Format: format,
		Value:  value,
	}
	// Should be able to add new doc to profile
	descriptorOpt, err := NewDescriptor("foo", nil, docOpt)
	if err != nil {
		t.Errorf("Unexpected Error during NewDescriptor: %v", err)
	}

	prof, err := NewProfile(docOpt, descriptorOpt)
	if err != nil {
		t.Errorf("Unexpected Error during NewProfile: %v", err)
	}

	if !reflect.DeepEqual(prof.profile.Doc, want) {
		t.Errorf("Should be able to add a new ext to profile. Wanted %+v, got %+v", want, prof.profile.Doc)
	}

	// Should be able to add new link to descriptor
	got := prof.profile.Descriptors[0].Doc
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Should be able to add a new ext to descriptor. Wanted %+v, got %+v", want, prof.profile.Doc)
	}

	// Should not be able to add link to unknown type
	err = func(opt Option) error {
		return opt(descriptor{})
	}(docOpt)
	if err != ErrTypeUnknown {
		t.Errorf("Should not be able to add link to unknown type. Wanted %v, got %v", ErrTypeUnknown, err)
	}
}
