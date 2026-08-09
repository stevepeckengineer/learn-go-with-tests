package shapes

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	rect := Rectangle{10.0, 10.0}
	want := 40.0

	assertFloatEquality(t, rect, rect.Perimeter(), want)
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		name string
		shape Shape
		hasArea  float64
	}{
		{name: "Rectangle", shape: Rectangle{12, 6}, hasArea: 72.0},
		{name: "Circle", shape: Circle{6}, hasArea: 113.09733552923255},
		{name: "Triangle", shape: Triangle{12, 6}, hasArea: 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			assertFloatEquality(t, tt.shape, tt.shape.Area(), tt.hasArea)
		})
	}
}

func assertFloatEquality(t testing.TB, shape Shape, got, want float64) {
	t.Helper()
	if got != want {
		t.Errorf("%#v got %g, want %g", shape, got, want)
	}
}
