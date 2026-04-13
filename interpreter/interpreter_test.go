package interpreter

import (
	"errors"
	"fmt"
	"math"
	"testing"

	"github.com/ByteHunter/glox/syntax/expression"
	"github.com/ByteHunter/glox/syntax/statement"
	"github.com/ByteHunter/glox/token"
	"github.com/ByteHunter/glox/utils"
)

func TestInterpreter_getFloat(t *testing.T) {
	var tests = []struct {
		value          any
		expectedResult any
		expectedNaN    bool
		expectedError  error
	}{
		{
			42, float64(42), false,
			nil,
		},
		{
			int(42), float64(42), false,
			nil,
		},
		{
			float64(42), float64(42), false,
			nil,
		},
		{
			nil, math.NaN(), true,
			errors.New("Cannot convert to float64, unexpected type (ConversionError)"),
		},
	}

	for set, test := range tests {
		testName := fmt.Sprintf("#%d", set)
		t.Run(testName, func(t *testing.T) {
			actual, err := NewInterpreter().getFloat(test.value)
			if !test.expectedNaN && actual != test.expectedResult {
				t.Errorf("Expected '%v' but got '%v'", test.expectedResult, actual)
			}
			if test.expectedNaN && !math.IsNaN(actual) {
				t.Errorf("Expected a NaN but got a number: %v", actual)
			}
			if test.expectedError == nil && err != nil {
				t.Errorf("Expected error to be nil, but got an error: '%v'", err)
			}
			if test.expectedError != nil && err != nil && test.expectedError.Error() != err.Error() {
				t.Errorf("Expected '%v' but got '%v'", test.expectedError, err)
			}
		})
	}
}

func TestInterpreter_parseTwoNumbers(t *testing.T) {
	var tests = []struct {
		left, right                             any
		expectedResultLeft, expectedResultRight any
		expectedNaNLeft, expectedNaNRight       bool
		expectedError                           error
	}{
		{
			42, 42,
			float64(42), float64(42),
			false, false,
			nil,
		},
		{
			nil, 42,
			math.NaN(), float64(42),
			true, false,
			errors.New("Cannot convert to float64, unexpected type (ConversionError)"),
		},
		{
			42, nil,
			float64(42), math.NaN(),
			false, true,
			errors.New("Cannot convert to float64, unexpected type (ConversionError)"),
		},
	}

	for set, test := range tests {
		testName := fmt.Sprintf("#%d", set)
		t.Run(testName, func(t *testing.T) {
			actualLeft, actualRight, err := NewInterpreter().parseTwoNumbers(test.left, test.right)
			if !test.expectedNaNLeft && actualLeft != test.expectedResultLeft {
				t.Errorf("Expected '%v' but got '%v'", test.expectedResultLeft, actualLeft)
			}
			if !test.expectedNaNRight && actualRight != test.expectedResultRight {
				t.Errorf("Expected '%v' but got '%v'", test.expectedResultRight, actualRight)
			}
			if test.expectedNaNLeft && !math.IsNaN(actualLeft) {
				t.Errorf("Expected a NaN but got a number: %v", actualLeft)
			}
			if test.expectedNaNRight && !math.IsNaN(actualRight) {
				t.Errorf("Expected a NaN but got a number: %v", actualRight)
			}
			if test.expectedError == nil && err != nil {
				t.Errorf("Expected error to be nil, but got an error: '%v'", err)
			}
			if test.expectedError != nil && err != nil && test.expectedError.Error() != err.Error() {
				t.Errorf("Expected '%v' but got '%v'", test.expectedError, err)
			}
		})
	}
}

func TestInterpreter_getBoolean(t *testing.T) {
	var tests = []struct {
		value    any
		expected any
	}{
		{true, true},
		{false, false},
		{nil, false},
		{42, true},
		{"true", true},
	}

	for set, test := range tests {
		testName := fmt.Sprintf("#%d", set)
		t.Run(testName, func(t *testing.T) {
			actual := NewInterpreter().getBoolean(test.value)
			if actual != test.expected {
				t.Errorf("Expected '%v' but got '%v'", test.expected, actual)
			}
		})
	}
}

func TestInterpreter_isEqual(t *testing.T) {
	var tests = []struct {
		a, b     any
		expected any
	}{
		{nil, nil, true},
		{nil, 42, false},
		{42, 42, true},
		{int(42), float64(42), false},
		{42, "42", false},
		{"42", "42", true},
		{"42", "0", false},
		{true, true, true},
		{true, false, false},
	}

	for set, test := range tests {
		testName := fmt.Sprintf("#%d", set)
		t.Run(testName, func(t *testing.T) {
			actual := NewInterpreter().isEqual(test.a, test.b)
			if actual != test.expected {
				t.Errorf("Expected '%v' but got '%v'", test.expected, actual)
			}
		})
	}
}

