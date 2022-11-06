package domain

type SortBySize struct {
	Items
}

func (s SortBySize) Less(i, j int) bool {
	iSize, _ := s.Items.Items[i].SizeInBytes()
	jSize, _ := s.Items.Items[j].SizeInBytes()

	return iSize > jSize
}
