package integers

import "testing"

func TestAdder(t *testing.T) {
	t.Run("add two numbers", func(t *testing.T) {
		sum := Add(2, 2)
		result := 4

		assert(t, sum, result)
	})
	t.Run("add two numbers", func(t *testing.T) {
		sum := Add(4, 2)
		result := 6

		assert(t, sum, result)
	})
}

func assert(t testing.TB, got, want any) {
	t.Helper()
	if got != want {
		t.Errorf("got %d wants %d", got, want)
	}
}
