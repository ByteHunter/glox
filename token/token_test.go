package token

import "testing"

func TestTokenToString(t *testing.T) {
	token := NewToken(LEFT_PAREN, "(", "(", 1)
	actual := token.toString()
	expected := "LEFT_PAREN ( ("

	if actual != expected {
		t.Errorf("Expected '%s' but got: '%s'", expected, actual)
	}
}

func TestTokenToStringWithNewline(t *testing.T) {
	token := NewToken(STRING, "\n", "\n", 1)
	actual := token.toString()
	expected := "STRING \\n \n"

	if actual != expected {
		t.Errorf("Expected '%s' but got: '%s'", expected, actual)
	}
}
