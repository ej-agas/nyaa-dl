package domain

import "strings"

type filter struct {
	value string
}

func (f filter) String() string {
	return f.value
}

var filters = map[string]filter{
	"no-filter":    {value: "0"},
	"no-remakes":   {value: "1"},
	"trusted-only": {value: "2"},
}

func NewFilter(s string) (filter, bool) {
	match, ok := filters[strings.ToLower(s)]

	return match, ok
}

func NoFilter() filter {
	return filter{value: "0"}
}
