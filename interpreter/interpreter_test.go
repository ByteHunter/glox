package interpreter

import (
	"fmt"

	"github.com/ByteHunter/glox/expression"
	"github.com/ByteHunter/glox/token"
)

// Evaluating Unary expressions

func ExampleInterpreter_Evaluate_unary_nil() {
	i := NewInterpreter()
	expr := expression.NewUnary(
		*token.NewToken(token.MINUS, "-", nil, 1),
		nil,
	)
	result := i.Evaluate(expr)
	fmt.Println(result)

	// Output:
	// [line 1] Error : Expected an expression, nil found (InterpreterError)
	// <nil>
}

func ExampleInterpreter_Evaluate_unary_unkown_token() {
	i := NewInterpreter()
	expr := expression.NewUnary(
		*token.NewToken(token.PLUS, "+", nil, 1),
		expression.NewLiteral(42),
	)
	result := i.Evaluate(expr)
	fmt.Println(result)

	// Output:
	// [line 1] Error : Unknown unary operator (InterpreterError)
	// <nil>
}

func ExampleInterpreter_Evaluate_unary_minus_nan() {
	i := NewInterpreter()
	expr := expression.NewUnary(
		*token.NewToken(token.MINUS, "-", nil, 1),
		expression.NewLiteral(true),
	)
	result := i.Evaluate(expr)
	fmt.Println(result)

	// Output:
	// [line 1] Error : Cannot convert to float64, unexpected type (ConversionError)
	// NaN
}

func ExampleInterpreter_Evaluate_unary_minus() {
	i := NewInterpreter()
	expr := expression.NewUnary(
		*token.NewToken(token.MINUS, "-", nil, 1),
		expression.NewLiteral(42),
	)
	result := i.Evaluate(expr)
	fmt.Println(result)

	// Output:
	// -42
}

func ExampleInterpreter_Evaluate_unary_minus2() {
	i := NewInterpreter()
	expr := expression.NewUnary(
		*token.NewToken(token.MINUS, "-", nil, 1),
		expression.NewLiteral(-42),
	)
	result := i.Evaluate(expr)
	fmt.Println(result)

	// Output:
	// 42
}

func ExampleInterpreter_Evaluate_unary_bang_true() {
	i := NewInterpreter()
	expr := expression.NewUnary(
		*token.NewToken(token.BANG, "!", nil, 1),
		expression.NewLiteral(true),
	)
	result := i.Evaluate(expr)
	fmt.Println(result)

	// Output:
	// false
}

func ExampleInterpreter_Evaluate_unary_bang_false() {
	i := NewInterpreter()
	expr := expression.NewUnary(
		*token.NewToken(token.BANG, "!", nil, 1),
		expression.NewLiteral(false),
	)
	result := i.Evaluate(expr)
	fmt.Println(result)

	// Output:
	// true
}

func ExampleInterpreter_Evaluate_unary_bang_other() {
	i := NewInterpreter()
	expr := expression.NewUnary(
		*token.NewToken(token.BANG, "!", nil, 1),
		expression.NewLiteral(42),
	)
	result := i.Evaluate(expr)
	fmt.Println(result)

	// Output:
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
	result := i.Evaluate(expr)
	fmt.Println(result)

	// Output:
	// [line 1] Error : Unknown binary operator (InterpreterError)
	// <nil>
}

func ExampleInterpreter_Evaluate_binary_missing_left_operand() {
	i := NewInterpreter()
	expr := expression.NewBinary(
		nil,
		*token.NewToken(token.MINUS, "-", nil, 1),
		expression.NewLiteral(1),
	)
	result := i.Evaluate(expr)
	fmt.Println(result)

	// Output:
	// [line 1] Error : Left operand expected to be an expression, nil found (InterpreterError)
	// <nil>
}

func ExampleInterpreter_Evaluate_binary_missing_right_operand() {
	i := NewInterpreter()
	expr := expression.NewBinary(
		expression.NewLiteral(42),
		*token.NewToken(token.MINUS, "-", nil, 1),
		nil,
	)
	result := i.Evaluate(expr)
	fmt.Println(result)

	// Output:
	// [line 1] Error : Right operand expected to be an expression, nil found (InterpreterError)
	// <nil>
}

