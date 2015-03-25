package cj

import "net/url"

const MediaType = "application/vnd.collection+json"
const V1 Version = "1.0"

type Version string

type Collection struct {
	collection collection
}

type collection struct {
	Version Version `json:"version"`
	Href    url.URL `json:"href"`
}
