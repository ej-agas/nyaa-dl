package domain

type SortByDownloads struct {
	ItemCollection
}

func (s SortByDownloads) SetItemCollection(i ItemCollection) Sorter {
	s.ItemCollection = i
	return s
}

func (s SortByDownloads) Less(i, j int) bool {
	return s.ItemCollection.Items[i].Downloads > s.ItemCollection.Items[j].Downloads
}