func TestInterpreter_VisitLiteral(t *testing.T) {
	var tests = []struct {
		expr     *expression.Literal
		expected any
	}{
		{
			expression.NewLiteral(42),
			42,
		},
		{
			expression.NewLiteral(42.1),
			42.1,
		},
		{
			expression.NewLiteral(int(42)),
			int(42),
		},
		{
			expression.NewLiteral(float64(42)),
			float64(42),
		},
		{
			expression.NewLiteral(true),
			true,
		},
		{
			expression.NewLiteral("hello"),
			"hello",
		},
		{
			expression.NewLiteral(nil),
			nil,
		},
	}

	for set, test := range tests {
		testName := fmt.Sprintf("#%d", set)
		t.Run(testName, func(t *testing.T) {
			actual, _ := NewInterpreter().VisitLiteralExpression(
				test.expr,
			)
			if actual != test.expected {
				t.Errorf("Expected '%v' but got '%v'", test.expected, actual)
			}
		})
	}
}

func TestInterpreter_VisitUnary(t *testing.T) {
	minusToken := token.NewToken(token.MINUS, "-", nil, 1)
	plusToken := token.NewToken(token.PLUS, "+", nil, 1)
	bangToken := token.NewToken(token.BANG, "!", nil, 1)
	var tests = []struct {
		expr           *expression.Unary
		expectedResult any
		expectedNaN    bool
		expectedError  error
	}{
		{
			expression.NewUnary(*minusToken, expression.NewLiteral(42)),
			float64(-42),
			false,
			nil,
		},
		{
			expression.NewUnary(*minusToken, expression.NewLiteral(-42)),
			float64(42),
			false,
			nil,
		},
		{
			expression.NewUnary(*minusToken, expression.NewLiteral(int(42))),
			float64(-42),
			false,
			nil,
		},
		{
			expression.NewUnary(*minusToken, expression.NewLiteral(float64(42))),
			float64(-42),
			false,
			nil,
		},
		{
			expression.NewUnary(*minusToken, nil),
			nil,
			false,
			NewRuntimeError(*minusToken, "Expected an expression, nil found"),
		},
		{
			expression.NewUnary(
				*minusToken,
				expression.NewUnary(*minusToken, expression.NewLiteral("42")),
			),
			nil,
			false,
			NewRuntimeError(*minusToken, "Cannot convert to float64, unexpected type (ConversionError)"),
		},
		{
			expression.NewUnary(*minusToken, expression.NewLiteral("42")),
			math.NaN(),
			true,
			NewRuntimeError(*minusToken, "Cannot convert to float64, unexpected type (ConversionError)"),
		},
		{
			expression.NewUnary(*plusToken, expression.NewLiteral(42)),
			nil,
			false,
			NewRuntimeError(*minusToken, "Unknown unary operator"),
		},
		{
			expression.NewUnary(*bangToken, expression.NewLiteral(true)),
			false,
			false,
			nil,
		},
		{
			expression.NewUnary(*bangToken, expression.NewLiteral(false)),
			true,
			false,
			nil,
		},
	}

	for set, test := range tests {
		testName := fmt.Sprintf("#%d", set)
		t.Run(testName, func(t *testing.T) {
			actual, err := NewInterpreter().VisitUnaryExpression(
				test.expr,
			)
			if !test.expectedNaN && actual != test.expectedResult {
				t.Errorf("Expected '%v' but got '%v'", test.expectedResult, actual)
			}
			if test.expectedNaN && !math.IsNaN(actual.(float64)) {
				t.Errorf("Expected '%v' but got '%v'", test.expectedResult, actual)
			}
			if test.expectedError != nil && err == nil {
				t.Errorf("Expected the error to be '%s', but got nil", test.expectedError)
			}
			if test.expectedError == nil && err != nil {
				t.Errorf("Expected error to be nil, but got an error: '%v'", err)
			}
			if test.expectedError != nil && err != nil && test.expectedError.Error() != err.Error() {
				t.Errorf("Expected '%v' but got '%v'", test.expectedError, err)
			}
		})
	}
}

