package constructor

import (
	"net/url"
	"reflect"
	"testing"
)

func TestExt(t *testing.T) {
	// Should be able to create a new link
	href, err := url.Parse("http://example.net/foo")
	if err != nil {
		t.Errorf("Unexpected Error during url.Parse: %v", err)
	}
	id := "foo"
	value := "wat?wat?foo!foo!"

	extOpt := NewExt(id, value, href)

	want := &ext{
		ID:    id,
		Href:  href.String(),
		Value: value,
	}
	// Should be able to add new link to profile
	descriptorOpt, err := NewDescriptor("foo", nil, extOpt)
	if err != nil {
		t.Errorf("Unexpected Error during NewDescriptor: %v", err)
	}

	prof, err := NewProfile(extOpt, descriptorOpt)
	if err != nil {
		t.Errorf("Unexpected Error during NewProfile: %v", err)
	}

	if !reflect.DeepEqual(prof.profile.Ext, want) {
		t.Errorf("Should be able to add a new ext to profile. Wanted %+v, got %+v", want, prof.profile.Ext)
	}

	// Should be able to add new link to descriptor
	got := prof.profile.Descriptors[0].Ext
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Should be able to add a new ext to descriptor. Wanted %+v, got %+v", want, prof.profile.Ext)
	}

	// Should not be able to add link to unknown type
	err = func(opt Option) error {
		return opt(descriptor{})
	}(extOpt)
	if err != ErrTypeUnknown {
		t.Errorf("Should not be able to add link to unknown type. Wanted %v, got %v", ErrTypeUnknown, err)
	}
}
