package main

import (
	"halstead/cli"
	"halstead/halstead"
)

func main() {

	filepath := cli.ParseCLI()

	metrics, err := halstead.AnalyzeSourceCode(filepath)
	if err != nil {
		panic(err)
	}

	metrics.Print()
}
