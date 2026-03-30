package astprinter

import (
	"fmt"

	"github.com/ByteHunter/glox/expression"
)

func ExampleAstPrinter_Print_nil_expression() {
	result := NewAstPrinter().Print(nil)

	fmt.Println(result)
	// Output:
	// nil
}

func ExampleAstPrinter_Parentesize() {
	result := NewAstPrinter().Parentesize("test")

	fmt.Println(result)
	// Output:
	// (test)
}

func ExampleAstPrinter_Parentesize_with_expression() {
	result := NewAstPrinter().Parentesize(
		"test",
		expression.NewLiteral(42),
	)

	fmt.Println(result)
	// Output:
	// (test 42)
}
