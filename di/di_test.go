package di

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	// Given
	buffer := bytes.Buffer{}
	want := "Hello, Chris"
	//When
	Greet(&buffer, "Chris")
	//Then
	got := buffer.String()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
