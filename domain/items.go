package domain

import (
	"sort"
)

type Items struct {
	Items []item `xml:"channel>item"`
}

func (i *Items) Sort(sortBy *string, order *string) {
	if *sortBy == "size" {
		sortBySize(i.Items, order)
	}

	switch *sortBy {
	case "size":
		sortBySize(i.Items, order)
	case "seeders":
		sortBySeeders(i.Items, order)
	}
}

func sortBySize(i []item, order *string) {
	sort.Slice(i, func(a, b int) bool {
		aSize, _ := i[a].SizeInBytes()
		bSize, _ := i[b].SizeInBytes()

		if *order == "asc" {
			return aSize < bSize
		}

		return aSize > bSize
	})
}

func sortBySeeders(i []item, order *string) {
	sort.Slice(i, func(a, b int) bool {
		if *order == "asc" {
			return i[a].Seeders < i[b].Seeders
		}

		return i[a].Seeders > i[b].Seeders
	})
}
