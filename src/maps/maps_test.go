package maps

import "testing"

func TestSearch(t *testing.T) {

	//Given
	dictionary := map[string]string{"test": "this is just a test"}

	want := "this is just a test"
	got := Search(dictionary, "test")

	//Asserting
	if got != want {
		t.Errorf("got %q want %q given, %q", got, want, "test")
	}

}
