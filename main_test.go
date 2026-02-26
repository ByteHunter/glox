package main

import "testing"

func TestRunExistingFileWithoutErrors(t *testing.T) {
	ans := runFile("main.lox")
	if ans != nil {
		t.Errorf("runFile() = %s, want nil", ans)
	}
}

func TestRunEmptyWithoutErrors(t *testing.T) {
	ans := run("")
	if ans != nil {
		t.Errorf("run() = %s, want nil", ans)
	}
}
