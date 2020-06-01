package mocking

import (
	"bytes"
	"testing"
)

func TestCountDown(t *testing.T) {
	//Given
	bf := &bytes.Buffer{}
	want := `3
2
1
Go!`
	//When
	CountDown(bf)
	got := bf.String()
	//Then
	if got != want {
		t.Errorf("want: %q but got: %q", want, got)
	}
}
