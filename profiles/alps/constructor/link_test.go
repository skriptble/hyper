package constructor

import (
	"net/url"
	"reflect"
	"testing"
)

func TestLink(t *testing.T) {
	// Should be able to create a new link
	href, err := url.Parse("http://example.net/foo")
	if err != nil {
		t.Errorf("Unexpected Error during url.Parse: %v", err)
	}
	rel := "foo"

	linkOpt := NewLink(href, rel)

	want := &link{
		Href: href.String(),
		Rel:  rel,
	}
	// Should be able to add new link to profile
	descriptorOpt, err := NewDescriptor("foo", nil, linkOpt)
	if err != nil {
		t.Errorf("Unexpected Error during NewDescriptor: %v", err)
	}

	prof, err := NewProfile(linkOpt, descriptorOpt)
	if err != nil {
		t.Errorf("Unexpected Error during NewProfile: %v", err)
	}

	if !reflect.DeepEqual(prof.profile.Link, want) {
		t.Errorf("Should be able to add a new link to profile. Wanted %v, got %v", want, prof.profile.Link)
	}

	// Should be able to add new link to descriptor
	got := prof.profile.Descriptors[0].Link
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Should be able to add a new link to descriptor. Wanted %v, got %v", want, prof.profile.Link)
	}

	// Should not be able to add link to unknown type
	err = func(opt Option) error {
		return opt(descriptor{})
	}(linkOpt)
	if err != ErrTypeUnknown {
		t.Errorf("Should not be able to add link to unknown type. Wanted %v, got %v", ErrTypeUnknown, err)
	}
}
