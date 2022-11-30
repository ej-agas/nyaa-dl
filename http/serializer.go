package http

import (
	"encoding/xml"

	"github.com/ej-agas/nyaa-dl/domain"
)

func SerializeToItems(b []byte) domain.ItemCollection {
	var items domain.ItemCollection

	xml.Unmarshal(b, &items)

	return items
}
