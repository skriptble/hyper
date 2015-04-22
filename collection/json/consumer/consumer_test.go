package consumer

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/skriptble/hyper/collection/json"
)

func TestCollection(t *testing.T) {
	// Should be able to unmarshal JSON into Collection
	doc := []byte(`{"collection":{"version":"1.0","href":"http://example.com"}}`)
	w := new(wrapper)
	err := json.Unmarshal(doc, w)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	got := w.C
	want := collection{
		Version: cj.V1,
		Href:    "http://example.com",
	}
	if !reflect.DeepEqual(want, got) {
		t.Error("Should be able to unmarshal JSON into Collection")
		t.Errorf("Wanted %+v, got %+v", want, got)
	}
}
