package main

import "testing"

func TestNewScanner(t *testing.T) {
	ans := NewScanner("")
	if ans == nil {
		t.Errorf("NewScanner() is nil")
	}
}

func TestScanTokensWithoutErrors(t *testing.T) {
	scanner := NewScanner("")
	err := scanner.scanTokens()
	if err != nil {
		t.Errorf("Scanner.scanTokens() want nil but got: %s", err.Error())
	}
}
