package main

import "testing"

func TestRunFileIsImplemented(t *testing.T) {
	ans := runFile("something")
	if ans != nil {
		t.Errorf("runFile() = %s, want nil", ans)
	}
}

func TestRunInteractiveIsImplemented(t *testing.T) {
	ans := runInteractive()
	if ans != nil {
		t.Errorf("runInteractive() = %s, want nil", ans)
	}
}

func TestRunIsImplemented(t *testing.T) {
	ans := run("")
	if ans != nil {
		t.Errorf("run() = %s, want nil", ans)
	}
}
