package main

import "testing"

func TestAdd(t *testing.T) {
	//Given, When
	expected := 12
	got := Add(6, 6)

	//Asserting
	if got != expected {
		t.Errorf("got: %d - expected: %d", got, expected)
	}

}
