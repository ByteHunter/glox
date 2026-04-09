package interpreter

import (
	"fmt"
	"testing"

	"github.com/ByteHunter/glox/expression"
	"github.com/ByteHunter/glox/parser"
	"github.com/ByteHunter/glox/scanner"
	"github.com/ByteHunter/glox/token"
	"github.com/ByteHunter/glox/utils"
)

// Evaluating Unary expressions

func ExampleInterpreter_Evaluate_unary_nil() {
	i := NewInterpreter()
	expr := expression.NewUnary(
		*token.NewToken(token.MINUS, "-", nil, 1),
		nil,
	)
	result, err := i.Evaluate(expr)
	fmt.Println(err)
	fmt.Println(result)

	// Output:
	// RuntimeError Expected an expression, nil found
	// <nil>
}

func ExampleInterpreter_Evaluate_unary_unkown_token() {
	i := NewInterpreter()
	expr := expression.NewUnary(
		*token.NewToken(token.PLUS, "+", nil, 1),
		expression.NewLiteral(42),
	)
	result, err := i.Evaluate(expr)
	fmt.Println(err)
	fmt.Println(result)

	// Output:
	// RuntimeError Unknown unary operator
	// <nil>
}

func ExampleInterpreter_Evaluate_unary_minus_nan() {
	i := NewInterpreter()
	expr := expression.NewUnary(
		*token.NewToken(token.MINUS, "-", nil, 1),
		expression.NewLiteral(true),
	)
	result, err := i.Evaluate(expr)
	fmt.Println(err)
	fmt.Println(result)

	// Output:
	// RuntimeError Cannot convert to float64, unexpected type (ConversionError)
	// NaN
}

func ExampleInterpreter_Evaluate_unary_minus() {
	i := NewInterpreter()
	expr := expression.NewUnary(
		*token.NewToken(token.MINUS, "-", nil, 1),
		expression.NewLiteral(42),
	)
	result, err := i.Evaluate(expr)
	fmt.Println(err)
	fmt.Println(result)

	// Output:
	// <nil>
	// -42
}

func ExampleInterpreter_Evaluate_unary_minus2() {
	i := NewInterpreter()
	expr := expression.NewUnary(
		*token.NewToken(token.MINUS, "-", nil, 1),
		expression.NewLiteral(-42),
	)
	result, err := i.Evaluate(expr)
	fmt.Println(err)
	fmt.Println(result)

	// Output:
	// <nil>
	// 42
}

func ExampleInterpreter_Evaluate_unary_bang_true() {
	i := NewInterpreter()
	expr := expression.NewUnary(
		*token.NewToken(token.BANG, "!", nil, 1),
		expression.NewLiteral(true),
	)
	result, err := i.Evaluate(expr)
	fmt.Println(err)
	fmt.Println(result)

	// Output:
	// <nil>
	// false
}

func ExampleInterpreter_Evaluate_unary_bang_false() {
	i := NewInterpreter()
	expr := expression.NewUnary(
		*token.NewToken(token.BANG, "!", nil, 1),
		expression.NewLiteral(false),
	)
	result, err := i.Evaluate(expr)
	fmt.Println(err)
	fmt.Println(result)

	// Output:
	// <nil>
	// true
}

func ExampleInterpreter_Evaluate_unary_bang_other() {
	i := NewInterpreter()
	expr := expression.NewUnary(
		*token.NewToken(token.BANG, "!", nil, 1),
		expression.NewLiteral(42),
	)
	result, err := i.Evaluate(expr)
	fmt.Println(err)
	fmt.Println(result)

	// Output:
	// <nil>
	// false
}

// Evaluating Binary Expressions

func ExampleInterpreter_Evaluate_binary_invalid_operator() {
	i := NewInterpreter()
	expr := expression.NewBinary(
		expression.NewLiteral(42),
		*token.NewToken(token.DOT, "-", nil, 1),
		expression.NewLiteral(1),
	)
	result, err := i.Evaluate(expr)
	fmt.Println(err)
	fmt.Println(result)

	// Output:
	// RuntimeError Unknown binary operator
	// <nil>
}

func ExampleInterpreter_Evaluate_binary_missing_left_operand() {
	i := NewInterpreter()
	expr := expression.NewBinary(
		nil,
		*token.NewToken(token.MINUS, "-", nil, 1),
		expression.NewLiteral(1),
	)
	result, err := i.Evaluate(expr)
	fmt.Println(err)
	fmt.Println(result)

	// Output:
	// RuntimeError Left operand expected to be an expression, nil found
	// <nil>
}

