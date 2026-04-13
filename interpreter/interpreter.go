package interpreter

import (
	"errors"
	"fmt"
	"math"
	"reflect"

	"github.com/ByteHunter/glox/reporting"
	"github.com/ByteHunter/glox/syntax/expression"
	"github.com/ByteHunter/glox/syntax/statement"
	"github.com/ByteHunter/glox/token"
)

type RuntimeError struct {
	Operator token.Token
	message  string
}

func NewRuntimeError(operator token.Token, message string) *RuntimeError {
	return &RuntimeError{
		Operator: operator,
		message:  message,
	}
}

func (r RuntimeError) Error() string {
	return fmt.Sprintf("RuntimeError %s", r.message)
}

type Interpreter struct{}

func NewInterpreter() *Interpreter {
	return &Interpreter{}
}

func (i *Interpreter) VisitExpressionStatement(stmt *statement.ExpressionStatement) (any, error) {
	i.Evaluate(stmt.Expr)
	return nil, nil
}

func (i *Interpreter) VisitPrintStatement(stmt *statement.PrintStatement) (any, error) {
	value, _ := i.Evaluate(stmt.Expr)
	fmt.Println(value)

	return nil, nil
}

func (i *Interpreter) VisitVariableStatement(stmt *statement.VariableStatement) (any, error) {
	return nil, NewRuntimeError(stmt.Name, "VariableStatement not implemented!")
}

func (i *Interpreter) VisitBinaryExpression(expr *expression.Binary) (any, error) {
	if expr.Left == nil {
		return nil, NewRuntimeError(
			expr.Operator,
			"Left operand expected to be an expression, nil found",
		)
	}
	if expr.Right == nil {
		return nil, NewRuntimeError(
			expr.Operator,
			"Right operand expected to be an expression, nil found",
		)
	}
	left, err := i.Evaluate(expr.Left)
	if err != nil {
		return nil, err
	}
	right, err := i.Evaluate(expr.Right)
	if err != nil {
		return nil, err
	}

	switch expr.Operator.Type {
	case token.GREATER:
		l, r, err := i.parseTwoNumbers(left, right)
		if err != nil {
			return nil, NewRuntimeError(expr.Operator, err.Error())
		}
		return l > r, nil
	case token.GREATER_EQUAL:
		l, r, err := i.parseTwoNumbers(left, right)
		if err != nil {
			return nil, NewRuntimeError(expr.Operator, err.Error())
		}
		return l >= r, nil
	case token.LESS:
		l, r, err := i.parseTwoNumbers(left, right)
		if err != nil {
			return nil, NewRuntimeError(expr.Operator, err.Error())
		}
		return l < r, nil
	case token.LESS_EQUAL:
		l, r, err := i.parseTwoNumbers(left, right)
		if err != nil {
			return nil, NewRuntimeError(expr.Operator, err.Error())
		}
		return l <= r, nil
	case token.BANG_EQUAL:
		return !i.isEqual(left, right), nil
	case token.EQUAL_EQUAL:
		return i.isEqual(left, right), nil
	case token.MINUS:
		l, r, err := i.parseTwoNumbers(left, right)
		if err != nil {
			return nil, NewRuntimeError(expr.Operator, err.Error())
		}
		return l - r, nil
	case token.SLASH:
		l, r, err := i.parseTwoNumbers(left, right)
		if err != nil {
			return nil, NewRuntimeError(expr.Operator, err.Error())
		}
		return l / r, nil
	case token.STAR:
		l, r, err := i.parseTwoNumbers(left, right)
		if err != nil {
			return nil, NewRuntimeError(expr.Operator, err.Error())
		}
		return l * r, nil
	case token.PLUS:
		left_type := reflect.TypeOf(left).String()
		right_type := reflect.TypeOf(right).String()

		if left_type == "string" && right_type == "string" {
			l := left.(string)
			r := right.(string)
			return l + r, nil
		}

		if left_type == "float64" && right_type == "float64" {
			return left.(float64) + right.(float64), nil
		}
		if left_type == "int" && right_type == "int" {
			return float64(left.(int)) + float64(right.(int)), nil
		}

		return nil, NewRuntimeError(expr.Operator, "Incompatible types in PLUS operation")
	}

	return nil, NewRuntimeError(expr.Operator, "Unknown binary operator")
}

func (i *Interpreter) VisitGroupingExpression(expr *expression.Grouping) (any, error) {
	return i.Evaluate(expr.Expr)
}

func (i *Interpreter) VisitLiteralExpression(expr *expression.Literal) (any, error) {
	return expr.Value, nil
}

func (i *Interpreter) VisitUnaryExpression(expr *expression.Unary) (any, error) {
	if expr.Right == nil {
		return nil, NewRuntimeError(expr.Operator, "Expected an expression, nil found")
	}
	right, err := i.Evaluate(expr.Right)
	if err != nil {
		return nil, err
	}

	switch expr.Operator.Type {
	case token.BANG:
		return !i.getBoolean(right), nil
	case token.MINUS:
		res, err := i.getFloat(right)
		if err != nil {
			return res, NewRuntimeError(expr.Operator, err.Error())
		}
		return -res, nil
	}

	return nil, NewRuntimeError(expr.Operator, "Unknown unary operator")
}

func (i *Interpreter) Evaluate(expr expression.Expression) (any, error) {
	if expr == nil {
		return nil, NewRuntimeError(token.Token{}, "Expected an expression, nil found")
	}
	return expr.Accept(i)
}

func (i *Interpreter) Interpret(stmts []statement.Statement) {
	for _, stmt := range stmts {
		_, err := i.Execute(stmt)
		if err != nil {
			reporting.LoxError(1, err.Error())
			return
		}
	}
}

func (i *Interpreter) Execute(stmt statement.Statement) (any, error) {
	return stmt.Accept(i)
}

func (i *Interpreter) getFloat(v any) (float64, error) {
	switch t := v.(type) {
	case int:
		return float64(t), nil
	case float64:
		return float64(t), nil
	default:
		return math.NaN(), errors.New("Cannot convert to float64, unexpected type (ConversionError)")
	}
}

func (i *Interpreter) parseTwoNumbers(left, right any) (float64, float64, error) {
	l, le := i.getFloat(left)
	r, re := i.getFloat(right)

	if le != nil {
		return l, r, le
	}
	if re != nil {
		return l, r, re
	}

	return l, r, nil
}

func (i *Interpreter) getBoolean(v any) bool {
	if v == nil || v == false {
		return false
	}
	return true
}

func (i *Interpreter) isEqual(a, b any) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil {
		return false
	}
	atype := reflect.TypeOf(a).String()
	btype := reflect.TypeOf(b).String()

	if atype != btype {
		return false
	}

	switch atype {
	case "int", "float64", "string", "bool":
		return a == b
	}

	return false
}
