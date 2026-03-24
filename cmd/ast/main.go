package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	os.Exit(RunMain())
}

func RunMain() int {
	flag.Parse()
	outputDirectory := flag.Arg(0)
	if len(outputDirectory) == 0 {
		fmt.Println("Usage go run cmd/ast/main.go <output-path>")
		return 1
	}

	file, err := getFile(outputDirectory, "Expression")
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return 1
	}
	defer file.Close()

	contents, err := generateContent("Expression")
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return 1
	}

	_, err = file.WriteString(contents)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return 1
	}

	return 0
}

func getFile(outputDir, baseName string) (*os.File, error) {
	path := fmt.Sprintf("%s/%s.go", outputDir, strings.ToLower(baseName))

	file, err := os.Create(path)

	if err != nil {
		return nil, err
	}

	return file, nil
}

func generateContent(baseName string) (string, error) {
	// Starting the file content
	contents := fmt.Sprintf("package %s\n\n", strings.ToLower(baseName))
	contents += fmt.Sprintf("type %s interface {\n", baseName)
	contents += "}\n\n"

	return contents, nil
}
