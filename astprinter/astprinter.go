package astprinter

import (
	"fmt"
	"strings"

	"github.com/ByteHunter/glox/expression"
)

type AstPrinter struct {
}

func NewAstPrinter() *AstPrinter {
	return &AstPrinter{}
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
