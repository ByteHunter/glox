package reporting

import (
	"testing"

	"github.com/ByteHunter/glox/utils"
)

func TestLoxError(t *testing.T) {
	actual := utils.CaptureStdout(t, func() {
		LoxError(0, "Test Message")
	})
	expected := "[line 0] Error : Test Message\n"
	if actual != expected {
		t.Errorf("Expecting %s, got: %s", expected, actual)
	}
}

func TestLoxReport(t *testing.T) {
	actual := utils.CaptureStdout(t, func() {
		LoxReport(0, "Somewhere", "Test Message")
	})
	expected := "[line 0] Error Somewhere: Test Message\n"
	if actual != expected {
		t.Errorf("Expecting %s, got: %s", expected, actual)
	}
}
