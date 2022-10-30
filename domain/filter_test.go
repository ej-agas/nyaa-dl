package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilterStringFunc(t *testing.T) {
	filter := filter{value: "0"}

	assert.Equal(t, "0", filter.String())
}

func TestNewFilter(t *testing.T) {
	badMatch, notOk := NewFilter("some-string")

	assert.Equal(t, filter{value: ""}, badMatch)
	assert.Equal(t, false, notOk)

	match, ok := NewFilter("no-filter")

	assert.Equal(t, filter{value: "0"}, match)
	assert.Equal(t, true, ok)
}

func TestNoFilter(t *testing.T) {
	noFilter := NoFilter()

	assert.Equal(t, filter{value: "0"}, noFilter)
}
