package greet

// Config is the command configuration.
type Config struct{ Name string }

// Run is the entry point.
func Run(cfg Config) string { return cfg.Name }
