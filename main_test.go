package main

import (
	"testing"
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
