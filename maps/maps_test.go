package maps

import (
	"testing"
)

func TestSearch(t *testing.T) {
	dic := map[string]string{"test": "111"}

	got := Search(dic, "test")
	want := "111"

	assert(t, got, want)
}

func assert(t testing.TB, want, got string) {
	t.Helper()
	if want != got {
		t.Errorf("got %q, want %q", got, want)
	}
}
