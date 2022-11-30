package domain

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortLeechers(t *testing.T) {
	itemCollection := ItemCollection{
		[]item{
			{Leechers: 1},
			{Leechers: 123},
			{Leechers: 321},
			{Leechers: 1100},
			{Leechers: 23},
			{Leechers: 41},
			{Leechers: 54},
			{Leechers: 69},
		},
	}

	sorter := SortByLeechers{ItemCollection: itemCollection}
	assert.False(t, sort.IsSorted(sorter))
	sort.Sort(sorter)
	assert.True(t, sort.IsSorted(sorter))
}