func TestInterpreter_VisitGrouping(t *testing.T) {
	actual, err := NewInterpreter().VisitGroupingExpression(
		expression.NewGrouping(
			expression.NewLiteral(42),
		),
	)
	expected := 42
	if actual != expected {
		t.Errorf("Expected '%v' but got '%v'", expected, actual)
	}
	if err != nil {
		t.Errorf("Expected error to be nil, but got an error: '%v'", err)
	}
}

func TestInterpreter_VisitBinary(t *testing.T) {
	greaterToken := token.NewToken(token.GREATER, ">", nil, 1)
	greaterEqualToken := token.NewToken(token.GREATER_EQUAL, ">=", nil, 1)
	lessToken := token.NewToken(token.LESS, "<", nil, 1)
	lessEqualToken := token.NewToken(token.LESS_EQUAL, "<=", nil, 1)
	bangEqualToken := token.NewToken(token.BANG_EQUAL, "!=", nil, 1)
	equalEqualToken := token.NewToken(token.EQUAL_EQUAL, "==", nil, 1)
	minusToken := token.NewToken(token.MINUS, "-", nil, 1)
	slashToken := token.NewToken(token.SLASH, "/", nil, 1)
	starToken := token.NewToken(token.STAR, "*", nil, 1)
	plusToken := token.NewToken(token.PLUS, "-", nil, 1)
	literal2 := expression.NewLiteral(2)
	literal42 := expression.NewLiteral(42)
	literal84 := expression.NewLiteral(84)
	var tests = []struct {
		expr           *expression.Binary
		expectedResult any
		expectedNaN    bool
		expectedError  error
	}{
		// Early returns
		{
			expression.NewBinary(nil, *greaterToken, literal42),
			nil,
			false,
			NewRuntimeError(*greaterToken, "Left operand expected to be an expression, nil found"),
		},
		{
			expression.NewBinary(literal42, *greaterToken, nil),
			nil,
			false,
			NewRuntimeError(*greaterToken, "Right operand expected to be an expression, nil found"),
		},
		{
			expression.NewBinary(
				expression.NewUnary(*minusToken, expression.NewLiteral("42")),
				*minusToken,
				literal42,
			),
			nil,
			false,
			NewRuntimeError(*greaterToken, "Cannot convert to float64, unexpected type (ConversionError)"),
		},
		{
			expression.NewBinary(
				literal42,
				*minusToken,
				expression.NewUnary(*minusToken, expression.NewLiteral("42")),
			),
			nil,
			false,
			NewRuntimeError(*greaterToken, "Cannot convert to float64, unexpected type (ConversionError)"),
		},
		// Unknown operator
		{
			expression.NewBinary(literal42, *token.NewToken(token.BANG, "!", nil, 1), literal42),
			nil,
			false,
			NewRuntimeError(*greaterToken, "Unknown binary operator"),
		},
		// Greater
		{
			expression.NewBinary(literal42, *greaterToken, literal42),
			false, false, nil,
		},
		{
			expression.NewBinary(literal84, *greaterToken, literal42),
			true, false, nil,
		},
		{
			expression.NewBinary(literal42, *greaterToken, literal84),
			false, false, nil,
		},
		{
			expression.NewBinary(literal42, *greaterToken, expression.NewLiteral("42")),
			nil, false,
			NewRuntimeError(*greaterToken, "Cannot convert to float64, unexpected type (ConversionError)"),
		},
		// Greater Equal
		{
			expression.NewBinary(literal42, *greaterEqualToken, literal42),
			true, false, nil,
		},
		{
			expression.NewBinary(literal84, *greaterEqualToken, literal42),
			true, false, nil,
		},
		{
			expression.NewBinary(literal42, *greaterEqualToken, literal84),
			false, false, nil,
		},
		{
			expression.NewBinary(literal42, *greaterEqualToken, expression.NewLiteral("42")),
			nil, false,
			NewRuntimeError(*greaterEqualToken, "Cannot convert to float64, unexpected type (ConversionError)"),
		},
		// Less
		{
			expression.NewBinary(literal42, *lessToken, literal42),
			false, false, nil,
		},
		{
			expression.NewBinary(literal84, *lessToken, literal42),
			false, false, nil,
		},
		{
			expression.NewBinary(literal42, *lessToken, literal84),
			true, false, nil,
		},
		{
			expression.NewBinary(literal42, *lessToken, expression.NewLiteral("42")),
			nil, false,
			NewRuntimeError(*lessToken, "Cannot convert to float64, unexpected type (ConversionError)"),
		},
		// Less Equal
		{
			expression.NewBinary(literal42, *lessEqualToken, literal42),
			true, false, nil,
		},
		{
			expression.NewBinary(literal84, *lessEqualToken, literal42),
			false, false, nil,
		},
		{
			expression.NewBinary(literal42, *lessEqualToken, literal84),
			true, false, nil,
		},
		{
			expression.NewBinary(literal42, *lessEqualToken, expression.NewLiteral("42")),
			nil, false,
			NewRuntimeError(*lessEqualToken, "Cannot convert to float64, unexpected type (ConversionError)"),
		},
		// Bang Equal
		{
			expression.NewBinary(literal42, *bangEqualToken, literal42),
			false, false, nil,
		},
		{
			expression.NewBinary(literal84, *bangEqualToken, literal42),
			true, false, nil,
		},
		{
			expression.NewBinary(literal42, *bangEqualToken, literal84),
			true, false, nil,
		},
		// Equal Equal
		{
			expression.NewBinary(literal42, *equalEqualToken, literal42),
			true, false, nil,
		},
		{
			expression.NewBinary(literal84, *equalEqualToken, literal42),
			false, false, nil,
		},
		{
			expression.NewBinary(literal42, *equalEqualToken, literal84),
			false, false, nil,
		},
		// Minus
		{
			expression.NewBinary(literal84, *minusToken, literal42),
			float64(42), false, nil,
		},
		{
			expression.NewBinary(literal42, *minusToken, expression.NewLiteral("42")),
			nil, false,
			NewRuntimeError(*minusToken, "Cannot convert to float64, unexpected type (ConversionError)"),
		},
		// Slash
		{
			expression.NewBinary(literal84, *slashToken, literal2),
			float64(42), false, nil,
		},
		{
			expression.NewBinary(literal42, *slashToken, expression.NewLiteral("42")),
			nil, false,
			NewRuntimeError(*slashToken, "Cannot convert to float64, unexpected type (ConversionError)"),
		},
		// Star
		{
			expression.NewBinary(literal42, *starToken, literal2),
			float64(84), false, nil,
		},
		{
			expression.NewBinary(literal42, *starToken, expression.NewLiteral("42")),
			nil, false,
			NewRuntimeError(*starToken, "Cannot convert to float64, unexpected type (ConversionError)"),
		},
		// Plus
		{
			expression.NewBinary(literal42, *plusToken, literal42),
			float64(84), false, nil,
		},
		{
			expression.NewBinary(expression.NewLiteral(float64(42)), *plusToken, expression.NewLiteral(float64(42))),
			float64(84), false, nil,
		},
		{
			expression.NewBinary(expression.NewLiteral("hello "), *plusToken, expression.NewLiteral("world")),
			"hello world", false, nil,
		},
		{
			expression.NewBinary(expression.NewLiteral("hello "), *plusToken, expression.NewLiteral(float64(42))),
			nil, false,
			NewRuntimeError(*plusToken, "Incompatible types in PLUS operation"),
		},
	}

	for set, test := range tests {
		testName := fmt.Sprintf("#%d", set)
		t.Run(testName, func(t *testing.T) {
			actual, err := NewInterpreter().VisitBinaryExpression(
				test.expr,
			)
			if !test.expectedNaN && actual != test.expectedResult {
				t.Errorf("Expected '%v' but got '%v'", test.expectedResult, actual)
			}
			if test.expectedNaN && !math.IsNaN(actual.(float64)) {
				t.Errorf("Expected '%v' but got '%v'", test.expectedResult, actual)
			}
			if test.expectedError != nil && err == nil {
				t.Errorf("Expected the error to be '%s', but got nil", test.expectedError)
			}
			if test.expectedError == nil && err != nil {
				t.Errorf("Expected error to be nil, but got an error: '%v'", err)
			}
			if test.expectedError != nil && err != nil && test.expectedError.Error() != err.Error() {
				t.Errorf("Expected '%v' but got '%v'", test.expectedError, err)
			}
		})
	}
}

func TestInterpreter_Interpret_nil(t *testing.T) {
	actual := utils.CaptureStdout(t, func() {
		NewInterpreter().Interpret(nil)
	})
	expected := ""

	if expected != actual {
		t.Errorf("Expected '%v' but got '%v'", expected, actual)
	}
}

func TestInterpreter_Interpret_valid(t *testing.T) {
	var statements []statement.Statement = []statement.Statement{
		statement.NewPrint(expression.NewLiteral(true)),
	}
	actual := utils.CaptureStdout(t, func() {
		NewInterpreter().Interpret(statements)
	})
	expected := "true\n"

	if expected != actual {
		t.Errorf("Expected '%v' but got '%v'", expected, actual)
	}
}
