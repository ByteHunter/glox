package parser

import (
	"fmt"

	"github.com/ByteHunter/glox/astprinter"
	"github.com/ByteHunter/glox/expression"
	"github.com/ByteHunter/glox/token"
)

func printExpression(expr expression.Expression) {
	fmt.Printf("%s", astprinter.NewAstPrinter().Print(expr))
}

func ExampleNewParser_empty() {
	parser := NewParser([]token.Token{})

	fmt.Println(parser)
	// Output:
	// &{[] 0}
}

func ExampleNewParser_with_tokens() {
	parser := NewParser([]token.Token{
		*token.NewToken(token.LEFT_PAREN, "(", nil, 1),
		*token.NewToken(token.NUMBER, "42", 42, 1),
		*token.NewToken(token.RIGHT_PAREN, ")", nil, 1),
	})

	fmt.Println(parser)
	// Output:
	// &{[{0 ( <nil> 1} {21 42 42 1} {1 ) <nil> 1}] 0}
}

func ExampleParser_Parse_empty_tokens_slice() {
	parser := NewParser([]token.Token{})

	fmt.Println(parser.Parse())
	// Output:
	// <nil>
}

func ExampleParser_Parse_only_eof() {
	parser := NewParser([]token.Token{
		*token.NewToken(token.EOF, "", nil, 1),
	})

	fmt.Println(parser.Parse())
	// Output:
	// <nil>
}

func ExampleParser_Primary_false() {
	p := NewParser([]token.Token{
		*token.NewToken(token.FALSE, "false", false, 1),
		*token.NewToken(token.EOF, "", nil, 1),
	})

	printExpression(p.Primary())
	// Output:
	// false
}

func ExampleParser_Primary_true() {
	p := NewParser([]token.Token{
		*token.NewToken(token.TRUE, "true", true, 1),
		*token.NewToken(token.EOF, "", nil, 1),
	})

	printExpression(p.Primary())
	// Output:
	// true
}

func ExampleParser_Primary_nil() {
	p := NewParser([]token.Token{
		*token.NewToken(token.NIL, "nil", nil, 1),
		*token.NewToken(token.EOF, "", nil, 1),
	})

	printExpression(p.Primary())
	// Output:
	// nil
}

func ExampleParser_Primary_number() {
	p := NewParser([]token.Token{
		*token.NewToken(token.NUMBER, "42", 42, 1),
		*token.NewToken(token.EOF, "", nil, 1),
	})

	printExpression(p.Primary())
	// Output:
	// 42
}

func ExampleParser_Primary_string() {
	p := NewParser([]token.Token{
		*token.NewToken(token.STRING, "test", "test", 1),
		*token.NewToken(token.EOF, "", nil, 1),
	})

	printExpression(p.Primary())
	// Output:
	// test
}

func ExampleParser_Primary_parenthesis() {
	p := NewParser([]token.Token{
		*token.NewToken(token.LEFT_PAREN, "(", nil, 1),
		*token.NewToken(token.NIL, "nil", nil, 1),
		*token.NewToken(token.RIGHT_PAREN, ")", nil, 1),
		*token.NewToken(token.EOF, "", nil, 1),
	})

	printExpression(p.Primary())
	// Output:
	// (group nil)
}

func ExampleParser_Primary_parenthesis_true() {
	p := NewParser([]token.Token{
		*token.NewToken(token.LEFT_PAREN, "(", nil, 1),
		*token.NewToken(token.TRUE, "true", true, 1),
		*token.NewToken(token.RIGHT_PAREN, ")", nil, 1),
		*token.NewToken(token.EOF, "", nil, 1),
	})

	printExpression(p.Primary())
	// Output:
	// (group true)
}

func ExampleParser_Primary_parenthesize_with_grouping() {
	p := NewParser([]token.Token{
		*token.NewToken(token.LEFT_PAREN, "(", nil, 1),
		*token.NewToken(token.NUMBER, "42", 42, 1),
		*token.NewToken(token.EQUAL_EQUAL, "==", nil, 1),
		*token.NewToken(token.NUMBER, "42", 42, 1),
		*token.NewToken(token.RIGHT_PAREN, ")", nil, 1),
		*token.NewToken(token.EOF, "", nil, 1),
	})

	printExpression(p.Primary())
	// Output:
	// (group (== 42 42))
}

