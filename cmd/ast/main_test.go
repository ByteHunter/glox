package main

import (
	"flag"
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

func TestRunMain(t *testing.T) {
	oldArgs := os.Args
	wd, _ := os.Getwd()
	os.Chdir("../../")
	defer func() {
		os.Args = oldArgs
		os.Chdir(wd)
	}()

	flag.NewFlagSet("Test flags", flag.ExitOnError)
	os.Args = append([]string{"Test flags"}, []string{"tests"}...)

	actualExit := -1
	actualStdout := utils.CaptureStdout(t, func() {
		actualExit = RunMain()
	})
	expectedExit := 0
	expectedStdout := ""

	if actualStdout != expectedStdout {
		t.Errorf("Expected '%v' but got '%v'", expectedStdout, actualStdout)
	}
	if actualExit != expectedExit {
		t.Errorf("Expected '%v' but got '%v'", expectedExit, actualExit)
	}
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

func TestGenerageContent(t *testing.T) {
	t.Skip()
}

func BenchmarkGenerageContentEmpty(b *testing.B) {
	for b.Loop() {
		generateContent("TestClass", EmptySubClassList)
	}
}

func BenchmarkBuildContentEmpty(b *testing.B) {
	for b.Loop() {
		buildContent("TestClass", EmptySubClassList)
	}
}

func BenchmarkGenerageContentSimple(b *testing.B) {
	for b.Loop() {
		generateContent("TestClass", SimpleSubClassList)
	}
}

func BenchmarkBuildContentSimple(b *testing.B) {
	for b.Loop() {
		buildContent("TestClass", SimpleSubClassList)
	}
}

func BenchmarkGenerageContentComplete(b *testing.B) {
	for b.Loop() {
		generateContent("TestClass", CompleteSubClassList)
	}
}

func BenchmarkBuildContentComplete(b *testing.B) {
	for b.Loop() {
		buildContent("TestClass", CompleteSubClassList)
	}
}
