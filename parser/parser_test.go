package parser

import (
	"fmt"

	"github.com/ByteHunter/glox/token"
)

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

func ExampleParser_Parse_empty() {
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

func ExampleParser_Parse_simple() {
	parser := NewParser([]token.Token{
		*token.NewToken(token.TRUE, "true", true, 1),
		*token.NewToken(token.EOF, "", nil, 1),
	})

	fmt.Println(parser.Parse())
	// Output:
	// &{<nil> true}
}
