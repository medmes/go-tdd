package arrays_slices

import "testing"

func TestSum(t *testing.T) {
	t.Run("Sum of collection of 5 numbers", func(t *testing.T) {
		// Given
		numbers := []int{1, 2, 3, 4, 5}
		want := 15
		//When
		got := Sum(numbers)

		// Asserting , Then
		if got != want {
			t.Errorf("want : [ %d ] but got : [ %d ] , given : %v !", want, got, numbers)
		}
	})

	t.Run("Sum of collection of any numbers", func(t *testing.T) {
		// Given
		numbers := []int{1, 2, 3}
		want := 6
		//When
		got := Sum(numbers)

		// Asserting , Then
		if got != want {
			t.Errorf("want : [ %d ] but got : [ %d ] , given : %v !", want, got, numbers)
		}
	})
}
