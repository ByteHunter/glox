package interpreter

import (
	"errors"
	"math"

	"github.com/ByteHunter/glox/expression"
	"github.com/ByteHunter/glox/reporting"
	"github.com/ByteHunter/glox/token"
)

type Interpreter struct{}

func NewInterpreter() *Interpreter {
	return &Interpreter{}
}

func (i *Interpreter) VisitBinaryExpression(expr *expression.Binary) any {
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
