package main

import (
	"flag"
	"fmt"
	"go/format"
	"os"
	"strings"
)

type FieldDefinition struct {
	key   string
	value string
}
type FieldList []FieldDefinition
type SubClassDefinition struct {
	name   string
	fields FieldList
}
type SubClassList []SubClassDefinition

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

	var classes = SubClassList{
		{
			"Binary",
			FieldList{
				{"left", "Expression"},
				{"operator", "token.Token"},
				{"right", "Expression"},
			},
		},
		{
			"Grouping",
			FieldList{
				{"expression", "Expression"},
			},
		},
	}

	contents, err := buildContent("Expression", classes)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return 1
	}

	doFormat := true
	if doFormat {
		formatted, err := format.Source([]byte(contents))
		if err != nil {
			fmt.Printf("Error: %s\n", err.Error())
			return 1
		}
		contents = string(formatted)
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

func generateContent(baseName string, classes SubClassList) (string, error) {
	// Starting the file content
	contents := fmt.Sprintf("package %s\n\n", strings.ToLower(baseName))
	contents += "import (\n"
	contents += "\"github.com/ByteHunter/glox/token\"\n"
	contents += ")\n"
	contents += fmt.Sprintf("type %s any\n\n", baseName)

	// Subclasses
	for _, subClass := range classes {
		// Define the subclass' struct
		contents += fmt.Sprintf("type %s struct {\n", subClass.name)
		contents += baseName + "\n"
		for _, field := range subClass.fields {
			contents += fmt.Sprintf("%s %s\n", field.key, field.value)
		}
		contents += "}\n"

		// Define the sublcass' constructor
		contents += fmt.Sprintf("func New%s(", subClass.name)
		for _, field := range subClass.fields {
			contents += fmt.Sprintf("%s %s, ", field.key, field.value)
		}
		contents += fmt.Sprintf(") *%s{\n", subClass.name)
		contents += fmt.Sprintf("return &%s{\n", subClass.name)
		for _, field := range subClass.fields {
			contents += fmt.Sprintf("%s: %s,\n", field.key, field.key)
		}
		contents += "}\n"
		contents += "}\n"
	}

	return contents, nil
}

func buildContent(baseName string, classes SubClassList) (string, error) {
	var buffer strings.Builder
	buffer.Reset()
	buffer.WriteString("")

	// Starting the file content
	buffer.WriteString("package " + strings.ToLower(baseName) + "\n\n")
	buffer.WriteString("import (\n")
	buffer.WriteString("\"github.com/ByteHunter/glox/token\"\n")
	buffer.WriteString(")\n")
	buffer.WriteString("type " + baseName + " any\n\n")

	// Subclasses
	for _, subClass := range classes {
		// Define the subclass' struct
		buffer.WriteString("type " + subClass.name + " struct {\n")
		buffer.WriteString(baseName + "\n")

		for _, field := range subClass.fields {
			buffer.WriteString("" + field.key + " " + field.value + "\n")
		}
		buffer.WriteString("}\n")

		// Define the sublcass' constructor
		buffer.WriteString("func New" + subClass.name + "(")
		for _, field := range subClass.fields {
			buffer.WriteString("" + field.key + " " + field.value + ", ")
		}
		buffer.WriteString(") *" + subClass.name + "{\n")
		buffer.WriteString("return &" + subClass.name + "{\n")
		for _, field := range subClass.fields {
			buffer.WriteString("" + field.key + ": " + field.key + ",\n")
		}
		buffer.WriteString("}\n")
		buffer.WriteString("}\n")
	}

	return buffer.String(), nil
}
