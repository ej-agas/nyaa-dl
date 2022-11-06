package domain

import (
	"sort"
	"strings"
)

func SortItems(items Items, sortBy string, orderBy string) {
	switch strings.ToLower(sortBy) {
	case "size":
		if strings.ToLower(orderBy) == "asc" {
			sortDesc(SortBySize{Items: items})
			break
		}

		sort.Sort(SortBySize{Items: items})
	case "seeders":
		if strings.ToLower(orderBy) == "asc" {
			sortDesc(SortBySeeders{Items: items})
			break
		}

		sort.Sort(SortBySeeders{Items: items})
	case "leechers":
		if strings.ToLower(orderBy) == "asc" {
			sortDesc(SortByLeechers{Items: items})
			break
		}

		sort.Sort(SortByLeechers{Items: items})
	case "downloads":
		if strings.ToLower(orderBy) == "asc" {
			sortDesc(SortByDownloads{Items: items})
			break
		}

		sort.Sort(SortByDownloads{Items: items})
	case "date":
		if strings.ToLower(orderBy) == "asc" {
			sortDesc(SortByDate{Items: items})
			break
		}

		sort.Sort(SortByDate{Items: items})
	}
}

func sortDesc(items sort.Interface) {
	sort.Sort(sort.Reverse(items))
}
