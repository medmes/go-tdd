package struct_interface

import (
	"testing"
)

func TestPerimeter(t *testing.T) {

	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	//Asserting
	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	assertThatArea := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	}
	t.Run("Calculate rectangle perimeter", func(t *testing.T) {
		// Given, When
		rectangle := Rectangle{12.0, 6.0}
		want := 72.0

		//Asserting
		assertThatArea(t, rectangle, want)
	})

	t.Run("circles", func(t *testing.T) {
		// Given, When
		circle := Circle{10}
		want := 314.1592653589793

		//Asserting
		assertThatArea(t, circle, want)
	})

	// Introducing table driven tests in Go
	// https://github.com/golang/go/wiki/TableDrivenTests
	t.Run("testing shape area with table driven tests", func(t *testing.T) {
		// declaring a struct and instancing it inline
		areaTests := []struct {
			shape Shape   //
			want  float64 // expecting
		}{
			{Rectangle{12, 6}, 72.0},
			{Circle{10}, 314.1592653589793},
			{Triangle{12.0, 6.0}, 36.0},
		}

		for _, tt := range areaTests {
			got := tt.shape.Area()
			if got != tt.want {
				t.Errorf("got %g want %g", got, tt.want)
			}
		}
	})
}
