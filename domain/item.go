package domain

import (
	"time"
)

type item struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Url         string `xml:"guid"`
	PublishDate string `xml:"pubDate"`
	Seeders     uint   `xml:"seeders"`
	Leechers    uint   `xml:"leechers"`
	Downloads   uint   `xml:"downloads"`
	InfoHash    string `xml:"infoHash"`
	CategoryID  string `xml:"categoryId"`
	Category    string `xml:"category"`
	Size        string `xml:"size"`
	Trusted     string `xml:"trusted"`
	Remake      string `xml:"remake"`
	Description string `xml:"description"`
}

func (i *item) PublishDateLocalTz() (string, error) {
	time, err := time.Parse("Mon, 02 Jan 2006 15:04:05 -0000", i.PublishDate)

	if err != nil {
		return "", err
	}

	return time.Local().Local().Format("Mon, 02 Jan 2006, 3:04:05 AM MST"), nil
}

func (i *item) SizeInBytes() (float64, error) {
	return ConvertToBytes(i.Size)
}
