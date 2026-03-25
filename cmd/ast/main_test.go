package main

import (
	"flag"
	"fmt"
	"os"
	"testing"

	"github.com/ByteHunter/glox/utils"
)

var EmptySubClassList = SubClassList{}
var SimpleSubClassList = SubClassList{
	SubClassDefinition{
		"TestClass",
		FieldList{
			{"field", "Expression"},
		},
	},
}
var CompleteSubClassList = SubClassList{
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
}

func TestRunMainWithoutArguments(t *testing.T) {
	oldArgs := os.Args
	wd, _ := os.Getwd()
	os.Chdir("../../")
	defer func() {
		os.Args = oldArgs
		os.Chdir(wd)
	}()

	flag.NewFlagSet("Test flags", flag.ExitOnError)
	os.Args = append([]string{"Test flags"}, []string{}...)

	actualExit := -1
	actualStdout := utils.CaptureStdout(t, func() {
		actualExit = RunMain()
	})
	expectedExit := 1
	expectedStdout := "Usage go run cmd/ast/main.go <output-path>\n"

	if actualStdout != expectedStdout {
		t.Errorf("Expected '%v' but got '%v'", expectedStdout, actualStdout)
	}
	if actualExit != expectedExit {
		t.Errorf("Expected '%v' but got '%v'", expectedExit, actualExit)
	}
}

func ExampleBuildContent() {
	content, _ := BuildContent("TestClass", EmptySubClassList)
	fmt.Print(content)
	// Output:
	// package testclass
	//
	// import (
	// "github.com/ByteHunter/glox/token"
	// )
	//
	// type TestClass interface {
	// accept(v Visitor) any
	// }
	//
	// type Visitor interface {
	// }
}

func ExampleBuildSubClassContent() {
	subClass := SubClassDefinition{
		"Literal",
		FieldList{
			{"value", "any"},
		},
	}
	subClassContent := BuildSubClassContent("TestClass", subClass)
	fmt.Print(subClassContent)
	// Output:
	// type Literal struct {
	// TestClass
	// value any
	// }
	//
	// func NewLiteral(value any, ) *Literal {
	// return &Literal{
	// value: value,
	// }
	// }
	//
	// func (literal *Literal) accept(v Visitor) any {
	// return v.visitLiteralExpression(literal)
	// }
}

func BenchmarkBuildContentEmpty(b *testing.B) {
	for b.Loop() {
		BuildContent("TestClass", EmptySubClassList)
	}
}

func BenchmarkBuildContentSimple(b *testing.B) {
	for b.Loop() {
		BuildContent("TestClass", SimpleSubClassList)
	}
}

func BenchmarkBuildContentComplete(b *testing.B) {
	for b.Loop() {
		BuildContent("TestClass", CompleteSubClassList)
	}
}
