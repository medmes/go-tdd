package arrays_slices

import "testing"

func TestSum(t *testing.T) {
	// Given
	numbers := [5]int{1, 2, 3, 4, 5}
	want := 15
	//When
	got := Sum(numbers)

	// Asserting , Then
	if got != want {
		t.Errorf("want : [ %d ] but got : [ %d ] !", want, got)

	}

}
