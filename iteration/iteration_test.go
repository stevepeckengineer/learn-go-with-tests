package iteration

import "testing"

func TestRepeat(t *testing.T) {
	t.Run("Default to 5 repeats", func(t *testing.T) {
		repeated := Repeat("a", 0)
		expected := "aaaaa"
		assertRepeat(t, repeated, expected)
	})

	t.Run("Accept a custom number of repeats", func(t *testing.T) {
		repeated := Repeat("z", 7)
		expected := "zzzzzzz"
		assertRepeat(t, repeated, expected)
	})

}

func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat("a", 0)
	}
}

func assertRepeat(t testing.TB, repeated, expected string) {
	t.Helper()

	if repeated != expected {
		t.Errorf("got %q, expected %q", repeated, expected)
	}
}
