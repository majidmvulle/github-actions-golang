package app_test

import (
	"testing"

	"github.com/majidmvulle/github-actions/app"
)

func TestAddTwoNumbers(t *testing.T) {
	t.Parallel()

	if sum := app.AddTwoNumbers(20, 12); sum != 32 {
		t.Fatalf(`addTwoNumbers(%d, %d) = %d, want  %d`, 20, 12, sum, 32)
	}
}
