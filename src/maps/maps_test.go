package maps

import "testing"

func TestSearch(t *testing.T) {
	//Given
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		// Given, When
		want := "this is just a test"
		got, _ := dictionary.Search("test")

		//Asserting
		assertStrings(t, want, got)
	})

	t.Run("unKnown word, asserting not found error", func(t *testing.T) {
		// Given, When
		_, err := dictionary.Search("xxxx")

		//Asserting
		assertError(t, err, notFoundError)
	})
}

// assert Helper
func assertStrings(t *testing.T, w string /* want */, g string /* got*/) {
	t.Helper()

	if g != w {
		t.Errorf("got %q want %q", g, w)
	}
}

func assertError(t *testing.T, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}
