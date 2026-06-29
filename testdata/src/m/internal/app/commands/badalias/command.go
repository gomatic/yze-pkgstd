package badalias

import (
	greet "m/internal/domain/greet" // want `domain.*alias`
)

const x = 1

func Command() greet.Config { return greet.Config{} }
