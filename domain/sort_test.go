package domain

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSort(t *testing.T) {
	itemCollection := ItemCollection{
		[]item{
			{PublishDate: "Tue, 11 Oct 2022 16:01:55 -0000", Downloads: 1, Seeders: 1, Leechers: 1, Size: "1 MiB"},
			{PublishDate: "Tue, 22 Nov 2022 17:01:37 -0000", Downloads: 123, Seeders: 123, Leechers: 123, Size: "20 MiB"},
			{PublishDate: "Tue, 08 Nov 2022 17:03:24 -0000", Downloads: 321, Seeders: 321, Leechers: 321, Size: "5 KiB"},
			{PublishDate: "Tue, 29 Nov 2022 17:01:33 -0000", Downloads: 1100, Seeders: 1100, Leechers: 1100, Size: "255 Bytes"},
			{PublishDate: "Tue, 25 Oct 2022 16:01:50 -0000", Downloads: 23, Seeders: 23, Leechers: 23, Size: "16 TiB"},
			{PublishDate: "Tue, 01 Nov 2022 16:02:14 -0000", Downloads: 41, Seeders: 41, Leechers: 41, Size: "25 GiB"},
			{PublishDate: "Tue, 15 Nov 2022 17:03:39 -0000", Downloads: 54, Seeders: 54, Leechers: 54, Size: "2 GiB"},
			{PublishDate: "Tue, 18 Oct 2022 16:01:45 -0000", Downloads: 69, Seeders: 69, Leechers: 69, Size: "551 MiB"},
		},
	}

	scenarios := []struct {
		testName string
		input    Sorter
		sortBy   string
		orderBy  string
	}{
		{
			testName: "Sort by Size",
			input:    SortBySize{itemCollection},
			sortBy:   "size",
			orderBy:  "asc",
		},
		{
			testName: "Sort by Seeders",
			input:    SortBySeeders{itemCollection},
			sortBy:   "seeders",
			orderBy:  "asc",
		},
		{
			testName: "Sort by Leechers",
			input:    SortByLeechers{itemCollection},
			sortBy:   "leechers",
			orderBy:  "asc",
		},
		{
			testName: "Sort by Downloads",
			input:    SortByDownloads{itemCollection},
			sortBy:   "downloads",
			orderBy:  "asc",
		},
		{
			testName: "Sort by Date",
			input:    SortByDate{itemCollection},
			sortBy:   "date",
			orderBy:  "asc",
		},
	}

	for _, scenario := range scenarios {
		t.Run(scenario.testName, func(t *testing.T) {
			SortItems(itemCollection, scenario.sortBy, scenario.sortBy)
			assert.True(t, sort.IsSorted(scenario.input))
		})
	}
}
