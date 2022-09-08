package main

import "testing"

func TestPerimeter(t *testing.T) {
	got := Perimeter(Rectangle{10.0, 10.0})
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "Rectangle area", shape: Rectangle{12, 6}, want: 72.0},
		{name: "Circle area", shape: Circle{10}, want: 314.1592653589793},
		{name: "Triangle area", shape: Triangle{12, 6}, want: 36.0},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("for %q got %g want %g", tt.name, got, tt.want)
		}
	}
}