func ExampleInterpreter_Evaluate_binary_missing_right_operand() {
	i := NewInterpreter()
	expr := expression.NewBinary(
		expression.NewLiteral(42),
		*token.NewToken(token.MINUS, "-", nil, 1),
		nil,
	)
	result, err := i.Evaluate(expr)
	fmt.Println(err)
	fmt.Println(result)

	// Output:
	// RuntimeError Right operand expected to be an expression, nil found
	// <nil>
}

func ExampleInterpreter_Evaluate_binary_minus() {
	i := NewInterpreter()
	expr := expression.NewBinary(
		expression.NewLiteral(43),
		*token.NewToken(token.MINUS, "-", nil, 1),
		expression.NewLiteral(1),
	)
	result, err := i.Evaluate(expr)
	fmt.Println(err)
	fmt.Println(result)

	// Output:
	// <nil>
	// 42
}

func ExampleInterpreter_Evaluate_binary_minus_error_left() {
	i := NewInterpreter()
	expr := expression.NewBinary(
		expression.NewLiteral(43),
		*token.NewToken(token.MINUS, "-", nil, 1),
		expression.NewLiteral(nil),
	)
	result, err := i.Evaluate(expr)
	fmt.Println(err)
	fmt.Println(result)

	// Output:
	// RuntimeError Cannot convert to float64, unexpected type (ConversionError)
	// <nil>
}

func ExampleInterpreter_Evaluate_binary_minus_error_right() {
	i := NewInterpreter()
	expr := expression.NewBinary(
		expression.NewLiteral(nil),
		*token.NewToken(token.MINUS, "-", nil, 1),
		expression.NewLiteral(1),
	)
	result, err := i.Evaluate(expr)
	fmt.Println(err)
	fmt.Println(result)

	// Output:
	// RuntimeError Cannot convert to float64, unexpected type (ConversionError)
	// <nil>
}

func ExampleInterpreter_Evaluate_binary_slash() {
	i := NewInterpreter()
	expr := expression.NewBinary(
		expression.NewLiteral(42),
		*token.NewToken(token.SLASH, "/", nil, 1),
		expression.NewLiteral(2),
	)
	result, err := i.Evaluate(expr)
	fmt.Println(err)
	fmt.Println(result)

	// Output:
	// <nil>
	// 21
}

func ExampleInterpreter_Evaluate_binary_slash_error_left() {
	i := NewInterpreter()
	expr := expression.NewBinary(
		expression.NewLiteral(nil),
		*token.NewToken(token.SLASH, "/", nil, 1),
		expression.NewLiteral(2),
	)
	result, err := i.Evaluate(expr)
	fmt.Println(err)
	fmt.Println(result)

	// Output:
	// RuntimeError Cannot convert to float64, unexpected type (ConversionError)
	// <nil>
}

func ExampleInterpreter_Evaluate_binary_slash_error_right() {
	i := NewInterpreter()
	expr := expression.NewBinary(
		expression.NewLiteral(42),
		*token.NewToken(token.SLASH, "/", nil, 1),
		expression.NewLiteral(nil),
	)
	result, err := i.Evaluate(expr)
	fmt.Println(err)
	fmt.Println(result)

	// Output:
	// RuntimeError Cannot convert to float64, unexpected type (ConversionError)
	// <nil>
}

func ExampleInterpreter_Evaluate_binary_star() {
	i := NewInterpreter()
	expr := expression.NewBinary(
		expression.NewLiteral(21),
		*token.NewToken(token.STAR, "*", nil, 1),
		expression.NewLiteral(2),
	)
	result, err := i.Evaluate(expr)
	fmt.Println(err)
	fmt.Println(result)

	// Output:
	// <nil>
	// 42
}

func ExampleInterpreter_Evaluate_binary_star_error_left() {
	i := NewInterpreter()
	expr := expression.NewBinary(
		expression.NewLiteral(nil),
		*token.NewToken(token.STAR, "*", nil, 1),
		expression.NewLiteral(2),
	)
	result, err := i.Evaluate(expr)
	fmt.Println(err)
	fmt.Println(result)

	// Output:
	// RuntimeError Cannot convert to float64, unexpected type (ConversionError)
	// <nil>
}

