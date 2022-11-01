package domain

import "time"

type Items struct {
	Items []item `xml:"channel>item"`
}

func (i Items) Len() int {
	return len(i.Items)
}

func (it Items) Swap(i, j int) {
	it.Items[i], it.Items[j] = it.Items[j], it.Items[i]
}

type SortBySize struct {
	Items
}

func (s SortBySize) Less(i, j int) bool {
	iSize, _ := s.Items.Items[i].SizeInBytes()
	jSize, _ := s.Items.Items[j].SizeInBytes()

	return iSize > jSize
}

type SortBySeeders struct {
	Items
}

func (s SortBySeeders) Less(i, j int) bool {
	return s.Items.Items[i].Seeders > s.Items.Items[j].Seeders
}

type SortByLeechers struct {
	Items
}

func (s SortByLeechers) Less(i, j int) bool {
	return s.Items.Items[i].Leechers > s.Items.Items[j].Leechers
}

type SortByDownloads struct {
	Items
}

func (s SortByDownloads) Less(i, j int) bool {
	return s.Items.Items[i].Downloads > s.Items.Items[j].Downloads
}

type SortByDate struct {
	Items
}

func (s SortByDate) Less(i, j int) bool {
	dateTimeFormat := "Mon, 02 Jan 2006 15:04:05 -0000"

	iTime, _ := time.Parse(dateTimeFormat, s.Items.Items[i].PublishDate)
	jTime, _ := time.Parse(dateTimeFormat, s.Items.Items[j].PublishDate)

	return iTime.After(jTime)
}
