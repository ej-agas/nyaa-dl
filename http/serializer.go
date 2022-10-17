package http

import (
	"encoding/xml"

	"github.com/ej-agas/nyaa-dl/domain"
)

func SerializeToItems(b []byte) domain.Items {
	var items domain.Items

	xml.Unmarshal(b, &items)

	return items
}
