package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/majidmvulle/github-actions/app"
)

func Handler(anApp app.App) error {
	if anApp.A == 0 || anApp.B == 0 {
		return fmt.Errorf("both A and B should be higher than 0")
	}

	fmt.Printf("(Sum of %d and %d is: %d\n)", anApp.A, anApp.B, app.AddTwoNumbers(anApp))
	fmt.Printf("(Difference of %d and %d is: %d\n)", anApp.A, anApp.B, app.SubtractTwoNumbers(anApp))

	return nil
}

func main() {
	appA := 48
	appB := 12

	if lambdaTaskRoot := os.Getenv("LAMBDA_TASK_ROOT"); lambdaTaskRoot != "" {
		lambda.Start(Handler)
	} else {
		err := Handler(app.App{
			A: appA,
			B: appB,
		})

		if err != nil {
			panic(err)
		}
	}
}