func ExampleInterpreter_Evaluate_binary_minus() {
	i := NewInterpreter()
	expr := expression.NewBinary(
		expression.NewLiteral(43),
		*token.NewToken(token.MINUS, "-", nil, 1),
		expression.NewLiteral(1),
	)
	result := i.Evaluate(expr)
	fmt.Println(result)

	// Output:
	// 42
}

func ExampleInterpreter_Evaluate_binary_minus_error_left() {
	i := NewInterpreter()
	expr := expression.NewBinary(
		expression.NewLiteral(43),
		*token.NewToken(token.MINUS, "-", nil, 1),
		expression.NewLiteral(nil),
	)
	result := i.Evaluate(expr)
	fmt.Println(result)

	// Output:
	// [line 1] Error : Cannot convert to float64, unexpected type (ConversionError)
	// <nil>
}

func ExampleInterpreter_Evaluate_binary_minus_error_right() {
	i := NewInterpreter()
	expr := expression.NewBinary(
		expression.NewLiteral(nil),
		*token.NewToken(token.MINUS, "-", nil, 1),
		expression.NewLiteral(1),
	)
	result := i.Evaluate(expr)
	fmt.Println(result)

	// Output:
	// [line 1] Error : Cannot convert to float64, unexpected type (ConversionError)
	// <nil>
}

func ExampleInterpreter_Evaluate_binary_slash() {
	i := NewInterpreter()
	expr := expression.NewBinary(
		expression.NewLiteral(42),
		*token.NewToken(token.SLASH, "/", nil, 1),
		expression.NewLiteral(2),
	)
	result := i.Evaluate(expr)
	fmt.Println(result)

	// Output:
	// 21
}

func ExampleInterpreter_Evaluate_binary_slash_error_left() {
	i := NewInterpreter()
	expr := expression.NewBinary(
		expression.NewLiteral(nil),
		*token.NewToken(token.SLASH, "/", nil, 1),
		expression.NewLiteral(2),
	)
	result := i.Evaluate(expr)
	fmt.Println(result)

	// Output:
	// [line 1] Error : Cannot convert to float64, unexpected type (ConversionError)
	// <nil>
}

func ExampleInterpreter_Evaluate_binary_slash_error_right() {
	i := NewInterpreter()
	expr := expression.NewBinary(
		expression.NewLiteral(42),
		*token.NewToken(token.SLASH, "/", nil, 1),
		expression.NewLiteral(nil),
	)
	result := i.Evaluate(expr)
	fmt.Println(result)

	// Output:
	// [line 1] Error : Cannot convert to float64, unexpected type (ConversionError)
	// <nil>
}

func ExampleInterpreter_Evaluate_binary_star() {
	i := NewInterpreter()
	expr := expression.NewBinary(
		expression.NewLiteral(21),
		*token.NewToken(token.STAR, "*", nil, 1),
		expression.NewLiteral(2),
	)
	result := i.Evaluate(expr)
	fmt.Println(result)

	// Output:
	// 42
}

func ExampleInterpreter_Evaluate_binary_star_error_left() {
	i := NewInterpreter()
	expr := expression.NewBinary(
		expression.NewLiteral(nil),
		*token.NewToken(token.STAR, "*", nil, 1),
		expression.NewLiteral(2),
	)
	result := i.Evaluate(expr)
	fmt.Println(result)

	// Output:
	// [line 1] Error : Cannot convert to float64, unexpected type (ConversionError)
	// <nil>
}

func ExampleInterpreter_Evaluate_binary_star_error_right() {
	i := NewInterpreter()
	expr := expression.NewBinary(
		expression.NewLiteral(21),
		*token.NewToken(token.STAR, "*", nil, 1),
		expression.NewLiteral(nil),
	)
	result := i.Evaluate(expr)
	fmt.Println(result)

	// Output:
	// [line 1] Error : Cannot convert to float64, unexpected type (ConversionError)
	// <nil>
}

func ExampleInterpreter_Evaluate_binary_plus_numbers() {
	i := NewInterpreter()
	expr := expression.NewBinary(
		expression.NewLiteral(42),
		*token.NewToken(token.PLUS, "+", nil, 1),
		expression.NewLiteral(42),
	)
	result := i.Evaluate(expr)
	fmt.Println(result)

	// Output:
	// 84
}

