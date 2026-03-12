package reporting

import (
	"fmt"
	"os"
)

func LoxError(line int, message string) {
	LoxReport(line, "", message)
}

func LoxReport(line int, where string, message string) {
	fmt.Fprintf(os.Stdout, "[line %d] Error %s: %s\n", line, where, message)
}
