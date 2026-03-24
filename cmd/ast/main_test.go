package main

import (
	"flag"
	"os"
	"testing"

	"github.com/ByteHunter/glox/utils"
)

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

func TestDefineAst(t *testing.T) {
	t.Skip()
}
