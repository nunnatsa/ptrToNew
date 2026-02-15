package main

import (
	"fmt"
	"os"
	"runtime"

	"golang.org/x/tools/go/analysis/singlechecker"

	"ptrtonew"
	"ptrtonew/myversion"
)

func main() {
	if len(os.Args) == 2 && os.Args[1] == "version" {
		fmt.Printf("ptrToNew version: %s\n", myversion.Version)
		fmt.Printf("go version:       %s\n", runtime.Version())
		os.Exit(0)
	}

	singlechecker.Main(ptrtonew.NewAnalyzer())
}
