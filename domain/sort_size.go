package domain

type SortBySize struct {
	ItemCollection
}

func (s SortBySize) SetItemCollection(i ItemCollection) Sorter {
	s.ItemCollection = i
	return s
}

func (s SortBySize) Less(i, j int) bool {
	iSize, _ := s.ItemCollection.Items[i].SizeInBytes()
	jSize, _ := s.ItemCollection.Items[j].SizeInBytes()

	return iSize > jSize
}
