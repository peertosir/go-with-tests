package oop

import (
	"math"
	"testing"
)

func checkFloat64Equal(t testing.TB, got, want float64) {
	t.Helper()
	if math.Abs(got-want) >= 0.01 {
		t.Errorf("got %.2f, but want %.2f", got, want)
	}
}

func TestPerimeter(t *testing.T) {
	perimeterTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{10.0, 10.0}, 40.0},
		{Circle{5.0}, 31.41},
		{Triangle{5.0, 5.0, 5.0}, 15.0},
	}

	for _, tt := range perimeterTests {
		got := tt.shape.Perimeter()
		checkFloat64Equal(t, got, tt.want)
	}
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{10.0, 10.0}, 100.0},
		{Circle{10.0}, 314.15},
		{Triangle{5.0, 4.0, 3.0}, 6.0},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		checkFloat64Equal(t, got, tt.want)
	}
}
