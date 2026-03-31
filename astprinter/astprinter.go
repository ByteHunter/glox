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

func (a *AstPrinter) VisitBinaryExpression(expr *expression.Binary) any {
	return a.Parentesize(expr.Operator.Lexeme, expr.Left, expr.Right)
}

func (a *AstPrinter) VisitGroupingExpression(expr *expression.Grouping) any {
	return a.Parentesize("group", expr.Expr)
}

func (a *AstPrinter) VisitLiteralExpression(expr *expression.Literal) any {
	if expr.Value == nil {
		return "nil"
	}

	return fmt.Sprintf("%v", expr.Value)
}

func (a *AstPrinter) VisitUnaryExpression(expr *expression.Unary) any {
	return a.Parentesize(expr.Operator.Lexeme, expr.Right)
}

func (a *AstPrinter) Print(expr expression.Expression) any {
	if expr == nil {
		return "nil"
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
		fmt.Fprintf(&buffer, " %v", expr.Accept(a))
	}
	buffer.WriteString(")")

	return buffer.String()
}
