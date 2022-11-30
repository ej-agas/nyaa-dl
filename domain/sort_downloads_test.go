package domain

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortDownloads(t *testing.T) {
	itemCollection := ItemCollection{
		[]item{
			{Downloads: 1},
			{Downloads: 123},
			{Downloads: 321},
			{Downloads: 1100},
			{Downloads: 23},
			{Downloads: 41},
			{Downloads: 54},
			{Downloads: 69},
		},
	}

	sorter := SortByDownloads{ItemCollection: itemCollection}
	assert.False(t, sort.IsSorted(sorter))
	sort.Sort(sorter)
	assert.True(t, sort.IsSorted(sorter))
}
