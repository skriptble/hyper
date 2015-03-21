package constructor

import (
	"net/url"
	"reflect"
	"testing"

	"github.com/skriptble/hyper/profiles/alps"
)

func TestDescriptor(t *testing.T) {
	// Should not be able to create descriptor without id or href
	_, err := NewDescriptor("", nil)
	if err != ErrMissingID {
		t.Errorf("Should not be able to create a descriptor without id or href. Wanted %+v, got %+v", ErrMissingID, err)
	}

	// Should be able to add name to descriptor
	nameOpt := NewName("foo")
	wantName := "foo"

	// Should not be able to attach name to unknown type
	err = failAttach(nameOpt)
	if err != ErrTypeUnknown {
		t.Error("Should not be able to attach name to unknown type")
		t.Errorf("Wanted %v, got %v", ErrTypeUnknown, err)
	}

	// Should be able to add type to descriptor
	typeOpt := NewType(alps.Safe)
	wantType := alps.Safe

	// Should not be able to attach type to unknown type
	err = failAttach(typeOpt)
	if err != ErrTypeUnknown {
		t.Error("Should not be able to attach type to unknown type")
		t.Errorf("Wanted %v, got %v", ErrTypeUnknown, err)
	}

	// Should be able to add rt to descriptor
	rtOpt := NewRt("bar")
	wantRt := "bar"

	// Should not be able to attach rt to an unknown type
	err = failAttach(rtOpt)
	if err != ErrTypeUnknown {
		t.Error("Should not be able to attach rt to unknown type")
		t.Errorf("Wanted %v, got %v", ErrTypeUnknown, err)
	}

	// Should be able to attach a descriptor to a descriptor
	quuxDescriptor, err := NewDescriptor("quux", nil)
	if err != nil {
		t.Errorf("Unexpected error from NewDescriptor: %v", err)
	}
	wantDescriptor := descriptor{ID: "quux", index: map[string]struct{}{"quux": struct{}{}}}

	// Should be able to create a descriptor with an href
	href, err := url.Parse("http://example.com/foo")
	if err != nil {
		t.Errorf("Unexpected error from url.Parse: %v", err)
	}
	wantHref := href.String()

	descriptorOpt, err := NewDescriptor("baz", href, nameOpt, typeOpt, rtOpt, quuxDescriptor)
	if err != nil {
		t.Errorf("Unexected error from NewDescriptor: %v", err)
	}

	// Should not be able to create a descriptor with an identifier conflict
	_, err = NewDescriptor("baz", nil, descriptorOpt)
	if err != ErrIDConflict {
		t.Error("Should not be able to create a descriptor with an identifier conflict")
		t.Errorf("Wanted %v, got %v", ErrIDConflict, err)
	}

	// Should not be able to attach a descriptor to an unknown type
	err = failAttach(descriptorOpt)
	if err != ErrTypeUnknown {
		t.Error("Should not be able to attach descriptor to unknown type")
		t.Errorf("Wanted %v, got %v", ErrTypeUnknown, err)
	}

	// Should be able to attach a descriptor to a profile
	prof, err := NewProfile(descriptorOpt)
	if err != nil {
		t.Errorf("Unexpected error from NewProfile: %v", err)
	}
	d0 := prof.profile.Descriptors[0]
	gotName := d0.Name
	gotType := d0.HyperType
	gotRt := d0.Rt
	gotDescriptor := d0.Descriptors[0]
	gotHref := d0.Href
	if wantName != gotName {
		t.Error("Should be able to attach name to descriptor")
		t.Errorf("Wanted %v, got %v", wantName, gotName)
	}
	if wantType != gotType {
		t.Error("Should be able to attach type to descriptor")
		t.Errorf("Wanted %v, got %v", wantType, gotType)
	}
	if wantRt != gotRt {
		t.Error("Should be able to attach rt to descriptor")
		t.Errorf("Wanted %v, got %v", wantRt, gotRt)
	}
	if !reflect.DeepEqual(wantDescriptor, gotDescriptor) {
		t.Error("Should be able to attach descriptor to descriptor")
		t.Errorf("Wanted %+v, got %+v", wantDescriptor, gotDescriptor)
	}
	if wantHref != gotHref {
		t.Error("Should be able to attach href to descriptor")
		t.Errorf("Wanted %v, got %v", wantHref, gotHref)
	}

	// Should not be able to attach a conflicting descriptor to a profile
	conflictDescriptor, err := NewDescriptor("quux", nil)
	if err != nil {
		t.Errorf("Unexected error from NewDescriptor: %v", err)
	}
	_, err = NewProfile(descriptorOpt, conflictDescriptor)
	if err != ErrIDConflict {
		t.Error("Should not be able to attach a conflicting descriptor to a profile")
		t.Errorf("Wanted %v, got %v", ErrIDConflict, err)
	}

}

func failAttach(opt Option) error {
	return opt(descriptor{})
}
