package domain

import "time"

type SortByDate struct {
	ItemCollection
}

func (s SortByDate) SetItemCollection(i ItemCollection) Sorter {
	s.ItemCollection = i
	return s
}

func (s SortByDate) Less(i, j int) bool {
	dateTimeFormat := "Mon, 02 Jan 2006 15:04:05 -0000"

	iTime, _ := time.Parse(dateTimeFormat, s.ItemCollection.Items[i].PublishDate)
	jTime, _ := time.Parse(dateTimeFormat, s.ItemCollection.Items[j].PublishDate)

	return iTime.After(jTime)
}
