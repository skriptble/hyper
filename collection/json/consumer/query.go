package consumer

import "strings"

// Query represents a Collection+JSON query.
type Query interface {
	// Set sets the key to the value, overwrites all values present.
	Set(key string, value string) Query
	// Add adds the value to the key, appends to all values present.
	Add(key string, value string) Query
	// Prompt returns the prompt value from the query
	Prompt() string
	URI() string
}

func (c collection) Query(rels ...string) []Query {
	queries := make([]Query, 0)
	tmpIndex := make(index)
	set := make(map[int]struct{}, 0)
	for k, v := range c.queries {
		tmpIndex[k] = v
	}

	for _, rel := range rels {
		for key := range tmpIndex {
			if !strings.Contains(key, rel) {
				delete(tmpIndex, key)
			}
		}
	}

	// flatten and dedupe
	for _, idx := range tmpIndex {
		for _, val := range idx {
			set[val] = struct{}{}
		}
	}

	for idx := range set {
		queries = append(queries, c.Queries[idx])
	}
	return queries
}

func (q query) Add(key string, value string) Query {
	return nil
}
func (q query) Set(key string, value string) Query {
	return nil
}
func (q query) Prompt() string {
	return ""
}
func (q query) URI() string {
	return ""
}
