package domain

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortDate(t *testing.T) {
	itemCollection := ItemCollection{
		[]item{
			{PublishDate: "Tue, 11 Oct 2022 16:01:55 -0000"},
			{PublishDate: "Tue, 22 Nov 2022 17:01:37 -0000"},
			{PublishDate: "Tue, 08 Nov 2022 17:03:24 -0000"},
			{PublishDate: "Tue, 29 Nov 2022 17:01:33 -0000"},
			{PublishDate: "Tue, 25 Oct 2022 16:01:50 -0000"},
			{PublishDate: "Tue, 01 Nov 2022 16:02:14 -0000"},
			{PublishDate: "Tue, 15 Nov 2022 17:03:39 -0000"},
			{PublishDate: "Tue, 18 Oct 2022 16:01:45 -0000"},
		},
	}

	sorter := SortByDate{ItemCollection: itemCollection}
	assert.False(t, sort.IsSorted(sorter))
	sort.Sort(sorter)
	assert.True(t, sort.IsSorted(sorter))
}
