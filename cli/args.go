package cli

import (
	"flag"
	"fmt"
	"os"
)

//cli input from args
func ParseCLI() string {
	var filepath string
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n", os.Args[0])
		fmt.Fprintf(flag.CommandLine.Output(), "  %s -filepath <filepath>\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.StringVar(&filepath, "filepath", "", "Filepath to the source code")
	flag.Parse()

	if filepath == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	return filepath
}
