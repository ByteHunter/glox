package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	os.Exit(RunMain())
}

func RunMain() int {
	fmt.Println("AST Generator")

	flag.Parse()
	outputDirectory := flag.Arg(0)

	if len(outputDirectory) == 0 {
		fmt.Println("Usage go run tool/main.go <output-path>")
		return 1
	}

	return 0
}
