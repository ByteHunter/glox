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

var Classes = SubClassList{
	SubClassDefinition{
		"Binary",
		FieldList{
			{"left", "Expression"},
			{"operator", "token.Token"},
			{"right", "Expression"},
		},
	},
	SubClassDefinition{
		"Grouping",
		FieldList{
			{"expression", "Expression"},
		},
	},
	SubClassDefinition{
		"Literal",
		FieldList{
			{"value", "any"},
		},
	},
	SubClassDefinition{
		"Unary",
		FieldList{
			{"operator", "token.Token"},
			{"right", "Expression"},
		},
	},
}

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

	contents, err := BuildContent("Expression", Classes)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return 1
	}

	formatted, err := format.Source([]byte(contents))
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return 1
	}
	contents = string(formatted)

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

func BuildContent(baseName string, classes SubClassList) (string, error) {
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
		subClassContent := BuildSubClassContent(baseName, subClass)
		buffer.WriteString(subClassContent)
	}

	return buffer.String(), nil
}

func BuildSubClassContent(baseName string, subClass SubClassDefinition) string {
	var buffer strings.Builder
	buffer.Reset()
	// Define the subclass' struct
	buffer.WriteString("type " + subClass.name + " struct {\n")
	buffer.WriteString(baseName + "\n")

	for _, field := range subClass.fields {
		buffer.WriteString("" + field.key + " " + field.value + "\n")
	}
	buffer.WriteString("}\n\n")

	// Define the sublcass' constructor
	buffer.WriteString("func New" + subClass.name + "(")
	for _, field := range subClass.fields {
		buffer.WriteString("" + field.key + " " + field.value + ", ")
	}
	buffer.WriteString(") *" + subClass.name + " {\n")
	buffer.WriteString("return &" + subClass.name + "{\n")
	for _, field := range subClass.fields {
		buffer.WriteString("" + field.key + ": " + field.key + ",\n")
	}
	buffer.WriteString("}\n")
	buffer.WriteString("}\n")

	return buffer.String()
}
