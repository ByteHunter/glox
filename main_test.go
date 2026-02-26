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

func TestLoxError(t *testing.T) {
	actual := captureStdout(t, func() {
		loxError(0, "Test Message")
	})
	expected := "[line 0] Error : Test Message\n"
	if actual != expected {
		t.Errorf("Expecting %s, got: %s", expected, actual)
	}
}

func TestLoxReport(t *testing.T) {
	actual := captureStdout(t, func() {
		loxReport(0, "Somewhere", "Test Message")
	})
	expected := "[line 0] Error Somewhere: Test Message\n"
	if actual != expected {
		t.Errorf("Expecting %s, got: %s", expected, actual)
	}
}
