package ptrtonew_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"

	"ptrtonew"
)

func TestAnalyzer(t *testing.T) {
	a := ptrtonew.NewAnalyzer()

	analysistest.RunWithSuggestedFixes(t, analysistest.TestData()+"/src/a", a)
}
