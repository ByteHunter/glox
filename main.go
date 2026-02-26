package main

import (
	"bufio"
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

func runFile(fileName string) error {
	fileData, err := os.ReadFile(fileName)
	if err != nil {
		return err
	}

	source := string(fileData)
	run(source)

	return nil
}

func runInteractive() error {
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Printf("> ")
		line, prefix, err := reader.ReadLine()

		if err != nil {
			return err
		}

		if prefix {
			fmt.Println("Prompt size was too large!")
		}

		source := string(line)
		if len(source) == 0 {
			break
		}

		run(source + "\n")
	}

	return nil
}

func run(source string) error {
	if len(source) == 0 {
		return nil
	}

	scanner := NewScanner(source)
	scanner.scanTokens()

	return nil
}

func loxError(line int, message string) {
	loxReport(line, "", message)
}

func loxReport(line int, where string, message string) {
	fmt.Fprintf(os.Stdout, "[line %d] Error %s: %s\n", line, where, message)
}
