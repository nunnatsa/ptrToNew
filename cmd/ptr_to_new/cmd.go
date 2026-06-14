package main

import (
	_ "embed"
	"fmt"
	"os"
	"runtime"

	"golang.org/x/tools/go/analysis/singlechecker"

	"github.com/nunnatsa/ptrtonew"
)

//go:embed version.txt
var version string

func main() {
	if len(os.Args) == 2 && os.Args[1] == "version" {
		fmt.Printf("ptrToNew version: %s\n", version)
		fmt.Printf("go version:       %s\n", runtime.Version())
		os.Exit(0)
	}

	singlechecker.Main(ptrtonew.NewAnalyzer())
}