func ExampleInterpreter_Evaluate_binary_star_error_right() {
	i := NewInterpreter()
	expr := expression.NewBinary(
		expression.NewLiteral(21),
		*token.NewToken(token.STAR, "*", nil, 1),
		expression.NewLiteral(nil),
	)
	result, err := i.Evaluate(expr)
	fmt.Println(err)
	fmt.Println(result)

	// Output:
	// RuntimeError Cannot convert to float64, unexpected type (ConversionError)
	// <nil>
}

func TestInterpreter_Plus_Numbers(t *testing.T) {
	scan := scanner.NewScanner("42 + 42")
	tokens, _ := scan.ScanTokens()
	parser := parser.NewParser(tokens)
	interpreter := NewInterpreter()

	actual := utils.CaptureStdout(t, func() {
		interpreter.Interpret(parser.Parse())
	})
	expected := "84\n"
	if actual != expected {
		t.Errorf("Expected '%v' but got '%v'", expected, actual)
	}
}

func ExampleInterpreter_Evaluate_binary_plus_numbers() {
	i := NewInterpreter()
	expr := expression.NewBinary(
		expression.NewLiteral(42),
		*token.NewToken(token.PLUS, "+", nil, 1),
		expression.NewLiteral(42),
	)
	result, err := i.Evaluate(expr)
	fmt.Println(err)
	fmt.Println(result)

	// Output:
	// <nil>
	// 84
}

func ExampleInterpreter_Evaluate_binary_plus_string() {
	i := NewInterpreter()
	expr := expression.NewBinary(
		expression.NewLiteral("hello "),
		*token.NewToken(token.PLUS, "+", nil, 1),
		expression.NewLiteral("world"),
	)
	result, err := i.Evaluate(expr)
	fmt.Println(err)
	fmt.Println(result)

	// Output:
	// <nil>
	// hello world
}

func ExampleInterpreter_Evaluate_binary_plus_incompatible_types() {
	i := NewInterpreter()
	expr := expression.NewBinary(
		expression.NewLiteral("42"),
		*token.NewToken(token.PLUS, "+", nil, 1),
		expression.NewLiteral(42),
	)
	result, err := i.Evaluate(expr)
	fmt.Println(err)
	fmt.Println(result)

	// Output:
	// RuntimeError Incompatible types in PLUS operation
	// <nil>
}

func ExampleInterpreter_Evaluate_binary_greater() {
	i := NewInterpreter()
	expr := expression.NewBinary(
		expression.NewLiteral(43),
		*token.NewToken(token.GREATER, ">", nil, 1),
		expression.NewLiteral(42),
	)
	result, err := i.Evaluate(expr)
	fmt.Println(err)
	fmt.Println(result)

	// Output:
	// <nil>
	// true
}

func ExampleInterpreter_Evaluate_binary_greater_error() {
	i := NewInterpreter()
	expr := expression.NewBinary(
		expression.NewLiteral(43),
		*token.NewToken(token.GREATER, ">", nil, 1),
		expression.NewLiteral(nil),
	)
	result, err := i.Evaluate(expr)
	fmt.Println(err)
	fmt.Println(result)

	// Output:
	// RuntimeError Cannot convert to float64, unexpected type (ConversionError)
	// <nil>
}

func ExampleInterpreter_Evaluate_binary_greater_equal() {
	i := NewInterpreter()
	expr := expression.NewBinary(
		expression.NewLiteral(42),
		*token.NewToken(token.GREATER_EQUAL, ">=", nil, 1),
		expression.NewLiteral(42),
	)
	result, err := i.Evaluate(expr)
	fmt.Println(err)
	fmt.Println(result)

	// Output:
	// <nil>
	// true
}

func ExampleInterpreter_Evaluate_binary_greater_equal_error() {
	i := NewInterpreter()
	expr := expression.NewBinary(
		expression.NewLiteral(42),
		*token.NewToken(token.GREATER_EQUAL, ">=", nil, 1),
		expression.NewLiteral(nil),
	)
	result, err := i.Evaluate(expr)
	fmt.Println(err)
	fmt.Println(result)

	// Output:
	// RuntimeError Cannot convert to float64, unexpected type (ConversionError)
	// <nil>
}

