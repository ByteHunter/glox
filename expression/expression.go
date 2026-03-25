package expression

import (
	"github.com/ByteHunter/glox/token"
)

type Expression interface {
	accept(v Visitor) any
}

type Visitor interface {
	VisitBinaryExpression(*Binary) any
	VisitGroupingExpression(*Grouping) any
	VisitLiteralExpression(*Literal) any
	VisitUnaryExpression(*Unary) any
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

func (binary *Binary) Accept(v Visitor) any {
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

func (grouping *Grouping) Accept(v Visitor) any {
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

func (literal *Literal) Accept(v Visitor) any {
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

func (unary *Unary) Accept(v Visitor) any {
	return v.VisitUnaryExpression(unary)
}
