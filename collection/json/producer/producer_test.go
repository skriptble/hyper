package producer

import (
	"encoding/json"
	"os"
	"reflect"
	"testing"

	"github.com/skriptble/hyper/collection/json"
)

func TestCollection(t *testing.T) {
	// Should be able to marshal json
	c := Collection{collection{
		Version: cj.V1,
	}}
	got, err := json.Marshal(c)
	if err != nil {
		t.Errorf("Unexpect error: %v", err)
	}
	os.Stdout.Write(got)
}

func TestError(t *testing.T) {
	errOpt := NewError("foo", "bar", "baz")
	// Should not be able to attach error to unknown type
	err := func(opt Option) error {
		return opt(collection{})
	}(errOpt)
	if err != ErrTypeUnknown {
		t.Error("Should not be able to attach an error to an unknown type")
		t.Errorf("Wanted %v, got %v", ErrTypeUnknown, err)
	}

	// Should be able to attach error to collection
	c, err := NewCollection(errOpt)
	if err != nil {
		t.Errorf("Unexpected error from NewCollection: %v", err)
	}
	want := &cjError{
		Title:   "foo",
		Code:    "bar",
		Message: "baz",
	}
	got := c.collection.Error
	if !reflect.DeepEqual(want, got) {
		t.Error("Should be able to attach an error to a collection")
		t.Errorf("Wanted %+v, got %+v", want, got)
	}
}

func TestTemplate(t *testing.T) {
	// Should not be able to attach incorrect option to template
	errOpt := NewError("foo", "bar", "baz")
	_, err := NewTemplate(errOpt)
	if err != ErrTypeUnknown {
		t.Error("Should not be able to attach incorrect option to template")
		t.Errorf("Wanted %v, got %v", ErrTypeUnknown, err)
	}

	// Should not be able to attach template to unknown type
	tmplOpt, err := NewTemplate()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	err = func(opt Option) error {
		return opt(struct{}{})
	}(tmplOpt)
	if err != ErrTypeUnknown {
		t.Error("Should not be able to attach template to unknown type")
		t.Errorf("Wanted %v, got %v", ErrTypeUnknown, err)
	}

	// Should be able to attach a template to a collection
	datumOpt := NewDatum("foo", "bar", "baz")
	tmplOpt, err = NewTemplate(datumOpt)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	c, err := NewCollection(tmplOpt)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	want := &template{
		Data: []datum{
			{
				Name:   "foo",
				Value:  "bar",
				Prompt: "baz",
			},
		},
	}
	got := c.collection.Template
	if !reflect.DeepEqual(want, got) {
		t.Error("Should be able to attach a template to a collection")
		t.Errorf("Wanted %+v, got %+v", want, got)
	}
}

func TestDatum(t *testing.T) {
	// Should not be able to attach datum to unknown type

	// Should be able to attach datum to a template

	// Should be able to attach datum to an item
}
