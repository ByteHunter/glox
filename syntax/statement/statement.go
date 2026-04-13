package syntax_statement

import (
	syntax_expression "github.com/ByteHunter/glox/syntax/expression"
)

type Statement interface {
	Accept(v Visitor) (any, error)
}

type Visitor interface {
	VisitExpressionStatement(*ExpressionStatement) (any, error)
	VisitPrintStatement(*PrintStatement) (any, error)
}

type ExpressionStatement struct {
	Statement
	Expr syntax_expression.Expression
}

func NewExpression(expr syntax_expression.Expression) *ExpressionStatement {
	return &ExpressionStatement{
		Expr: expr,
	}
}

func (e *ExpressionStatement) Accept(v Visitor) (any, error) {
	return v.VisitExpressionStatement(e)
}

type PrintStatement struct {
	Statement
	Expr syntax_expression.Expression
}

func NewPrint(expr syntax_expression.Expression) *PrintStatement {
	return &PrintStatement{
		Expr: expr,
	}
}

func (p *PrintStatement) Accept(v Visitor) (any, error) {
	return v.VisitPrintStatement(p)
}
