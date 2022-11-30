package domain

import (
	"sort"
	"strings"
)

type Sorter interface {
	sort.Interface
	SetItemCollection(ItemCollection) Sorter
}

func SortItems(items ItemCollection, sortBy string, orderBy string) {
	switch strings.ToLower(sortBy) {
	case "size":
		if strings.ToLower(orderBy) == "asc" {
			sortDesc(SortBySize{items})
			break
		}

		sortAsc(SortBySize{items})
	case "seeders":
		if strings.ToLower(orderBy) == "asc" {
			sortDesc(SortBySeeders{items})
			break
		}

		sortAsc(SortBySeeders{items})
	case "leechers":
		if strings.ToLower(orderBy) == "asc" {
			sortDesc(SortByLeechers{items})
			break
		}

		sortAsc(SortByLeechers{items})
	case "downloads":
		if strings.ToLower(orderBy) == "asc" {
			sortDesc(SortByDownloads{items})
			break
		}

		sortAsc(SortByDownloads{items})
	case "date":
		if strings.ToLower(orderBy) == "asc" {
			sortDesc(SortByDate{items})
			break
		}

		sortAsc(SortByDate{items})
	}
}

func sortAsc(items sort.Interface) {
	sort.Sort(items)
}

func sortDesc(items sort.Interface) {
	sort.Sort(sort.Reverse(items))
}
