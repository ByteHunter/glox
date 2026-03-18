package utils

import (
	"io"
	"os"
	"testing"
)

func FailOnError(t *testing.T, e error) {
	if e != nil {
		t.Fatal(e)
	}
}

func CaptureStdout(t *testing.T, runnable func()) string {
	realStdout := os.Stdout
	defer func() {
		os.Stdout = realStdout
	}()

	testStdin, testStdout, err := os.Pipe()
	FailOnError(t, err)
	os.Stdout = testStdout
	runnable()
	err = testStdout.Close()
	FailOnError(t, err)

	outBytes, err := io.ReadAll(testStdin)
	FailOnError(t, err)
	err = testStdin.Close()
	FailOnError(t, err)

	return string(outBytes)
}
