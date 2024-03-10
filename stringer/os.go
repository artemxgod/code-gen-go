package stringer

//go:generate stringer -type=OS
type OS uint8

const (
	Windows OS = iota
	Darwin
	Linux
)
