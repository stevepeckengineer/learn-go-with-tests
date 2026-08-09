package sum

import (
	"slices"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("Add array of any size", func(t *testing.T) {
		nums := []int{1, 4, 75}
		got := Sum(nums)
		want := 80
		assertEqual(t, got, want, nums)
	})
}

func TestSumAll(t *testing.T) {
	t.Run("Add a single slice", func(t *testing.T) {
		nums := []int{1, 4, 75}
		got := Sum(nums)
		want := 80
		assertEqual(t, got, want, nums)
	})

	t.Run("Add multiple slices", func(t *testing.T) {
		got := SumAll([]int{3, 4}, []int{5, 6})
		want := []int{7, 11}
		if !slices.Equal(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}

func assertEqual(t testing.TB, got, want int, input []int) {
	t.Helper()
	if got != want {
		t.Errorf("got: %d, wanted: %d, input: %v", got, want, input)
	}
}
