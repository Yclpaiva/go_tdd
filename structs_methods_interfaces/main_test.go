package main

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := rectangle.Perimeter()
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	type Shape interface {
		Area() float64
	}

	checkArea := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("%#v got %g want %g", shape, got, want)
		}
	}

	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{12.0, 6.0}
		want := 72.0

		checkArea(t, rectangle, want)
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10}
		want := 314.1592653589793

		checkArea(t, circle, want)
	})
}

func TestAreaTableDrivenTests(t *testing.T) {
	type Shape interface {
		Area() float64
	}

	checkArea := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("%#v got %g want %g", shape, got, want)
		}
	}

	// Table driven tests
	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{shape: Rectangle{12.0, 6.0}, want: 72.0},
		{shape: Circle{10}, want: 314.1592653589793},
		{shape: Circle{0}, want: 0},
		{shape: Circle{2}, want: 12.566370614359172},
		{shape: Rectangle{0, 0}, want: 0},
		{shape: Rectangle{2, 2}, want: 4},
		{shape: Triangle{12, 6}, want: 36},
		{shape: Triangle{0, 0}, want: 0},
		{shape: Triangle{2, 2}, want: 2},
		{shape: Triangle{3, 4}, want: 6},
	}
	for _, test := range areaTests {
		checkArea(t, test.shape, test.want)
	}
}
