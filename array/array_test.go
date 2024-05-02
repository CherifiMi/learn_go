package array

import "testing"

func TestSum(t *testing.T) {
	var nums = [5]int{1, 2, 3, 4, 5}

	got := Sum(nums)
	wants := 15
	assert(t, got, wants, nums)
}

func assert(t testing.TB, got, want, nums any) {
	t.Helper()
	if got != want {
		t.Errorf("got %d wants %d, %v", got, want, nums)
	}
}
