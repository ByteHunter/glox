package reporting

import (
	"fmt"
	"os"

	"github.com/ByteHunter/glox/token"
)

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
