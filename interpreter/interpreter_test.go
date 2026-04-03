package interpreter

import (
	"fmt"

	"github.com/ByteHunter/glox/expression"
	"github.com/ByteHunter/glox/token"
)

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
