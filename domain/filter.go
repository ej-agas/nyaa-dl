package domain

type Filter = string

const (
	NoFilter    Filter = "0"
	NoRemakes   Filter = "1"
	TrustedOnly Filter = "2"
)
