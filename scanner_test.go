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

func TestScanMultipleTokens(t *testing.T) {
	scanner := NewScanner("(){},.;-+*")
    scanner.scanTokens()
	actual := scanner.tokens
	expected := []Token{
        {LEFT_PAREN, "(", nil, 1},
        {RIGHT_PAREN, ")", nil, 1},
        {LEFT_BRACE, "{", nil, 1},
        {RIGHT_BRACE, "}", nil, 1},
        {COMMA, ",", nil, 1},
        {DOT, ".", nil, 1},
        {SEMICOLON, ";", nil, 1},
        {MINUS, "-", nil, 1},
        {PLUS, "+", nil, 1},
        {STAR, "*", nil, 1},
        {EOF, "", nil, 1},
    }
	if len(actual) != len(expected) {
		t.Errorf("Expected '%v' but got '%v'", len(expected), len(actual))
	}
    for i := range expected {
        e := expected[i]
        a := actual[i]
        if a != e {
            t.Errorf("Expected '%v' but got '%v' at index %d", e, a, i)
        }
    }
}
