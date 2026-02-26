package main

import "testing"

func TestNewScanner(t *testing.T) {
	scanner := NewScanner("")
	if scanner == nil {
		t.Errorf("Expected to be not nil")
	}
	var tests = []struct {
		name             string
		actual, expected any
	}{
		{"source", scanner.source, ""},
		{"start", scanner.start, 0},
		{"current", scanner.current, 0},
		{"line", scanner.line, 1},
	}
	for _, test := range tests {
		if test.actual != test.expected {
			t.Errorf(
				"Expected '%s' to be '%v' but got '%v'",
				test.name, test.expected, test.actual,
			)
		}
	}
}

func TestScanTokensWithoutErrors(t *testing.T) {
	scanner := NewScanner("")
	tokens, err := scanner.scanTokens()
	if err != nil {
		t.Errorf("Expected '%v' but got '%v'", nil, err.Error())
	}
	actual := len(tokens)
	expected := 1
	if actual != expected {
		t.Errorf("Expected length '%v' but got '%v'", expected, actual)
	}
}

func TestIsAtEnd(t *testing.T) {
	scanner := NewScanner("")
	actual := scanner.isAtEnd()
	expected := true
	if actual != expected {
		t.Errorf("Expected to be '%v' but got '%v'", expected, actual)
	}
}

func TestAdvanceWithSource(t *testing.T) {
	scanner := NewScanner("(")
	actual := scanner.advance()
	expected := byte('(')
	if actual != expected {
		t.Errorf("Expected '%v' but got '%v'", expected, actual)
	}
}

func TestAddSimpleToken(t *testing.T) {
	scanner := NewScanner("")
	scanner.addSimpleToken(EOF)
	actual := len(scanner.tokens)
	expected := 1
	if actual != expected {
		t.Errorf("Expected '%v' but got '%v'", expected, actual)
	}
}

func TestAddToken(t *testing.T) {
	scanner := NewScanner("")
	scanner.addToken(EOF, nil)
	actual := len(scanner.tokens)
	expected := 1
	if actual != expected {
		t.Errorf("Expected '%v' but got '%v'", expected, actual)
	}
}
