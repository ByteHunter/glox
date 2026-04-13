package astprinter

import (
	"fmt"

	syntax_expression "github.com/ByteHunter/glox/syntax"
	"github.com/ByteHunter/glox/token"
)

func ExampleAstPrinter_Parentesize() {
	result := NewAstPrinter().Parentesize("test")

	fmt.Println(result)
	// Output:
	// (test)
}

func ExampleAstPrinter_Parentesize_without_expressions() {
	result := NewAstPrinter().Parentesize(
		"test",
	)

	fmt.Println(result)
	// Output:
	// (test)
}

func ExampleAstPrinter_Parentesize_with_nil_expression() {
	result := NewAstPrinter().Parentesize(
		"test",
		nil,
	)

	fmt.Println(result)
	// Output:
	// (test nil)
}

func ExampleAstPrinter_Parentesize_with_multiple_expressions() {
	result := NewAstPrinter().Parentesize(
		"test",
		nil,
		nil,
	)

	fmt.Println(result)
	// Output:
	// (test nil nil)
}

func ExampleAstPrinter_Parentesize_binary() {
	result := NewAstPrinter().Parentesize(
		"test",
		syntax_expression.NewBinary(nil, *token.NewToken(token.PLUS, "+", nil, 1), nil),
	)

	fmt.Println(result)
	// Output:
	// (test (+ nil nil))
}

func ExampleAstPrinter_Parentesize_grouping() {
	result := NewAstPrinter().Parentesize(
		"test",
		syntax_expression.NewGrouping(nil),
	)

	fmt.Println(result)
	// Output:
	// (test (group nil))
}

func ExampleAstPrinter_Parentesize_literal() {
	result := NewAstPrinter().Parentesize(
		"test",
		syntax_expression.NewLiteral(42),
	)

	fmt.Println(result)
	// Output:
	// (test 42)
}

func ExampleAstPrinter_Parentesize_literal_nil() {
	result := NewAstPrinter().Parentesize(
		"test",
		syntax_expression.NewLiteral(nil),
	)

	fmt.Println(result)
	// Output:
	// (test nil)
}

func ExampleAstPrinter_Parentesize_unary() {
	result := NewAstPrinter().Parentesize(
		"test",
		syntax_expression.NewUnary(*token.NewToken(token.PLUS, "+", nil, 1), nil),
	)

	fmt.Println(result)
	// Output:
	// (test (+ nil))
}

func ExampleAstPrinter_Print_nil_expression() {
	result, _ := NewAstPrinter().Print(nil)

	fmt.Println(result)
	// Output:
	// nil
}

func ExampleAstPrinter_Print() {
	// Testing expression: -123 * 45.67
	result, _ := NewAstPrinter().Print(
		syntax_expression.NewBinary(
			syntax_expression.NewUnary(
				*token.NewToken(token.MINUS, "-", nil, 1),
				syntax_expression.NewLiteral(123),
			),
			*token.NewToken(token.STAR, "*", nil, 1),
			syntax_expression.NewGrouping(
				syntax_expression.NewLiteral(45.67),
			),
		),
	)

	fmt.Println(result)
	// Output:
	// (* (- 123) (group 45.67))
}