func ExampleParser_Primary_parenthesize_with_error() {
	p := NewParser([]token.Token{
		*token.NewToken(token.LEFT_PAREN, "(", nil, 1),
		*token.NewToken(token.NUMBER, "42", 42, 1),
		*token.NewToken(token.EQUAL_EQUAL, "==", nil, 1),
		*token.NewToken(token.NUMBER, "42", 42, 1),
		*token.NewToken(token.EOF, "", nil, 1),
	})

	printExpression(p.Primary())
	// Output:
	// Expect ')' after expression.
	// (group (== 42 42))
}

func ExampleParser_Comparison() {
	var examples = []struct {
		tokens []token.Token
	}{
		{tokens: []token.Token{
			*token.NewToken(token.NUMBER, "42", 42, 1),
			*token.NewToken(token.GREATER, ">", nil, 1),
			*token.NewToken(token.NUMBER, "42", 42, 1),
			*token.NewToken(token.EOF, "", nil, 1),
		}},
		{tokens: []token.Token{
			*token.NewToken(token.NUMBER, "42", 42, 1),
			*token.NewToken(token.GREATER_EQUAL, ">=", nil, 1),
			*token.NewToken(token.NUMBER, "42", 42, 1),
			*token.NewToken(token.EOF, "", nil, 1),
		}},
		{tokens: []token.Token{
			*token.NewToken(token.NUMBER, "42", 42, 1),
			*token.NewToken(token.LESS, "<", nil, 1),
			*token.NewToken(token.NUMBER, "42", 42, 1),
			*token.NewToken(token.EOF, "", nil, 1),
		}},
		{tokens: []token.Token{
			*token.NewToken(token.NUMBER, "42", 42, 1),
			*token.NewToken(token.LESS_EQUAL, "<=", nil, 1),
			*token.NewToken(token.NUMBER, "42", 42, 1),
			*token.NewToken(token.EOF, "", nil, 1),
		}},
	}

	for _,e := range examples {
		p := NewParser(e.tokens)
		printExpression(p.Comparison())
		fmt.Println()
	}
	// Output:
	// (> 42 42)
	// (>= 42 42)
	// (< 42 42)
	// (<= 42 42)
}

func ExampleParser_Term_minus() {
	p := NewParser([]token.Token{
		*token.NewToken(token.NUMBER, "42", 42, 1),
		*token.NewToken(token.MINUS, "-", nil, 1),
		*token.NewToken(token.NUMBER, "42", 42, 1),
		*token.NewToken(token.EOF, "", nil, 1),
	})

	printExpression(p.Term())
	// Output:
	// (- 42 42)
}

func ExampleParser_Term_plus() {
	p := NewParser([]token.Token{
		*token.NewToken(token.NUMBER, "42", 42, 1),
		*token.NewToken(token.PLUS, "+", nil, 1),
		*token.NewToken(token.NUMBER, "42", 42, 1),
		*token.NewToken(token.EOF, "", nil, 1),
	})

	printExpression(p.Term())
	// Output:
	// (+ 42 42)
}

func ExampleParser_Factor_slash() {
	p := NewParser([]token.Token{
		*token.NewToken(token.NUMBER, "42", 42, 1),
		*token.NewToken(token.SLASH, "/", nil, 1),
		*token.NewToken(token.NUMBER, "42", 42, 1),
		*token.NewToken(token.EOF, "", nil, 1),
	})

	printExpression(p.Factor())
	// Output:
	// (/ 42 42)
}

func ExampleParser_Factor_star() {
	p := NewParser([]token.Token{
		*token.NewToken(token.NUMBER, "21", 21, 1),
		*token.NewToken(token.STAR, "*", nil, 1),
		*token.NewToken(token.NUMBER, "2", 2, 1),
		*token.NewToken(token.EOF, "", nil, 1),
	})

	printExpression(p.Factor())
	// Output:
	// (* 21 2)
}

func ExampleParser_Unary_bang() {
	p := NewParser([]token.Token{
		*token.NewToken(token.BANG, "!", nil, 1),
		*token.NewToken(token.TRUE, "true", true, 1),
		*token.NewToken(token.EOF, "", nil, 1),
	})

	printExpression(p.Unary())
	// Output:
	// (! true)
}

func ExampleParser_Unary_minus() {
	p := NewParser([]token.Token{
		*token.NewToken(token.MINUS, "-", nil, 1),
		*token.NewToken(token.NUMBER, "42", 42, 1),
		*token.NewToken(token.EOF, "", nil, 1),
	})

	printExpression(p.Unary())
	// Output:
	// (- 42)
}
