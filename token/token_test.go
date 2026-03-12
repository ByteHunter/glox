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
