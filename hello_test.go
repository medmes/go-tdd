package hello

import "testing"

func TestHello(t *testing.T) {
	expected := "Hello Mo!"
	got := Hello("Mo!")

	//Asserting
	if got != expected {
		t.Errorf("got %q excpected %q", got, expected)
	}
}
