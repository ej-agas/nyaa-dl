package domain

type SortBySeeders struct {
	Items
}

func (s SortBySeeders) Less(i, j int) bool {
	return s.Items.Items[i].Seeders > s.Items.Items[j].Seeders
}
