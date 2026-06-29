package noconst

import domain "m/internal/domain/greet"

var cfg domain.Config // want `const block`

func Command() domain.Config { return cfg }
