package scanner

import (
	"testing"

	"github.com/ByteHunter/glox/token"
	"github.com/ByteHunter/glox/utils"
)

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
	tokens, err := scanner.ScanTokens()
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
	scanner.addSimpleToken(token.EOF)
	actual := len(scanner.tokens)
	expected := 1
	if actual != expected {
		t.Errorf("Expected '%v' but got '%v'", expected, actual)
	}
}

func TestAddToken(t *testing.T) {
	scanner := NewScanner("")
	scanner.addToken(token.EOF, nil)
	actual := len(scanner.tokens)
	expected := 1
	if actual != expected {
		t.Errorf("Expected '%v' but got '%v'", expected, actual)
	}
}

func TestSimpleTokens(t *testing.T) {
	scanner := NewScanner("(){},.;-+*\n")
	scanner.ScanTokens()
	actual := scanner.tokens
	expected := []token.Token{
		{Type: token.LEFT_PAREN, Lexeme: "(", Literal: nil, Line: 1},
		{Type: token.RIGHT_PAREN, Lexeme: ")", Literal: nil, Line: 1},
		{Type: token.LEFT_BRACE, Lexeme: "{", Literal: nil, Line: 1},
		{Type: token.RIGHT_BRACE, Lexeme: "}", Literal: nil, Line: 1},
		{Type: token.COMMA, Lexeme: ",", Literal: nil, Line: 1},
		{Type: token.DOT, Lexeme: ".", Literal: nil, Line: 1},
		{Type: token.SEMICOLON, Lexeme: ";", Literal: nil, Line: 1},
		{Type: token.MINUS, Lexeme: "-", Literal: nil, Line: 1},
		{Type: token.PLUS, Lexeme: "+", Literal: nil, Line: 1},
		{Type: token.STAR, Lexeme: "*", Literal: nil, Line: 1},
		{Type: token.EOF, Lexeme: "", Literal: nil, Line: 2},
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

func TestOperatorTokens(t *testing.T) {
	scanner := NewScanner("! = < > != == <= >=\n")
	scanner.ScanTokens()
	actual := scanner.tokens
	expected := []token.Token{
		{Type: token.BANG, Lexeme: "!", Literal: nil, Line: 1},
		{Type: token.EQUAL, Lexeme: "=", Literal: nil, Line: 1},
		{Type: token.LESS, Lexeme: "<", Literal: nil, Line: 1},
		{Type: token.GREATER, Lexeme: ">", Literal: nil, Line: 1},
		{Type: token.BANQ_EQUAL, Lexeme: "!=", Literal: nil, Line: 1},
		{Type: token.EQUAL_EQUAL, Lexeme: "==", Literal: nil, Line: 1},
		{Type: token.LESS_EQUAL, Lexeme: "<=", Literal: nil, Line: 1},
		{Type: token.GREATER_EQUAL, Lexeme: ">=", Literal: nil, Line: 1},
		{Type: token.EOF, Lexeme: "", Literal: nil, Line: 2},
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

func TestComment(t *testing.T) {
	content := "// This is a comment\n/\n" +
		".// This is also a comment\n"
	scanner := NewScanner(content)
	scanner.ScanTokens()
	actual := scanner.tokens
	expected := []token.Token{
		{Type: token.SLASH, Lexeme: "/", Literal: nil, Line: 2},
		{Type: token.DOT, Lexeme: ".", Literal: nil, Line: 3},
		{Type: token.EOF, Lexeme: "", Literal: nil, Line: 4},
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

func TestString(t *testing.T) {
	content := "\"Example string\"\n"
	scanner := NewScanner(content)
	scanner.ScanTokens()
	actual := scanner.tokens
	expected := []token.Token{
		{Type: token.STRING, Lexeme: "\"Example string\"", Literal: "Example string", Line: 1},
		{Type: token.EOF, Lexeme: "", Literal: nil, Line: 2},
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

func TestAllowMultilineStrings(t *testing.T) {
	content := "\"Example of a multiline string\nAnother line\"\n"
	scanner := NewScanner(content)
	scanner.ScanTokens()
	actual := scanner.tokens
	expected := []token.Token{
		{Type: token.STRING, Lexeme: "\"Example of a multiline string\nAnother line\"", Literal: "Example of a multiline string\nAnother line", Line: 2},
		{Type: token.EOF, Lexeme: "", Literal: nil, Line: 3},
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

func TestUnterminatedString(t *testing.T) {
	content := "\"Example string\n"
	scanner := NewScanner(content)
	actual := utils.CaptureStdout(t, func() {
		scanner.ScanTokens()
	})

	expected := "[line 2] Error : Unterminated string.\n"
	if actual != expected {
		t.Errorf("Expected '%v' but got '%v'", expected, actual)
	}
}

func TestNumber(t *testing.T) {
	var tests = []struct {
		name     string
		content  string
		expected []token.Token
	}{
		{
			"Valid number",
			"123\n",
			[]token.Token{
				{Type: token.NUMBER, Lexeme: "123", Literal: float64(123), Line: 1},
				{Type: token.EOF, Lexeme: "", Literal: nil, Line: 2},
			},
		},
		{
			"Valid number",
			"00001\n",
			[]token.Token{
				{Type: token.NUMBER, Lexeme: "00001", Literal: float64(1), Line: 1},
				{Type: token.EOF, Lexeme: "", Literal: nil, Line: 2},
			},
		},
		{
			"Invalid number",
			"123.\n",
			[]token.Token{
				{Type: token.NUMBER, Lexeme: "123", Literal: float64(123), Line: 1},
				{Type: token.DOT, Lexeme: ".", Literal: nil, Line: 1},
				{Type: token.EOF, Lexeme: "", Literal: nil, Line: 2},
			},
		},
	}

	for set, test := range tests {

		scanner := NewScanner(test.content)
		scanner.ScanTokens()
		actual := scanner.tokens
		if len(actual) != len(test.expected) {
			t.Errorf("Data Set #%d Expected '%v' but got '%v'", set, len(test.expected), len(actual))
			// t.Errorf("%v\n", actual)
		}
		for i := range test.expected {
			e := test.expected[i]
			a := actual[i]
			if a != e {
				t.Errorf("Data Set #%d Expected '%v' but got '%v' at index %d", set, e, a, i)
			}
		}
	}

}

func TestUnexpectedCharacter(t *testing.T) {
	scanner := NewScanner("?")

	actual := utils.CaptureStdout(t, func() {
		scanner.ScanTokens()
	})
	expected := "[line 1] Error : Unexpected character ?.\n"
	if actual != expected {
		t.Errorf("Expected '%v' but got '%v'", expected, actual)
	}
}
