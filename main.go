package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/majidmvulle/github-actions/app"
)

func Handler(anApp app.App) {
	fmt.Printf("Sum of %d and %d is: %d\n", anApp.A, anApp.B, app.AddTwoNumbers(anApp))
	fmt.Printf("Difference of %d and %d is: %d\n", anApp.A, anApp.B, app.SubtractTwoNumbers(anApp))
}

func main() {
	appA := 48
	appB := 12

	if lambdaTaskRoot := os.Getenv("LAMBDA_TASK_ROOT"); lambdaTaskRoot != "" {
		lambda.Start(Handler)
	} else {
		Handler(app.App{
			A: appA,
			B: appB,
		})
	}
}
