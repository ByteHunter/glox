package astprinter

import (
	"fmt"
	"strings"

	syntax_expression "github.com/ByteHunter/glox/syntax/expression"
)

type AstPrinter struct {
}

func NewAstPrinter() *AstPrinter {
	return &AstPrinter{}
}

func (a *AstPrinter) VisitBinaryExpression(expr *syntax_expression.Binary) (any, error) {
	return a.Parentesize(expr.Operator.Lexeme, expr.Left, expr.Right), nil
}

func (a *AstPrinter) VisitGroupingExpression(expr *syntax_expression.Grouping) (any, error) {
	return a.Parentesize("group", expr.Expr), nil
}

func (a *AstPrinter) VisitLiteralExpression(expr *syntax_expression.Literal) (any, error) {
	if expr.Value == nil {
		return "nil", nil
	}

	return fmt.Sprintf("%v", expr.Value), nil
}

func (a *AstPrinter) VisitUnaryExpression(expr *syntax_expression.Unary) (any, error) {
	return a.Parentesize(expr.Operator.Lexeme, expr.Right), nil
}

func (a *AstPrinter) Print(expr syntax_expression.Expression) (any, error) {
	if expr == nil {
		return "nil", nil
	}

	return expr.Accept(a)
}

func (a *AstPrinter) Parentesize(name string, exprs ...syntax_expression.Expression) string {
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
