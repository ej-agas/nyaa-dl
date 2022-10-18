package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilterFromString(t *testing.T) {
	match, ok := FilterFromString("some-string", Filters)

	assert.Equal(t, "", match)
	assert.Equal(t, false, ok)

	match2, ok2 := FilterFromString("no-filter", Filters)

	assert.Equal(t, NoFilter, match2)
	assert.Equal(t, true, ok2)
}
