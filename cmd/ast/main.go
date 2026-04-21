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

var ExpressionClasses = SubClassList{
	SubClassDefinition{
		"Assign",
		FieldList{
			{"Name", "token.Token"},
			{"Expr", "Expression"},
		},
	},
	SubClassDefinition{
		"Binary",
		FieldList{
			{"Left", "Expression"},
			{"Operator", "token.Token"},
			{"Right", "Expression"},
		},
	},
	SubClassDefinition{
		"Grouping",
		FieldList{
			{"Expr", "Expression"},
		},
	},
	SubClassDefinition{
		"Literal",
		FieldList{
			{"Value", "any"},
		},
	},
	SubClassDefinition{
		"Unary",
		FieldList{
			{"Operator", "token.Token"},
			{"Right", "Expression"},
		},
	},
	SubClassDefinition{
		"Variable",
		FieldList{
			{"Name", "token.Token"},
		},
	},
}

var StatementClasses = SubClassList{
	SubClassDefinition{
		"Expression",
		FieldList{
			{"Expr", "expression.Expression"},
		},
	},
	SubClassDefinition{
		"Print",
		FieldList{
			{"Expr", "expression.Expression"},
		},
	},
	SubClassDefinition{
		"Variable",
		FieldList{
			{"Name", "token.Token"},
			{"Initializer", "expression.Expression"},
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

	res := generateAst(outputDirectory, "Expression", ExpressionClasses)
	if res != 0 {
		return res
	}
	res = generateAst(outputDirectory, "Statement", StatementClasses)
	if res != 0 {
		return res
	}

	return 0
}

func generateAst(outputDirectory, baseName string, classes SubClassList) int {
	file, err := getFile(outputDirectory, baseName)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return 1
	}
	defer file.Close()

	contents, err := BuildContent(baseName, classes)
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
	path := fmt.Sprintf("%s/%s/%s.go", outputDir, strings.ToLower(baseName), strings.ToLower(baseName))

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
	buffer.WriteString(")\n\n")

	// Expression interface
	buffer.WriteString("type " + baseName + " interface {\n")
	buffer.WriteString("Accept(v Visitor) (any, error)\n")
	buffer.WriteString("}\n\n")

	// Visitor interface
	visitorInterfaceContent := BuildVistitorInterface(baseName, classes)
	buffer.WriteString(visitorInterfaceContent)

	// Subclasses
	for _, subClass := range classes {
		subClassContent := BuildSubClassContent(baseName, subClass)
		buffer.WriteString(subClassContent)
	}

	return buffer.String(), nil
}

func BuildVistitorInterface(baseName string, subClasses SubClassList) string {
	var buffer strings.Builder
	buffer.Reset()

	buffer.WriteString("type Visitor interface {\n")
	for _, subClass := range subClasses {
		buffer.WriteString("Visit" + subClass.name + baseName + "(*" + subClass.name + baseName + ") (any, error)\n")
	}
	buffer.WriteString("}\n\n")

	return buffer.String()
}

func BuildSubClassContent(baseName string, subClass SubClassDefinition) string {
	var buffer strings.Builder
	buffer.Reset()
	// Define the struct
	buffer.WriteString("type " + subClass.name + baseName + " struct {\n")
	buffer.WriteString(baseName + "\n")

	// Define the fields
	for _, field := range subClass.fields {
		buffer.WriteString("" + field.key + " " + field.value + "\n")
	}
	buffer.WriteString("}\n\n")

	// Define the  constructor
	buffer.WriteString("func New" + subClass.name + "(")
	for _, field := range subClass.fields {
		buffer.WriteString("" + strings.ToLower(field.key) + " " + field.value + ", ")
	}
	buffer.WriteString(") *" + subClass.name + baseName + " {\n")
	buffer.WriteString("return &" + subClass.name + baseName + "{\n")
	for _, field := range subClass.fields {
		buffer.WriteString("" + field.key + ": " + strings.ToLower(field.key) + ",\n")
	}
	buffer.WriteString("}\n")
	buffer.WriteString("}\n\n")

	// Define the accept method
	buffer.WriteString("func (" + strings.ToLower(subClass.name) + " *" + subClass.name + baseName + ") Accept(v Visitor) (any, error) {\n")
	buffer.WriteString("return v.Visit" + subClass.name + baseName + "(" + strings.ToLower(subClass.name) + ")\n")
	buffer.WriteString("}\n")

	return buffer.String()
}
