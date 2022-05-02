package app_test

import (
	"testing"

	"github.com/majidmvulle/github-actions/app"
)

func TestAddTwoNumbers(t *testing.T) {
	t.Parallel()

	thing := app.App{
		A: 20,
		B: 12,
	}

	if sum := app.AddTwoNumbers(thing); sum != 32 {
		t.Fatalf(`addTwoNumbers(%v) = %d, want  %d`, thing, sum, 32)
	}
}

func TestSubtractTwoNumbers(t *testing.T) {
	t.Parallel()

	thing := app.App{
		A: 20,
		B: 12,
	}

	if difference := app.SubtractTwoNumbers(thing); difference != 8 {
		t.Fatalf(`addTwoNumbers(%v) = %d, want  %d`, thing, difference, 8)
	}
}
