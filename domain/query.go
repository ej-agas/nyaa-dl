package domain

import (
	"net/url"
)

type Query struct {
	filter   filter
	category Category
	search   string
}

func NewQuery(f filter, c Category, s string) *Query {
	return &Query{f, c, s}
}

func (q *Query) UrlEncode() string {
	params := url.Values{}
	params.Add("page", "rss")
	params.Add("f", q.filter.String())
	params.Add("c", string(q.category))
	params.Add("q", string(q.search))

	return params.Encode()
}
