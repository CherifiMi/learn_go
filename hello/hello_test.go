package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("say hi to name", func(t *testing.T) {
		g := hello("mito")
		w := "hello, mito"
		assert(t, g, w)
	})
	t.Run("say hi to world", func(t *testing.T) {
		g := hello("")
		w := "hello, world"
		assert(t, g, w)
	})
}

func assert(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q wants %q", got, want)
	}
}
