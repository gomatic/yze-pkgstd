// Command yze-go-pkgstd runs the pkgstd analyzer as a standalone go/analysis
// checker (text, -json, and -fix output, and as a `go vet -vettool`).
package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"

	pkgstd "github.com/gomatic/yze-pkgstd"
)

// run is the analysis entry point, indirected so the binary's wiring is testable
// without invoking the real driver (which loads packages and exits the process).
var run = singlechecker.Main

func main() { run(pkgstd.Analyzer) }
