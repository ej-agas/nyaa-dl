package domain

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestItem(t *testing.T) {
	item := item{
		Title:       "[SubsPlease] Bocchi the Rock! - 05 (1080p) [192004DE].mkv",
		Link:        "https://nyaa.si/download/1599201.torrent",
		Url:         "https://nyaa.si/view/1599201",
		PublishDate: "Sat, 05 Nov 2022 16:33:33 -0000",
		Seeders:     0x414,
		Leechers:    0x31,
		Downloads:   0xdea,
		InfoHash:    "ef0d35adc2c45d6c226de092249eab36c43e5630",
		CategoryID:  "1_2",
		Category:    "Anime - English-translated",
		Size:        "1.3 GiB",
		Trusted:     "Yes",
		Remake:      "No",
		Description: "<a href=\"https://nyaa.si/view/1599201\">#1599201 | [SubsPlease] Bocchi the Rock! - 05 (1080p) [192004DE].mkv</a> | 1.3 GiB | Anime - English-translated | ef0d35adc2c45d6c226de092249eab36c43e5630",
	}

	assert.Equal(t, "Sun, 06 Nov 2022, 12:33:33 AM PST", item.PublishDateLocalTz())

	size, err := item.SizeInBytes()

	if err != nil {
		fmt.Fprintf(os.Stderr, "test failed: %s\n", err)
	}

	assert.Equal(t, 1.3958643712e+09, size)
}
