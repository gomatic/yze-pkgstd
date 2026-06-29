// Package pkgstd provides a go/analysis analyzer enforcing the gomatic three-tier
// layout's per-package standards. For a command package
// (internal/app/commands/<cmd>): the first declaration is the const block, a
// Command() entry point exists, and the domain package is imported under the
// "domain" alias. (Cross-package correspondence is the layout analyzer's job.)
package pkgstd

import (
	"go/ast"
	"go/token"
	"strings"

	goyze "github.com/gomatic/go-yze"
	"golang.org/x/tools/go/analysis"
)

// Analyzer reports per-package violations of the three-tier command-package layout.
var Analyzer = &analysis.Analyzer{
	Name: "pkgstd",
	Doc:  "reports command packages that violate the gomatic three-tier package standards",
	Run:  run,
}

// Registration declares this analyzer to the yze framework.
var Registration = goyze.Registration{
	Name:       "pkgstd",
	Categories: []goyze.Category{"structure"},
	URL:        "https://docs.gomatic.dev/yze/go/pkgstd",
	Analyzer:   Analyzer,
}

// run checks command packages for the per-package layout standards.
func run(pass *analysis.Pass) (any, error) {
	if isCommandPackage(pass.Pkg.Path()) {
		checkConstFirst(pass)
		checkCommandFunc(pass)
		checkDomainAlias(pass)
	}
	return nil, nil
}

// isCommandPackage reports whether a package path is a command package.
func isCommandPackage(pkgPath string) bool {
	return strings.Contains(pkgPath, "/internal/app/commands/")
}

// checkConstFirst reports when the first non-import declaration is not a const
// block.
func checkConstFirst(pass *analysis.Pass) {
	for _, decl := range pass.Files[0].Decls {
		if isImportDecl(decl) {
			continue
		}
		if !isConstDecl(decl) {
			pass.Reportf(decl.Pos(), "command package: the first declaration must be the const block (name, usage, ...)")
		}
		return
	}
}

func isImportDecl(decl ast.Decl) bool {
	gen, ok := decl.(*ast.GenDecl)
	return ok && gen.Tok == token.IMPORT
}

func isConstDecl(decl ast.Decl) bool {
	gen, ok := decl.(*ast.GenDecl)
	return ok && gen.Tok == token.CONST
}

// checkCommandFunc reports when the package has no Command() entry point.
func checkCommandFunc(pass *analysis.Pass) {
	for _, file := range pass.Files {
		for _, decl := range file.Decls {
			if isCommandFunc(decl) {
				return
			}
		}
	}
	pass.Reportf(pass.Files[0].Name.Pos(), "command package: missing the Command() entry point")
}

func isCommandFunc(decl ast.Decl) bool {
	fn, ok := decl.(*ast.FuncDecl)
	return ok && fn.Recv == nil && fn.Name.Name == "Command"
}

// checkDomainAlias reports domain imports not aliased as "domain".
func checkDomainAlias(pass *analysis.Pass) {
	for _, file := range pass.Files {
		for _, imp := range file.Imports {
			checkDomainImport(pass, imp)
		}
	}
}

func checkDomainImport(pass *analysis.Pass, imp *ast.ImportSpec) {
	path := strings.Trim(imp.Path.Value, `"`)
	if !strings.Contains(path, "/internal/domain/") {
		return
	}
	if imp.Name == nil || imp.Name.Name != "domain" {
		pass.Reportf(imp.Pos(), "command package: import the domain package with the \"domain\" alias")
	}
}
