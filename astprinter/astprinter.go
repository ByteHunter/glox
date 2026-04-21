package astprinter

import (
	"fmt"
	"strings"

	"github.com/ByteHunter/glox/syntax/expression"
	"github.com/ByteHunter/glox/syntax/statement"
)

type AstPrinter struct {
}

func NewAstPrinter() *AstPrinter {
	return &AstPrinter{}
}

func (a *AstPrinter) VisitAssignExpression(expr *expression.Assign) (any, error) {
	return a.Parentesize(expr.Name.Lexeme, expr.Value), nil
}

func (a *AstPrinter) VisitBinaryExpression(expr *expression.Binary) (any, error) {
	return a.Parentesize(expr.Operator.Lexeme, expr.Left, expr.Right), nil
}

func (a *AstPrinter) VisitGroupingExpression(expr *expression.Grouping) (any, error) {
	return a.Parentesize("group", expr.Expr), nil
}

func (a *AstPrinter) VisitLiteralExpression(expr *expression.Literal) (any, error) {
	if expr.Value == nil {
		return "nil", nil
	}

	return fmt.Sprintf("%v", expr.Value), nil
}

func (a *AstPrinter) VisitVariableExpression(*expression.Variable) (any, error) {
	return nil, nil
}

func (a *AstPrinter) VisitExpressionStatement(*statement.ExpressionStatement) (any, error) {
	return nil, nil
}

func (a *AstPrinter) VisitPrintStatement(*statement.PrintStatement) (any, error) {
	return nil, nil
}

func (a *AstPrinter) VisitVariableStatement(*statement.VariableStatement) (any, error) {
	return nil, nil
}

func (a *AstPrinter) VisitUnaryExpression(expr *expression.Unary) (any, error) {
	return a.Parentesize(expr.Operator.Lexeme, expr.Right), nil
}

func (a *AstPrinter) Print(expr expression.Expression) (any, error) {
	if expr == nil {
		return "nil", nil
	}

	return expr.Accept(a)
}

func (a *AstPrinter) Parentesize(name string, exprs ...expression.Expression) string {
	var buffer strings.Builder
	buffer.Reset()
	buffer.WriteString("(" + name)

	for _, expr := range exprs {
		if expr == nil {
			buffer.WriteString(" nil")
			continue
		}
		r, _ := expr.Accept(a)
		fmt.Fprintf(&buffer, " %v", r)
	}
	buffer.WriteString(")")

	return buffer.String()
}
