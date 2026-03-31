package astprinter

import (
	"fmt"

	"github.com/ByteHunter/glox/expression"
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
		expression.NewBinary(nil, *token.NewToken(token.PLUS, "+", nil, 1), nil),
	)

	fmt.Println(result)
	// Output:
	// (test (+ nil nil))
}

func ExampleAstPrinter_Parentesize_grouping() {
	result := NewAstPrinter().Parentesize(
		"test",
		expression.NewGrouping(nil),
	)

	fmt.Println(result)
	// Output:
	// (test (group nil))
}

func ExampleAstPrinter_Parentesize_literal() {
	result := NewAstPrinter().Parentesize(
		"test",
		expression.NewLiteral(42),
	)

	fmt.Println(result)
	// Output:
	// (test 42)
}

func ExampleAstPrinter_Parentesize_literal_nil() {
	result := NewAstPrinter().Parentesize(
		"test",
		expression.NewLiteral(nil),
	)

	fmt.Println(result)
	// Output:
	// (test nil)
}

func ExampleAstPrinter_Parentesize_unary() {
	result := NewAstPrinter().Parentesize(
		"test",
		expression.NewUnary(*token.NewToken(token.PLUS, "+", nil, 1), nil),
	)

	fmt.Println(result)
	// Output:
	// (test (+ nil))
}

func ExampleAstPrinter_Print_nil_expression() {
	result := NewAstPrinter().Print(nil)

	fmt.Println(result)
	// Output:
	// nil
}

func ExampleAstPrinter_Print() {
	// Testing expression: -123 * 45.67
	result := NewAstPrinter().Print(
		expression.NewBinary(
			expression.NewUnary(
				*token.NewToken(token.MINUS, "-", nil, 1),
				expression.NewLiteral(123),
			),
			*token.NewToken(token.STAR, "*", nil, 1),
			expression.NewGrouping(
				expression.NewLiteral(45.67),
			),
		),
	)

	fmt.Println(result)
	// Output:
	// (* (- 123) (group 45.67))
}
