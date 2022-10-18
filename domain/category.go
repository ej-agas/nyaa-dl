package domain

import "strings"

type Category = string

const (
	AllCategories Category = "0_0"
)

var Categories = map[string]Category{
	"all": AllCategories,
}

func CategoryFromString(s string, c map[string]Category) (Category, bool) {
	match, ok := c[strings.ToLower(s)]

	return match, ok
}
