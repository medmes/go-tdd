package arrays_slices

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	// Given
	numbers := []int{1, 2, 3, 4, 5}
	want := 15
	//When
	got := Sum(numbers)

	// Asserting , Then
	if got != want {
		t.Errorf("want : [ %d ] but got : [ %d ] , given : %v !", want, got, numbers)
	}
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

//Our tests have some repeated code around assertion again, let's extract that into a function
// So we need to optimize our test code:
func TestSumAllTails(t *testing.T) {

	checkSums := func(t *testing.T, got, want []int) {
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}
	t.Run("make the sums of some slices", func(t *testing.T) {
		want := []int{2, 9}
		got := SumAllTails([]int{1, 2}, []int{0, 9})

		// Asserting
		checkSums(t, got, want)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		want := []int{0, 9}
		got := SumAllTails([]int{}, []int{3, 4, 5})

		// Asserting
		checkSums(t, got, want)
	})
}
