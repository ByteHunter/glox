package main

import (
	"flag"
	"os"
	"testing"

	"github.com/ByteHunter/glox/utils"
)

func TestRunExistingFileWithoutErrors(t *testing.T) {
	err := runFile("./tests/resources/main.lox")
	if err != nil {
		t.Errorf("runFile() want nil but got: %s", err)
	}
}

func TestRunFileWithErrors(t *testing.T) {
	err := runFile("")
	if err == nil {
		t.Errorf("Expected an error but got nil")
	}
	actual := err.Error()
	expected := "open : no such file or directory"
	if actual != expected {
		t.Errorf("Expected length '%v' but got '%v'", expected, actual)
	}
}

func TestRunEmptyWithoutErrors(t *testing.T) {
	err := run("")
	if err != nil {
		t.Errorf("run() want nil but got: %s", err)
	}
}

func TestRunMainTooManyArguments(t *testing.T) {
	oldArgs := os.Args
	defer func() {
		os.Args = oldArgs
	}()

	flag.NewFlagSet("Test flags", flag.ExitOnError)
	os.Args = append([]string{"Test flags"}, []string{"file", "file"}...)

	actualExit := -1
	actualStdout := utils.CaptureStdout(t, func() {
		actualExit = RunMain()
	})
	expectedExit := 64
	expectedStdout := "Error: Too many arguments!\n"

	if actualStdout != expectedStdout {
		t.Errorf("Expected '%v' but got '%v'", expectedStdout, actualStdout)
	}
	if actualExit != expectedExit {
		t.Errorf("Expected '%v' but got '%v'", expectedExit, actualExit)
	}
}

func TestRunMainRunFile(t *testing.T) {
	oldArgs := os.Args
	defer func() {
		os.Args = oldArgs
	}()

	flag.NewFlagSet("Test flags", flag.ExitOnError)
	os.Args = append([]string{"Test flags"}, []string{"./tests/resources/main.lox"}...)

	actualExit := -1
	actualStdout := utils.CaptureStdout(t, func() {
		actualExit = RunMain()
	})
	expectedExit := 0
	expectedStdout := "(== 42 42)\n"

	if actualStdout != expectedStdout {
		t.Errorf("Expected '%v' but got '%v'", expectedStdout, actualStdout)
	}
	if actualExit != expectedExit {
		t.Errorf("Expected '%v' but got '%v'", expectedExit, actualExit)
	}
}
