package main

import (
	"fmt"

	"github.com/majidmvulle/github-actions/app"
)

func main() {
	a := 10
	b := 5

	fmt.Printf("Sum of %d and %d is: %d\n", a, b, app.AddTwoNumbers(a, b))
	fmt.Printf("Difference of %d and %d is: %d\n", a, b, app.SubtractTwoNumbers(a, b))
}
