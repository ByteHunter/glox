package interpreter

import (
	"errors"
	"math"
	"reflect"

	"github.com/ByteHunter/glox/expression"
	"github.com/ByteHunter/glox/reporting"
	"github.com/ByteHunter/glox/token"
)

type Interpreter struct{}

func NewInterpreter() *Interpreter {
	return &Interpreter{}
}

func (i *Interpreter) VisitBinaryExpression(expr *expression.Binary) any {
	if expr.Left == nil {
		reporting.LoxError(
			expr.Operator.Line,
			"Left operand expected to be an expression, nil found (InterpreterError)",
		)
		return nil
	}
	if expr.Right == nil {
		reporting.LoxError(
			expr.Operator.Line,
			"Right operand expected to be an expression, nil found (InterpreterError)",
		)
		return nil
	}
	left := i.Evaluate(expr.Left)
	right := i.Evaluate(expr.Right)

	switch expr.Operator.Type {
	case token.MINUS:
		l, err := i.getFloat(left)
		if err != nil {
			reporting.LoxError(expr.Operator.Line, err.Error())
			return nil
		}
		r, err := i.getFloat(right)
		if err != nil {
			reporting.LoxError(expr.Operator.Line, err.Error())
			return nil
		}
		return l - r
	case token.SLASH:
		l, err := i.getFloat(left)
		if err != nil {
			reporting.LoxError(expr.Operator.Line, err.Error())
			return nil
		}
		r, err := i.getFloat(right)
		if err != nil {
			reporting.LoxError(expr.Operator.Line, err.Error())
			return nil
		}
		return l / r
	case token.STAR:
		l, err := i.getFloat(left)
		if err != nil {
			reporting.LoxError(expr.Operator.Line, err.Error())
			return nil
		}
		r, err := i.getFloat(right)
		if err != nil {
			reporting.LoxError(expr.Operator.Line, err.Error())
			return nil
		}
		return l * r
	case token.PLUS:
		left_type := reflect.TypeOf(left).String()
		right_type := reflect.TypeOf(right).String()
		if left_type == "int" && right_type == "int" {
			return float64(left.(int)) + float64(right.(int))
		}
		if left_type == "string" && right_type == "string" {
			l := left.(string)
			r := right.(string)
			return l + r
		}

		reporting.LoxError(expr.Operator.Line, "Incompatible types in PLUS operation (InterpreterError)")
		return nil
	}

	reporting.LoxError(expr.Operator.Line, "Unknown binary operator (InterpreterError)")
	return nil
}

func (i *Interpreter) VisitGroupingExpression(expr *expression.Grouping) any {
	return i.Evaluate(expr.Expr)
}

func (i *Interpreter) VisitLiteralExpression(expr *expression.Literal) any {
	return expr.Value
}

func (i *Interpreter) VisitUnaryExpression(expr *expression.Unary) any {
	if expr.Right == nil {
		reporting.LoxError(expr.Operator.Line, "Expected an expression, nil found (InterpreterError)")
		return nil
	}
	right := i.Evaluate(expr.Right)

	switch expr.Operator.Type {
	case token.BANG:
		return !i.getBoolean(right)
	case token.MINUS:
		res, err := i.getFloat(right)
		if err != nil {
			reporting.LoxError(expr.Operator.Line, err.Error())
			return res
		}
		return -res
	}

	reporting.LoxError(expr.Operator.Line, "Unknown unary operator (InterpreterError)")
	return nil
}

func (i *Interpreter) Evaluate(expr expression.Expression) any {
	return expr.Accept(i)
}

func (i *Interpreter) getFloat(v any) (float64, error) {
	switch t := v.(type) {
	case int:
		return float64(t), nil
	default:
		return math.NaN(), errors.New("Cannot convert to float64, unexpected type (ConversionError)")
	}
}

func (i *Interpreter) getBoolean(v any) bool {
	if v == nil || v == false {
		return false
	}
	return true
}
