package reporting

import (
	"fmt"
	"os"

	"github.com/ByteHunter/glox/token"
)

type RuntimeError struct {
	Operator token.Token
	message  string
}

func NewRuntimeError(operator token.Token, message string) *RuntimeError {
	return &RuntimeError{
		Operator: operator,
		message:  message,
	}
}

func (r RuntimeError) Error() string {
	return fmt.Sprintf("RuntimeError %s", r.message)
}

func LoxTokenError(t token.Token, message string) {
	if t.Type == token.EOF {
		LoxReport(t.Line, "at end", message)
		return
	}

	LoxReport(t.Line, "at '" + t.Lexeme + "'", message)
}

func LoxError(line int, message string) {
	LoxReport(line, "", message)
}

func LoxReport(line int, where string, message string) {
	fmt.Fprintf(os.Stdout, "[line %d] Error %s: %s\n", line, where, message)
}
