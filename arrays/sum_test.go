package sum

import "testing"

func TestSum(t *testing.T) {
	t.Run("Add an array of 5 nums", func(t *testing.T) {
		nums := [5]int{1, 2, 3, 4, 5}
		got := Sum(nums[:])
		want := 15
		assertEqual(t, got, want, nums[:])
	})

	t.Run("Add array of any size", func(t *testing.T) {
		nums := []int{1, 4, 75}
		got := Sum(nums)
		want := 80
		assertEqual(t, got, want, nums)
	})
}

func assertEqual(t testing.TB, got, want int, input []int) {
	t.Helper()
	if got != want {
		t.Errorf("got: %d, wanted: %d, input: %v", got, want, input)
	}
}
