package domain

type SortBySeeders struct {
	ItemCollection
}

func (s SortBySeeders) SetItemCollection(i ItemCollection) Sorter {
	s.ItemCollection = i
	return s
}

func (s SortBySeeders) Less(i, j int) bool {
	return s.ItemCollection.Items[i].Seeders > s.ItemCollection.Items[j].Seeders
}
