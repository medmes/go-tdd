package maps

import (
	"testing"
)

func TestDictionary_Search(t *testing.T) {
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
		assertError(t, err, NotFoundError)
	})
}

func TestDictionary_Add(t *testing.T) {

	dictionary := Dictionary{}
	dictionary.Add("test", "this is just a test")
	assertStrings(t, dictionary["test"], "this is just a test")
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
