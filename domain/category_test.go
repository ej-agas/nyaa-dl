package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCategoryFromString(t *testing.T) {
	match1, ok1 := CategoryFromString("some-string", Categories)

	assert.Equal(t, "", match1)
	assert.Equal(t, false, ok1)

	match2, ok2 := CategoryFromString("all", Categories)

	assert.Equal(t, AllCategories, match2)
	assert.Equal(t, true, ok2)
}