func ExampleInterpreter_Evaluate_binary_plus_string() {
	i := NewInterpreter()
	expr := expression.NewBinary(
		expression.NewLiteral("hello "),
		*token.NewToken(token.PLUS, "+", nil, 1),
		expression.NewLiteral("world"),
	)
	result := i.Evaluate(expr)
	fmt.Println(result)

	// Output:
	// hello world
}

func ExampleInterpreter_Evaluate_binary_plus_incompatible_types() {
	i := NewInterpreter()
	expr := expression.NewBinary(
		expression.NewLiteral("42"),
		*token.NewToken(token.PLUS, "+", nil, 1),
		expression.NewLiteral(42),
	)
	result := i.Evaluate(expr)
	fmt.Println(result)

	// Output:
	// [line 1] Error : Incompatible types in PLUS operation (InterpreterError)
	// <nil>
}

func ExampleInterpreter_Evaluate_binary_greater() {
	i := NewInterpreter()
	expr := expression.NewBinary(
		expression.NewLiteral(43),
		*token.NewToken(token.GREATER, ">", nil, 1),
		expression.NewLiteral(42),
	)
	result := i.Evaluate(expr)
	fmt.Println(result)

	// Output:
	// true
}

func ExampleInterpreter_Evaluate_binary_greater_error() {
	i := NewInterpreter()
	expr := expression.NewBinary(
		expression.NewLiteral(43),
		*token.NewToken(token.GREATER, ">", nil, 1),
		expression.NewLiteral(nil),
	)
	result := i.Evaluate(expr)
	fmt.Println(result)

	// Output:
	// [line 1] Error : Cannot convert to float64, unexpected type (ConversionError)
	// <nil>
}

func ExampleInterpreter_Evaluate_binary_greater_equal() {
	i := NewInterpreter()
	expr := expression.NewBinary(
		expression.NewLiteral(42),
		*token.NewToken(token.GREATER_EQUAL, ">=", nil, 1),
		expression.NewLiteral(42),
	)
	result := i.Evaluate(expr)
	fmt.Println(result)

	// Output:
	// true
}

func ExampleInterpreter_Evaluate_binary_greater_equal_error() {
	i := NewInterpreter()
	expr := expression.NewBinary(
		expression.NewLiteral(42),
		*token.NewToken(token.GREATER_EQUAL, ">=", nil, 1),
		expression.NewLiteral(nil),
	)
	result := i.Evaluate(expr)
	fmt.Println(result)

	// Output:
	// [line 1] Error : Cannot convert to float64, unexpected type (ConversionError)
	// <nil>
}

func ExampleInterpreter_Evaluate_binary_less() {
	i := NewInterpreter()
	expr := expression.NewBinary(
		expression.NewLiteral(41),
		*token.NewToken(token.LESS, "<", nil, 1),
		expression.NewLiteral(42),
	)
	result := i.Evaluate(expr)
	fmt.Println(result)

	// Output:
	// true
}

func ExampleInterpreter_Evaluate_binary_less_error() {
	i := NewInterpreter()
	expr := expression.NewBinary(
		expression.NewLiteral(41),
		*token.NewToken(token.LESS, "<", nil, 1),
		expression.NewLiteral(nil),
	)
	result := i.Evaluate(expr)
	fmt.Println(result)

	// Output:
	// [line 1] Error : Cannot convert to float64, unexpected type (ConversionError)
	// <nil>
}

func ExampleInterpreter_Evaluate_binary_less_equal() {
	i := NewInterpreter()
	expr := expression.NewBinary(
		expression.NewLiteral(42),
		*token.NewToken(token.LESS_EQUAL, "<=", nil, 1),
		expression.NewLiteral(42),
	)
	result := i.Evaluate(expr)
	fmt.Println(result)

	// Output:
	// true
}

func ExampleInterpreter_Evaluate_binary_less_equal_error() {
	i := NewInterpreter()
	expr := expression.NewBinary(
		expression.NewLiteral(42),
		*token.NewToken(token.LESS_EQUAL, "<=", nil, 1),
		expression.NewLiteral(nil),
	)
	result := i.Evaluate(expr)
	fmt.Println(result)

	// Output:
	// [line 1] Error : Cannot convert to float64, unexpected type (ConversionError)
	// <nil>
}
