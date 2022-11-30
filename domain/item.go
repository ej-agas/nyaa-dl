package domain

import (
	"regexp"
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

func (i *item) PublishDateLocalTz() string {
	time, err := time.Parse("Mon, 02 Jan 2006 15:04:05 -0000", i.PublishDate)

	if err != nil {
		return ""
	}

	return time.Local().Local().Format("Mon, 02 Jan 2006, 3:04:05 AM MST")
}

func (i *item) SizeInBytes() (float64, error) {
	return ConvertToBytes(i.Size)
}

func (i item) Id() string {
	r := regexp.MustCompile(`\d+`)
	results := r.FindAllString(i.Link, -1)

	if len(results) == 0 {
		return ""
	}

	return results[0]
}
