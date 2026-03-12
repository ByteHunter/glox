package utils

import (
	"fmt"
	"testing"
)

func TestCaptureStdout(t *testing.T) {
	actual := CaptureStdout(t, func() {
		fmt.Println("Example string")
	})
	expected := "Example string\n"

	if actual != expected {
		t.Errorf("Expected length '%v' but got '%v'", expected, actual)
	}
}
