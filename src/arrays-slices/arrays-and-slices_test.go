package arrays_slices

import (
	"reflect"
	"testing"
)

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
}

func TestSumAll(t *testing.T) {
	// Given , When
	want := []int{3, 9}
	got := SumAll([]int{1, 2}, []int{0, 9})

	// Asserting, Then
	if !reflect.DeepEqual(got, want) {
		t.Errorf("want :  %d but got :  %d !", want, got)

	}
}
