package expression

import (
	"github.com/ByteHunter/glox/token"
)

type Expression interface {
	accept(v Visitor) any
}

type Visitor interface {
	visitBinaryExpression(*Binary) any
	visitGroupingExpression(*Grouping) any
	visitLiteralExpression(*Literal) any
	visitUnaryExpression(*Unary) any
}

type Binary struct {
	Expression
	left     Expression
	operator token.Token
	right    Expression
}

func NewBinary(left Expression, operator token.Token, right Expression) *Binary {
	return &Binary{
		left:     left,
		operator: operator,
		right:    right,
	}
}

func (binary *Binary) accept(v Visitor) any {
	return v.visitBinaryExpression(binary)
}

type Grouping struct {
	Expression
	expression Expression
}

func NewGrouping(expression Expression) *Grouping {
	return &Grouping{
		expression: expression,
	}
}

func (grouping *Grouping) accept(v Visitor) any {
	return v.visitGroupingExpression(grouping)
}

type Literal struct {
	Expression
	value any
}

func NewLiteral(value any) *Literal {
	return &Literal{
		value: value,
	}
}

func (literal *Literal) accept(v Visitor) any {
	return v.visitLiteralExpression(literal)
}

type Unary struct {
	Expression
	operator token.Token
	right    Expression
}

func NewUnary(operator token.Token, right Expression) *Unary {
	return &Unary{
		operator: operator,
		right:    right,
	}
}

func (unary *Unary) accept(v Visitor) any {
	return v.visitUnaryExpression(unary)
}
