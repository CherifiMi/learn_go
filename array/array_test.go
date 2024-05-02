package array

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	var nums = []int{1, 2, 3, 4, 5}

	got := Sum(nums)
	wants := 15
	assert(t, got, wants)
}
func TestSumAll(t *testing.T) {
	g := SumAll([]int{1, 2}, []int{0, 9})
	w := []int{3, 9}

	assert(t, g, w)
}
func TestSumAllTails(t *testing.T) {
	t.Run("make the sums of sume slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}

		assert(t, got, want)
	})
	t.Run("make the sums of empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{0, 9})
		want := []int{0, 9}

		assert(t, got, want)
	})
}

func assert(t testing.TB, got, want any) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %q wants %q", got, want)
	}
}
