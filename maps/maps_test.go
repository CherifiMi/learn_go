package maps

import (
	"testing"
)

func TestSearch(t *testing.T) {
	dic := Dic{"test": "111"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dic.Search("test")
		want := "111"

		assert(t, got, want)
	})
	t.Run("unknown word", func(t *testing.T) {
		_, err := dic.Search("mito")
		want := "could not find the word"

		assert(t, err.Error(), want)
	})
}
func TestAdd(t *testing.T) {
	dic := Dic{}
	dic.Add("test", "111")

	got, err := dic.Search("test")
	want := "111"

	if err != nil {
		t.Fatal("should found the word", err)
	}

	assert(t, got, want)

}

func assert(t testing.TB, want, got string) {
	t.Helper()
	if want != got {
		t.Errorf("got %q, want %q", got, want)
	}
}
