package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"os"

	"github.com/majidmvulle/github-actions/app"
)

func Handler(a int, b int) {
	fmt.Printf("Sum of %d and %d is: %d\n", a, b, app.AddTwoNumbers(a, b))
	fmt.Printf("Difference of %d and %d is: %d\n", a, b, app.SubtractTwoNumbers(a, b))
}

func main() {

	if lambdaTaskRoot := os.Getenv("LAMBDA_TASK_ROOT"); lambdaTaskRoot != "" {
		lambda.Start(Handler)
	} else {
		Handler(48, 21)
	}
}
