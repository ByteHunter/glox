package interpreter

import (
	"errors"
	"fmt"
	"math"
	"testing"

	syntax_expression "github.com/ByteHunter/glox/syntax/expression"
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
		expr     *syntax_expression.Literal
		expected any
	}{
		{
			syntax_expression.NewLiteral(42),
			42,
		},
		{
			syntax_expression.NewLiteral(42.1),
			42.1,
		},
		{
			syntax_expression.NewLiteral(int(42)),
			int(42),
		},
		{
			syntax_expression.NewLiteral(float64(42)),
			float64(42),
		},
		{
			syntax_expression.NewLiteral(true),
			true,
		},
		{
			syntax_expression.NewLiteral("hello"),
			"hello",
		},
		{
			syntax_expression.NewLiteral(nil),
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
		expr           *syntax_expression.Unary
		expectedResult any
		expectedNaN    bool
		expectedError  error
	}{
		{
			syntax_expression.NewUnary(*minusToken, syntax_expression.NewLiteral(42)),
			float64(-42),
			false,
			nil,
		},
		{
			syntax_expression.NewUnary(*minusToken, syntax_expression.NewLiteral(-42)),
			float64(42),
			false,
			nil,
		},
		{
			syntax_expression.NewUnary(*minusToken, syntax_expression.NewLiteral(int(42))),
			float64(-42),
			false,
			nil,
		},
		{
			syntax_expression.NewUnary(*minusToken, syntax_expression.NewLiteral(float64(42))),
			float64(-42),
			false,
			nil,
		},
		{
			syntax_expression.NewUnary(*minusToken, nil),
			nil,
			false,
			NewRuntimeError(*minusToken, "Expected an expression, nil found"),
		},
		{
			syntax_expression.NewUnary(
				*minusToken,
				syntax_expression.NewUnary(*minusToken, syntax_expression.NewLiteral("42")),
			),
			nil,
			false,
			NewRuntimeError(*minusToken, "Cannot convert to float64, unexpected type (ConversionError)"),
		},
		{
			syntax_expression.NewUnary(*minusToken, syntax_expression.NewLiteral("42")),
			math.NaN(),
			true,
			NewRuntimeError(*minusToken, "Cannot convert to float64, unexpected type (ConversionError)"),
		},
		{
			syntax_expression.NewUnary(*plusToken, syntax_expression.NewLiteral(42)),
			nil,
			false,
			NewRuntimeError(*minusToken, "Unknown unary operator"),
		},
		{
			syntax_expression.NewUnary(*bangToken, syntax_expression.NewLiteral(true)),
			false,
			false,
			nil,
		},
		{
			syntax_expression.NewUnary(*bangToken, syntax_expression.NewLiteral(false)),
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
		syntax_expression.NewGrouping(
			syntax_expression.NewLiteral(42),
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
	literal2 := syntax_expression.NewLiteral(2)
	literal42 := syntax_expression.NewLiteral(42)
	literal84 := syntax_expression.NewLiteral(84)
	var tests = []struct {
		expr           *syntax_expression.Binary
		expectedResult any
		expectedNaN    bool
		expectedError  error
	}{
		// Early returns
		{
			syntax_expression.NewBinary(nil, *greaterToken, literal42),
			nil,
			false,
			NewRuntimeError(*greaterToken, "Left operand expected to be an expression, nil found"),
		},
		{
			syntax_expression.NewBinary(literal42, *greaterToken, nil),
			nil,
			false,
			NewRuntimeError(*greaterToken, "Right operand expected to be an expression, nil found"),
		},
		{
			syntax_expression.NewBinary(
				syntax_expression.NewUnary(*minusToken, syntax_expression.NewLiteral("42")),
				*minusToken,
				literal42,
			),
			nil,
			false,
			NewRuntimeError(*greaterToken, "Cannot convert to float64, unexpected type (ConversionError)"),
		},
		{
			syntax_expression.NewBinary(
				literal42,
				*minusToken,
				syntax_expression.NewUnary(*minusToken, syntax_expression.NewLiteral("42")),
			),
			nil,
			false,
			NewRuntimeError(*greaterToken, "Cannot convert to float64, unexpected type (ConversionError)"),
		},
		// Unknown operator
		{
			syntax_expression.NewBinary(literal42, *token.NewToken(token.BANG, "!", nil, 1), literal42),
			nil,
			false,
			NewRuntimeError(*greaterToken, "Unknown binary operator"),
		},
		// Greater
		{
			syntax_expression.NewBinary(literal42, *greaterToken, literal42),
			false, false, nil,
		},
		{
			syntax_expression.NewBinary(literal84, *greaterToken, literal42),
			true, false, nil,
		},
		{
			syntax_expression.NewBinary(literal42, *greaterToken, literal84),
			false, false, nil,
		},
		{
			syntax_expression.NewBinary(literal42, *greaterToken, syntax_expression.NewLiteral("42")),
			nil, false,
			NewRuntimeError(*greaterToken, "Cannot convert to float64, unexpected type (ConversionError)"),
		},
		// Greater Equal
		{
			syntax_expression.NewBinary(literal42, *greaterEqualToken, literal42),
			true, false, nil,
		},
		{
			syntax_expression.NewBinary(literal84, *greaterEqualToken, literal42),
			true, false, nil,
		},
		{
			syntax_expression.NewBinary(literal42, *greaterEqualToken, literal84),
			false, false, nil,
		},
		{
			syntax_expression.NewBinary(literal42, *greaterEqualToken, syntax_expression.NewLiteral("42")),
			nil, false,
			NewRuntimeError(*greaterEqualToken, "Cannot convert to float64, unexpected type (ConversionError)"),
		},
		// Less
		{
			syntax_expression.NewBinary(literal42, *lessToken, literal42),
			false, false, nil,
		},
		{
			syntax_expression.NewBinary(literal84, *lessToken, literal42),
			false, false, nil,
		},
		{
			syntax_expression.NewBinary(literal42, *lessToken, literal84),
			true, false, nil,
		},
		{
			syntax_expression.NewBinary(literal42, *lessToken, syntax_expression.NewLiteral("42")),
			nil, false,
			NewRuntimeError(*lessToken, "Cannot convert to float64, unexpected type (ConversionError)"),
		},
		// Less Equal
		{
			syntax_expression.NewBinary(literal42, *lessEqualToken, literal42),
			true, false, nil,
		},
		{
			syntax_expression.NewBinary(literal84, *lessEqualToken, literal42),
			false, false, nil,
		},
		{
			syntax_expression.NewBinary(literal42, *lessEqualToken, literal84),
			true, false, nil,
		},
		{
			syntax_expression.NewBinary(literal42, *lessEqualToken, syntax_expression.NewLiteral("42")),
			nil, false,
			NewRuntimeError(*lessEqualToken, "Cannot convert to float64, unexpected type (ConversionError)"),
		},
		// Bang Equal
		{
			syntax_expression.NewBinary(literal42, *bangEqualToken, literal42),
			false, false, nil,
		},
		{
			syntax_expression.NewBinary(literal84, *bangEqualToken, literal42),
			true, false, nil,
		},
		{
			syntax_expression.NewBinary(literal42, *bangEqualToken, literal84),
			true, false, nil,
		},
		// Equal Equal
		{
			syntax_expression.NewBinary(literal42, *equalEqualToken, literal42),
			true, false, nil,
		},
		{
			syntax_expression.NewBinary(literal84, *equalEqualToken, literal42),
			false, false, nil,
		},
		{
			syntax_expression.NewBinary(literal42, *equalEqualToken, literal84),
			false, false, nil,
		},
		// Minus
		{
			syntax_expression.NewBinary(literal84, *minusToken, literal42),
			float64(42), false, nil,
		},
		{
			syntax_expression.NewBinary(literal42, *minusToken, syntax_expression.NewLiteral("42")),
			nil, false,
			NewRuntimeError(*minusToken, "Cannot convert to float64, unexpected type (ConversionError)"),
		},
		// Slash
		{
			syntax_expression.NewBinary(literal84, *slashToken, literal2),
			float64(42), false, nil,
		},
		{
			syntax_expression.NewBinary(literal42, *slashToken, syntax_expression.NewLiteral("42")),
			nil, false,
			NewRuntimeError(*slashToken, "Cannot convert to float64, unexpected type (ConversionError)"),
		},
		// Star
		{
			syntax_expression.NewBinary(literal42, *starToken, literal2),
			float64(84), false, nil,
		},
		{
			syntax_expression.NewBinary(literal42, *starToken, syntax_expression.NewLiteral("42")),
			nil, false,
			NewRuntimeError(*starToken, "Cannot convert to float64, unexpected type (ConversionError)"),
		},
		// Plus
		{
			syntax_expression.NewBinary(literal42, *plusToken, literal42),
			float64(84), false, nil,
		},
		{
			syntax_expression.NewBinary(syntax_expression.NewLiteral(float64(42)), *plusToken, syntax_expression.NewLiteral(float64(42))),
			float64(84), false, nil,
		},
		{
			syntax_expression.NewBinary(syntax_expression.NewLiteral("hello "), *plusToken, syntax_expression.NewLiteral("world")),
			"hello world", false, nil,
		},
		{
			syntax_expression.NewBinary(syntax_expression.NewLiteral("hello "), *plusToken, syntax_expression.NewLiteral(float64(42))),
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
	expected := "[line 1] Error : RuntimeError Expected an expression, nil found\n"

	if expected != actual {
		t.Errorf("Expected '%v' but got '%v'", expected, actual)
	}
}

func TestInterpreter_Interpret_valid(t *testing.T) {
	actual := utils.CaptureStdout(t, func() {
		NewInterpreter().Interpret(syntax_expression.NewLiteral(42))
	})
	expected := "42\n"

	if expected != actual {
		t.Errorf("Expected '%v' but got '%v'", expected, actual)
	}
}
