package interpreter

import (
	"errors"
	"fmt"
	"github.com/ByteHunter/glox/expression"
	"math"
	"testing"
	// "github.com/ByteHunter/glox/parser"
	// "github.com/ByteHunter/glox/scanner"
	"github.com/ByteHunter/glox/token"
	// "github.com/ByteHunter/glox/utils"
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
