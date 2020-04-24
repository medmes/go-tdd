package hello

import "testing"

/**
* Introducing t.Run tool, running several subTest cases for a given function:
 */
func TestHello(t *testing.T) {

	// Refactoring test assertion by introducing
	//a common function and helper to help us to catch which number the test has been failed
	// In Go, we can assign a function to a variable and call it as normal function!
	// Inner-function:

	assertThat := func(t *testing.T, got string, expected string) {
		t.Helper() // it provide us useful information such like line number where the test has been failed.
		if got != expected {
			t.Errorf("got %q excpected %q", got, expected)
		}
	}

	t.Run("saying hello to people!", func(t *testing.T) {
		// Given, When
		expected := "Hello Mo!"
		got := Hello("Mo!")

		//Asserting
		assertThat(t, got, expected)

	})

	t.Run("saying hello world, when no argument iss passed!", func(t *testing.T) {
		// Given, When
		expected := "Hello World!"
		got := Hello("")

		//Asserting
		assertThat(t, got, expected)
	})

}
