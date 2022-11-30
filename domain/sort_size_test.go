package domain

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortSize(t *testing.T) {
	itemCollection := ItemCollection{
		[]item{
			{Size: "1 MiB"},
			{Size: "20 MiB"},
			{Size: "5 KiB"},
			{Size: "255 Bytes"},
			{Size: "16 TiB"},
			{Size: "25 GiB"},
			{Size: "2 GiB"},
			{Size: "551 MiB"},
		},
	}

	sorter := SortBySize{ItemCollection: itemCollection}
	assert.False(t, sort.IsSorted(sorter))
	sort.Sort(sorter)
	assert.True(t, sort.IsSorted(sorter))
}
