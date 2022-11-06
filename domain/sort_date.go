package domain

import "time"

type SortByDate struct {
	Items
}

func (s SortByDate) Less(i, j int) bool {
	dateTimeFormat := "Mon, 02 Jan 2006 15:04:05 -0000"

	iTime, _ := time.Parse(dateTimeFormat, s.Items.Items[i].PublishDate)
	jTime, _ := time.Parse(dateTimeFormat, s.Items.Items[j].PublishDate)

	return iTime.After(jTime)
}
