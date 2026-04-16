package expression

import (
	"github.com/ByteHunter/glox/token"
)

type Expression interface {
	Accept(v Visitor) (any, error)
}

type Visitor interface {
	VisitAssignExpression(*Assign) (any, error)
	VisitBinaryExpression(*Binary) (any, error)
	VisitGroupingExpression(*Grouping) (any, error)
	VisitLiteralExpression(*Literal) (any, error)
	VisitUnaryExpression(*Unary) (any, error)
	VisitVariableExpression(*Variable) (any, error)
}

type Assign struct {
	Expression
	Name token.Token
	Expr Expression
}

func NewAssign(name token.Token, expr Expression) *Assign {
	return &Assign{
		Name: name,
		Expr: expr,
	}
}

func (assign *Assign) Accept(v Visitor) (any, error) {
	return v.VisitAssignExpression(assign)
}

type Binary struct {
	Expression
	Left     Expression
	Operator token.Token
	Right    Expression
}

func NewBinary(left Expression, operator token.Token, right Expression) *Binary {
	return &Binary{
		Left:     left,
		Operator: operator,
		Right:    right,
	}
}

func (binary *Binary) Accept(v Visitor) (any, error) {
	return v.VisitBinaryExpression(binary)
}

type Grouping struct {
	Expression
	Expr Expression
}

func NewGrouping(expr Expression) *Grouping {
	return &Grouping{
		Expr: expr,
	}
}

func (grouping *Grouping) Accept(v Visitor) (any, error) {
	return v.VisitGroupingExpression(grouping)
}

type Literal struct {
	Expression
	Value any
}

func NewLiteral(value any) *Literal {
	return &Literal{
		Value: value,
	}
}

func (literal *Literal) Accept(v Visitor) (any, error) {
	return v.VisitLiteralExpression(literal)
}

type Unary struct {
	Expression
	Operator token.Token
	Right    Expression
}

func NewUnary(operator token.Token, right Expression) *Unary {
	return &Unary{
		Operator: operator,
		Right:    right,
	}
}

func (unary *Unary) Accept(v Visitor) (any, error) {
	return v.VisitUnaryExpression(unary)
}

type Variable struct {
	Expression
	Name token.Token
}

func NewVariable(name token.Token) *Variable {
	return &Variable{
		Name: name,
	}
}

func (variable *Variable) Accept(v Visitor) (any, error) {
	return v.VisitVariableExpression(variable)
}
