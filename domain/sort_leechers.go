package domain

type SortByLeechers struct {
	Items
}

func (s SortByLeechers) Less(i, j int) bool {
	return s.Items.Items[i].Leechers > s.Items.Items[j].Leechers
}
