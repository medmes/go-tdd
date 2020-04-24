package iteration

import "testing"

func TestRepeat(t *testing.T) {

	//Given, When
	expected := "aaaaa"
	repeated := Repeat("a")

	//Asserting
	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}