func ExampleInterpreter_Evaluate_binary_less() {
	i := NewInterpreter()
	expr := expression.NewBinary(
		expression.NewLiteral(41),
		*token.NewToken(token.LESS, "<", nil, 1),
		expression.NewLiteral(42),
	)
	result, err := i.Evaluate(expr)
	fmt.Println(err)
	fmt.Println(result)

	// Output:
	// <nil>
	// true
}

func ExampleInterpreter_Evaluate_binary_less_error() {
	i := NewInterpreter()
	expr := expression.NewBinary(
		expression.NewLiteral(41),
		*token.NewToken(token.LESS, "<", nil, 1),
		expression.NewLiteral(nil),
	)
	result, err := i.Evaluate(expr)
	fmt.Println(err)
	fmt.Println(result)

	// Output:
	// RuntimeError Cannot convert to float64, unexpected type (ConversionError)
	// <nil>
}

func ExampleInterpreter_Evaluate_binary_less_equal() {
	i := NewInterpreter()
	expr := expression.NewBinary(
		expression.NewLiteral(42),
		*token.NewToken(token.LESS_EQUAL, "<=", nil, 1),
		expression.NewLiteral(42),
	)
	result, err := i.Evaluate(expr)
	fmt.Println(err)
	fmt.Println(result)

	// Output:
	// <nil>
	// true
}

func ExampleInterpreter_Evaluate_binary_less_equal_error() {
	i := NewInterpreter()
	expr := expression.NewBinary(
		expression.NewLiteral(42),
		*token.NewToken(token.LESS_EQUAL, "<=", nil, 1),
		expression.NewLiteral(nil),
	)
	result, err := i.Evaluate(expr)
	fmt.Println(err)
	fmt.Println(result)

	// Output:
	// RuntimeError Cannot convert to float64, unexpected type (ConversionError)
	// <nil>
}

func ExampleInterpreter_Evaluate_binary_bang_equal() {
	i := NewInterpreter()
	expr := expression.NewBinary(
		expression.NewLiteral(43),
		*token.NewToken(token.BANQ_EQUAL, "!=", nil, 1),
		expression.NewLiteral(42),
	)
	result, err := i.Evaluate(expr)
	fmt.Println(err)
	fmt.Println(result)

	// Output:
	// <nil>
	// true
}

func ExampleInterpreter_Evaluate_binary_equal_equal() {
	i := NewInterpreter()
	expr := expression.NewBinary(
		expression.NewLiteral(42),
		*token.NewToken(token.EQUAL_EQUAL, "==", nil, 1),
		expression.NewLiteral(42),
	)
	result, err := i.Evaluate(expr)
	fmt.Println(err)
	fmt.Println(result)

	// Output:
	// <nil>
	// true
}

func ExampleInterpreter_Evaluate_binary_equal_both_nil() {
	i := NewInterpreter()
	expr := expression.NewBinary(
		expression.NewLiteral(nil),
		*token.NewToken(token.EQUAL_EQUAL, "==", nil, 1),
		expression.NewLiteral(nil),
	)
	result, err := i.Evaluate(expr)
	fmt.Println(err)
	fmt.Println(result)

	// Output:
	// <nil>
	// true
}

func ExampleInterpreter_Evaluate_binary_equal_one_nil() {
	i := NewInterpreter()
	expr := expression.NewBinary(
		expression.NewLiteral(nil),
		*token.NewToken(token.EQUAL_EQUAL, "==", nil, 1),
		expression.NewLiteral(42),
	)
	result, err := i.Evaluate(expr)
	fmt.Println(err)
	fmt.Println(result)

	// Output:
	// <nil>
	// false
}

func ExampleInterpreter_Evaluate_binary_equal_not_same_types() {
	i := NewInterpreter()
	expr := expression.NewBinary(
		expression.NewLiteral(42),
		*token.NewToken(token.EQUAL_EQUAL, "==", nil, 1),
		expression.NewLiteral(true),
	)
	result, err := i.Evaluate(expr)
	fmt.Println(err)
	fmt.Println(result)

	// Output:
	// <nil>
	// false
}

func ExampleInterpreter_Evaluate_binary_equal_false() {
	i := NewInterpreter()
	expr := expression.NewBinary(
		expression.NewLiteral(true),
		*token.NewToken(token.EQUAL_EQUAL, "==", nil, 1),
		expression.NewLiteral(false),
	)
	result, err := i.Evaluate(expr)
	fmt.Println(err)
	fmt.Println(result)

	// Output:
	// <nil>
	// false
}
