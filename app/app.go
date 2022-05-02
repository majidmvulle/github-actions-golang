package app

type App struct {
	A int
	B int
}

func AddTwoNumbers(app App) int {
	return app.A + app.B
}

func SubtractTwoNumbers(app App) int {
	return app.A - app.B
}
