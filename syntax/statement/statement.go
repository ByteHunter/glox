package statement

import (
	"github.com/ByteHunter/glox/syntax/expression"
	"github.com/ByteHunter/glox/token"
)

type Statement interface {
	Accept(v Visitor) (any, error)
}

type Visitor interface {
	VisitExpressionStatement(*ExpressionStatement) (any, error)
	VisitPrintStatement(*PrintStatement) (any, error)
	VisitVariableStatement(*VariableStatement) (any, error)
}

type ExpressionStatement struct {
	Statement
	Expr expression.Expression
}

func NewExpression(expr expression.Expression) *ExpressionStatement {
	return &ExpressionStatement{
		Expr: expr,
	}
}

func (e *ExpressionStatement) Accept(v Visitor) (any, error) {
	return v.VisitExpressionStatement(e)
}

type PrintStatement struct {
	Statement
	Expr expression.Expression
}

func NewPrint(expr expression.Expression) *PrintStatement {
	return &PrintStatement{
		Expr: expr,
	}
}

func (p *PrintStatement) Accept(v Visitor) (any, error) {
	return v.VisitPrintStatement(p)
}

type VariableStatement struct {
	Statement
	Name        token.Token
	Initializer expression.Expression
}

func NewVariable(name token.Token, initializer expression.Expression) *VariableStatement {
	return &VariableStatement{
		Name:        name,
		Initializer: initializer,
	}
}

func (variable *VariableStatement) Accept(v Visitor) (any, error) {
	return v.VisitVariableStatement(variable)
}
