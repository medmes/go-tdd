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
	spySleeper := &SpySleeper{}
	//When
	CountDown(bf, spySleeper)
	got := bf.String()
	//Then
	if got != want {
		t.Errorf("want: %q but got: %q", want, got)
	}

	if spySleeper.Calls != 4 {
		t.Errorf("not enough calls to sleeper, want 4 got %d", spySleeper.Calls)
	}
}
