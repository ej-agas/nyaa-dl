package domain

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortSeeders(t *testing.T) {
	itemCollection := ItemCollection{
		[]item{
			{Seeders: 1},
			{Seeders: 123},
			{Seeders: 321},
			{Seeders: 1100},
			{Seeders: 23},
			{Seeders: 41},
			{Seeders: 54},
			{Seeders: 69},
		},
	}

	sorter := SortBySeeders{ItemCollection: itemCollection}
	assert.False(t, sort.IsSorted(sorter))
	sort.Sort(sorter)
	assert.True(t, sort.IsSorted(sorter))
}
