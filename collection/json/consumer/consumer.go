package consumer

import (
	"encoding/json"

	"github.com/skriptble/hyper/collection/json"
)

// Collection represents a Collection+JSON document.
type Collection interface {
	Query(rels ...string) []Query
}

type index map[string][]int

// NewCollection converts a slice of bytes into a Collection.
func NewCollection(b []byte) (Collection, error) {
	var wrapper struct {
		collection `json:"collection"`
	}
	err := json.Unmarshal(b, &wrapper)
	if err != nil {
		return nil, err
	}
	c := wrapper.collection
	c.links = make(index)
	c.queries = make(index)
	// Build the indexes
	for idx, q := range c.Queries {
		c.queries[q.Rel] = append(c.queries[q.Rel], idx)
		c.queries[q.Name] = append(c.queries[q.Name], idx)
	}

	return c, nil
}

type collection struct {
	Version  cj.Version `json:"version"`
	Href     string     `json:"href"`
	Links    []link     `json:"links"`
	Items    []item     `json:"items"`
	Queries  []query    `json:"queries"`
	Template template   `json:"template"`
	Error    cjError    `json:"error"`

	// Indexes for the Queries and Links slices
	queries index
	links   index
}

type link struct {
	Href   string `json:"href"`
	Rel    string `json:"rel"`
	Name   string `json:"name,omitempty"`
	Render string `json:"render,omitempty"`
	Prompt string `json:"prompt,omitempty"`
}

type item struct {
	Href  string  `json:"href"`
	Data  []datum `json:"data,omitempty"`
	Links []link  `json:"links,omitempty"`
}

type query struct {
	Href      string  `json:"href"`
	Rel       string  `json:"rel"`
	Name      string  `json:"name,omitempty"`
	PromptStr string  `json:"prompt,omitempty"`
	Data      []datum `json:"data,omitempty"`
}

type template struct {
	Data []datum `json:"data"`
}

type datum struct {
	Name   string `json:"name"`
	Value  string `json:"value,omitempty"`
	Prompt string `json:"prompt,omitempty"`
}

type cjError struct {
	TTitle  string `json:"title,omitempty"`
	Code    string `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}
