package pkgstd_test

import (
	"testing"

	pkgstd "github.com/gomatic/yze-go-pkgstd"
	"github.com/stretchr/testify/assert"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestCommandPackageStandards(t *testing.T) {
	analysistest.Run(
		t, analysistest.TestData(), pkgstd.Analyzer,
		"m/internal/domain/greet",
		"m/internal/app/commands/greet",
		"m/internal/app/commands/badalias",
		"m/internal/app/commands/noconst",
		"m/internal/app/commands/nocmd",
	)
}

func TestRegistrationIsWellFormed(t *testing.T) {
	assert.NoError(t, pkgstd.Registration.Validate())
	assert.Equal(t, "yze/go/pkgstd", pkgstd.Registration.RuleID())
	assert.Same(t, pkgstd.Analyzer, pkgstd.Registration.Analyzer)
}
