package nocmd

import "testing"

// TestNocmd exists only to produce the in-package test variant (command.go plus
// a _test.go file); that variant duplicates the primary build, so primaryBuild
// skips it and the missing-Command diagnostic is reported exactly once.
func TestNocmd(t *testing.T) { _ = t }
