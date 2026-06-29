package greet

import (
	"strings"

	domain "m/internal/domain/greet"
)

const (
	name  = "greet"
	usage = "greet someone"
)

var cfg domain.Config

// Command is the entry point.
func Command() string {
	_ = cfg
	_ = usage
	return strings.ToUpper(name)
}
