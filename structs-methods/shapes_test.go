package shapes

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	rect := Rectangle{10.0, 10.0}
	got := Perimeter(rect)
	want := 40.0

	assertFloatEquality(t, got, want)
}

func TestArea(t *testing.T) {
	t.Run("Area for Rectangle", func(t *testing.T) {
		rect := Rectangle{10.0, 10.0}
		got := rect.Area()
		want := 100.0

		assertFloatEquality(t, got, want)
	})

	t.Run("Area for Circle", func(t *testing.T) {
		cir := Circle{6.0}
		got := cir.Area()
		want := 113.09733552923255

		assertFloatEquality(t, got, want)
	})

}

func assertFloatEquality(t testing.TB, got, want float64) {
	t.Helper()
	if got != want {
		t.Errorf("got %g want %g", got, want)
	}
}
