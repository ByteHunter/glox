package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
)

func main() {
	flag.Parse()

	if flag.NArg() > 1 {
		fmt.Println("Error: Too many arguments!")
		os.Exit(64)
	}

	if flag.NArg() == 1 {
		runFile(flag.Arg(0))
	} else {
		runInteractive()
	}
}

func runFile(filename string) error {
	return errors.New("Not implemented")
}

func runInteractive() error {
	return errors.New("Not implemented")
}

func run(source string) error {
	return errors.New("Not implemented")
}
