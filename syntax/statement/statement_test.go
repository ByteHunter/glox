package statement

import "fmt"

func ExampleNewExpression() {
	fmt.Println(NewExpression(nil))
	// Output:
	// &{<nil> <nil>}
}

func ExampleNewPrint() {
	fmt.Println(NewPrint(nil))
	// Output:
	// &{<nil> <nil>}
}
