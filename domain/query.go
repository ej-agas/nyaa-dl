package domain

import (
	"net/url"
)

type Query struct {
	filter   Filter
	category Category
	search   string
}

func NewQuery(f Filter, c Category, s string) *Query {
	return &Query{f, c, s}
}

func (q *Query) UrlEncode() string {
	params := url.Values{}
	params.Add("page", "rss")
	params.Add("f", q.filter)
	params.Add("c", q.category)
	params.Add("q", q.search)

	return params.Encode()
}
