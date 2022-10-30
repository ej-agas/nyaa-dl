package domain

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewQuery(t *testing.T) {
	assert := assert.New(t)

	filter, ok := NewFilter("no-remakes")

	if ok != true {
		filter = NoFilter()
	}

	category := AllCategories
	search := "Foo"

	q := NewQuery(filter, category, search)

	assert.Equal(filter, q.filter)
	assert.Equal(category, q.category)
	assert.Equal(search, q.search)
}

func TestUrlEncode(t *testing.T) {
	params := url.Values{}
	params.Add("page", "rss")
	params.Add("f", NoFilter().value)
	params.Add("c", AllCategories)
	params.Add("q", "Foo")
	expected := params.Encode()

	q := NewQuery(NoFilter(), AllCategories, "Foo")
	actual := q.UrlEncode()

	assert.Equal(t, expected, actual)
}
