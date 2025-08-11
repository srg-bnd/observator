// Custom Analyzers
package custom

import (
	"golang.org/x/tools/go/analysis"
)

// An analyzer that prohibits the use of a direct call to `os.Exit` in the main function of the main package
func NewNoOsExitInMain() *analysis.Analyzer {
	return &analysis.Analyzer{
		Name: "noosexitinmain",
		Doc:  "forbids direct calls to os.Exit in main.main function",
		Run:  runNoOsExitInMain,
	}
}

func runNoOsExitInMain(pass *analysis.Pass) (interface{}, error) {
	return nil, nil
}
