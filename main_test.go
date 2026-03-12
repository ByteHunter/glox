package main

import (
	"testing"
)

func TestRunExistingFileWithoutErrors(t *testing.T) {
	err := runFile("main.lox")
	if err != nil {
		t.Errorf("runFile() want nil but got: %s", err)
	}
}

func TestRunEmptyWithoutErrors(t *testing.T) {
	err := run("")
	if err != nil {
		t.Errorf("run() want nil but got: %s", err)
	}
}
