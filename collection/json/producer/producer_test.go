package producer

import (
	"encoding/json"
	"fmt"
	"net/url"
	"reflect"
	"testing"

	"github.com/skriptble/hyper/collection/json"
)

func TestCollection(t *testing.T) {
	// Should be able to marshal json
	c := Collection{collection{
		Version: cj.V1,
	}}
	b, err := json.Marshal(c)
	if err != nil {
		t.Errorf("Unexpect error: %v", err)
	}
	want := `{"collection":{"version":"1.0"}}`
	got := fmt.Sprintf("%s", b)
	if got != want {
		t.Error("Should be able to marshal collection into JSON")
		t.Errorf("Wanted %v, got %v", want, got)
	}

	// Should not be able to attach incorrect option to collection
	datumOpt := NewDatum("foo", "bar", "baz")
	_, err = NewCollection(datumOpt)
	if err != ErrTypeUnknown {
		t.Error("Should not be able to attach incorrect option to collection")
		t.Errorf("Wanted %v, got %v", ErrTypeUnknown, err)
	}
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
	datumOpt := NewDatum("foo", "bar", "baz")
	err := func(opt Option) error {
		return opt(struct{}{})
	}(datumOpt)
	if err != ErrTypeUnknown {
		t.Error("Should not be able to attach template to unknown type")
		t.Errorf("Wanted %v, got %v", ErrTypeUnknown, err)
	}

	// Should be able to attach datum to a template
	_, err = NewTemplate(datumOpt)
	if err != nil {
		t.Error("Should be able to attach datum to a template")
		t.Errorf("Wanted nil, but got %v", err)
	}

	// Should be able to attach datum to an item
	_, err = NewItem(url.URL{}, datumOpt)
	if err != nil {
		t.Error("Should be able to attach datum to a template")
		t.Errorf("Wanted nil, but got %v", err)
	}

	// Should be able to attach datum to a query
	_, err = NewQuery(url.URL{}, "", "", "", datumOpt)
	if err != nil {
		t.Error("Should be able to attach datum to a template")
		t.Errorf("Wanted nil, but got %v", err)
	}
}

func TestItem(t *testing.T) {
	// Should not be able attach incorrect option to item
	errOpt := NewError("foo", "bar", "baz")
	_, err := NewItem(url.URL{}, errOpt)
	if err != ErrTypeUnknown {
		t.Error("Should not be able to attach incorrect option to template")
		t.Errorf("Wanted %v, got %v", ErrTypeUnknown, err)
	}
	// Should not be able to attach item to an unknown type
	itemOpt, err := NewItem(url.URL{})
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	err = func(opt Option) error {
		return opt(struct{}{})
	}(itemOpt)
	if err != ErrTypeUnknown {
		t.Error("Should not be able to attach template to unknown type")
		t.Errorf("Wanted %v, got %v", ErrTypeUnknown, err)
	}
	// Should be able to attach an item to a collection
	datumOpt := NewDatum("foo", "bar", "baz")
	href := url.URL{Host: "example.com", Scheme: "http"}
	itemOpt, err = NewItem(href, datumOpt)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	c, err := NewCollection(itemOpt)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	want := item{
		Href: "http://example.com",
		Data: []datum{
			{
				Name:   "foo",
				Value:  "bar",
				Prompt: "baz",
			},
		},
	}
	got := c.collection.Items[0]
	if !reflect.DeepEqual(want, got) {
		t.Error("Should be able to attach an item to a collection")
		t.Errorf("Wanted %+v, got %+v", want, got)
	}

}

func TestLink(t *testing.T) {
	// Should not be able to add link to an unknown type
	href := url.URL{Host: "example.com", Scheme: "http"}
	linkOpt := NewLink(href, "foo", "bar", "baz", "qux")
	err := func(opt Option) error {
		return opt(struct{}{})
	}(linkOpt)
	if err != ErrTypeUnknown {
		t.Error("Should not be able to attach link to an unknown type")
		t.Errorf("Wanted %v, got %v", ErrTypeUnknown, err)
	}

	// Should be able to add link to an item
	href.Host = "example.org"
	itemOpt, err := NewItem(href, linkOpt)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Should be able to add link to a collection
	c, err := NewCollection(linkOpt, itemOpt)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	want := link{
		Href:   "http://example.com",
		Rel:    "foo",
		Name:   "bar",
		Render: "baz",
		Prompt: "qux",
	}
	got := c.collection.Links[0]
	if !reflect.DeepEqual(want, got) {
		t.Error("Should be able to attach a link to a collection")
		t.Errorf("Wanted %+v, got %+v", want, got)
	}
	got = c.collection.Items[0].Links[0]
	if !reflect.DeepEqual(want, got) {
		t.Error("Should be able to attach a link to an item")
		t.Errorf("Wanted %+v, got %+v", want, got)
	}

}

func TestQuery(t *testing.T) {
	// Should not be able to attach incorrect option to query
	errOpt := NewError("foo", "bar", "baz")
	_, err := NewQuery(url.URL{}, "", "", "", errOpt)
	if err != ErrTypeUnknown {
		t.Error("Should not be able to attach incorrect option to query")
		t.Errorf("Wanted %v, got %v", ErrTypeUnknown, err)
	}

	// Should not be able to attach query to unknown type
	href := url.URL{Host: "example.com", Scheme: "http"}
	queryOpt, err := NewQuery(href, "foo", "bar", "baz")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	err = func(opt Option) error {
		return opt(struct{}{})
	}(queryOpt)
	if err != ErrTypeUnknown {
		t.Error("Should not be able to attach query to unknown type")
		t.Errorf("Wanted %v, got %v", ErrTypeUnknown, err)
	}

	// Should be able to attach a query to a collection
	datumOpt := NewDatum("foo", "bar", "baz")
	queryOpt, err = NewQuery(href, "foo", "bar", "baz", datumOpt)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	c, err := NewCollection(queryOpt)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	want := query{
		Href:   "http://example.com",
		Rel:    "foo",
		Name:   "bar",
		Prompt: "baz",
		Data: []datum{
			{
				Name:   "foo",
				Value:  "bar",
				Prompt: "baz",
			},
		},
	}
	got := c.collection.Queries[0]
	if !reflect.DeepEqual(want, got) {
		t.Error("Should be able to attach a template to a collection")
		t.Errorf("Wanted %+v, got %+v", want, got)
	}
}
