package domain

type SortByLeechers struct {
	ItemCollection
}

func (s SortByLeechers) SetItemCollection(i ItemCollection) Sorter {
	s.ItemCollection = i
	return s
}

func (s SortByLeechers) Less(i, j int) bool {
	return s.ItemCollection.Items[i].Leechers > s.ItemCollection.Items[j].Leechers
}
