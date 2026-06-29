// Command yze-go-pkgstd runs the pkgstd analyzer as a standalone go/analysis
// checker (text, -json, and -fix output, and as a `go vet -vettool`).
package main

import (
	pkgstd "github.com/gomatic/yze-go-pkgstd"
	"golang.org/x/tools/go/analysis/singlechecker"
)

// run is the analysis entry point, indirected so the binary's wiring is testable
// without invoking the real driver (which loads packages and exits the process).
var run = singlechecker.Main

func main() { run(pkgstd.Analyzer) }
