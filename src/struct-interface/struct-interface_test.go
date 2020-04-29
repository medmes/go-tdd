package struct_interface

import "testing"

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
	assertThat := func(t *testing.T, got, want float64) {
		t.Helper()
		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	}
	t.Run("Calculate rectangle perimeter", func(t *testing.T) {
		// Given, When
		rectangle := Rectangle{12.0, 6.0}
		got := rectangle.Area()
		want := 72.0

		//Asserting
		assertThat(t, got, want)
	})

	t.Run("circles", func(t *testing.T) {
		// Given, When
		circle := Circle{10}
		got := circle.Area()
		want := 314.1592653589793

		//Asserting
		assertThat(t, got, want)
	})
}
