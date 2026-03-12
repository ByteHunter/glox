package utils

import (
	"io"
	"os"
	"testing"
)

func CaptureStdout(t *testing.T, runnable func()) string {
	realStdout := os.Stdout
	defer func() {
		os.Stdout = realStdout
	}()

	testStdin, testStdout, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	os.Stdout = testStdout
	runnable()
	err = testStdout.Close()
	if err != nil {
		t.Fatal(err)
	}

	outBytes, err := io.ReadAll(testStdin)
	if err != nil {
		t.Fatal(err)
	}
	err = testStdin.Close()
	if err != nil {
		t.Fatal(err)
	}

	return string(outBytes)
}
