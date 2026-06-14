package ptrtonew_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"

	"github.com/nunnatsa/ptrtonew"
)

func TestAnalyzer(t *testing.T) {
	a := ptrtonew.NewAnalyzer()

	testCases := map[string]string{
		"test k8s.io/utils/ptr":     "./ptr",
		"test k8s.io/utils/pointer": "./pointer",
	}

	for testName, testDir := range testCases {
		t.Run(testName, func(t *testing.T) {
			analysistest.RunWithSuggestedFixes(t, analysistest.TestData()+"/src/a", a, testDir)
		})
	}
}
