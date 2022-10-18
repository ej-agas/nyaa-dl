package domain

import "strings"

type Filter = string

const (
	NoFilter    Filter = "0"
	NoRemakes   Filter = "1"
	TrustedOnly Filter = "2"
)

var Filters = map[string]Filter{
	"no-filter":    NoFilter,
	"no-remakes":   NoRemakes,
	"trusted-only": TrustedOnly,
}

func FilterFromString(s string, f map[string]Filter) (Filter, bool) {
	match, ok := f[strings.ToLower(s)]

	return match, ok
}
