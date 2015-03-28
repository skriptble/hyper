package producer

import (
	"encoding/json"
	"errors"

	"github.com/skriptble/hyper/collection/json"
)

// ErrTypeUnknown is returned when an option is passed into a New function for a
// type that does match the types it knows how to configure.
var ErrTypeUnknown = errors.New("producer: option given as argument for mismatching type")

// Option is a configuration option that can be passed into New functions.
// Options are safe to reuse in multiple invocations of New functions. If the
// Option is passed into a New function for a type it does not support it will
// return an ErrTypeUnknown error.
type Option func(interface{}) error

// Collection is created from a call to NewCollection and is the final
// representation of a Collection. It can be directly marshaled into JSON via
// the MarshalJSON method.
type Collection struct {
	collection collection
}

// encoding/json Marshaler implementation
func (c Collection) MarshalJSON() ([]byte, error) {
	document := struct {
		collection `json:"collection"`
	}{
		collection: c.collection,
	}
	return json.Marshal(document)
}

type collection struct {
	Version  cj.Version `json:"version"`
	Href     string     `json:"href,omitempty"`
	Links    []link     `json:"links,omitempty"`
	Items    []item     `json:"items,omitempty"`
	Queries  []query    `json:"queries,omitempty"`
	Template *template  `json:"template,omitempty"`
	Error    *cjError   `json:"error,omitempty"`
}

func NewCollection(opts ...Option) (Collection, error) {
	c := new(collection)
	c.Version = cj.V1
	for _, opt := range opts {
		err := opt(c)
		if err != nil {
			return Collection{}, err
		}
	}

	return Collection{*c}, nil
}

type link struct {
}

type item struct {
}

type query struct {
}

type template struct {
	Data []datum `json:"data"`
}

func NewTemplate(opts ...Option) (Option, error) {
	t := new(template)
	for _, opt := range opts {
		err := opt(t)
		if err != nil {
			return nil, err
		}
	}

	return func(i interface{}) error {
		c, ok := i.(*collection)
		if !ok {
			return ErrTypeUnknown
		}

		c.Template = t
		return nil
	}, nil
}

type datum struct {
	Name   string `json:"name"`
	Value  string `json:"value,omitempty"`
	Prompt string `json:"prompt,omitempty"`
}

func NewDatum(name, value, prompt string) Option {
	d := datum{
		Name:   name,
		Value:  value,
		Prompt: prompt,
	}
	return func(i interface{}) error {
		switch t := i.(type) {
		case *template:
			t.Data = append(t.Data, d)
		default:
			return ErrTypeUnknown
		}
		return nil
	}
}

type cjError struct {
	Title   string `json:"title,omitempty"`
	Code    string `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

func NewError(title, code, message string) Option {
	cjErr := cjError{
		Title:   title,
		Code:    code,
		Message: message,
	}
	return func(i interface{}) error {
		c, ok := i.(*collection)
		if !ok {
			return ErrTypeUnknown
		}

		c.Error = &cjErr
		return nil
	}
}
