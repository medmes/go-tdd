package mocking

import (
	"bytes"
	"reflect"
	"testing"
)

func TestCountDown(t *testing.T) {
	t.Run("prints 3 to Go!", func(t *testing.T) {
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
	})

	t.Run("sleep before every print", func(t *testing.T) {
		//Given
		want := []string{
			sleep,
			sleep,
			sleep,
			write,
			write,
			write,
			sleep,
			write,
		}
		spySleeperPrint := &CountdownOperationsSpy{}
		//When
		CountDown(spySleeperPrint, spySleeperPrint) // that implement io.Writer, Sleeper interfaces
		//Then
		if !reflect.DeepEqual(want, spySleeperPrint.Calls) {
			t.Errorf("wanted calls: %v got: %v", want, spySleeperPrint.Calls)
		}
	})
}
