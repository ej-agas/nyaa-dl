package domain

type SortByDownloads struct {
	Items
}

func (s SortByDownloads) Less(i, j int) bool {
	return s.Items.Items[i].Downloads > s.Items.Items[j].Downloads
}
