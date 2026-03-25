package expression

import (
	"fmt"

	"github.com/ByteHunter/glox/token"
)

func ExampleNewBinary() {
	fmt.Println(NewBinary(nil, token.Token{}, nil))
	// Output:
	// &{<nil> <nil> {0  <nil> 0} <nil>}
}

func ExampleGrouping() {
	fmt.Println(NewGrouping(nil))
	// Output:
	// &{<nil> <nil>}
}

func ExampleNewLiteral() {
	fmt.Println(NewLiteral(nil))
	// Output:
	// &{<nil> <nil>}
}

func ExampleNewUnary() {
	fmt.Println(NewUnary(token.Token{}, nil))
	// Output:
	// &{<nil> {0  <nil> 0} <nil>}
}
