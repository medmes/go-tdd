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

// benchmark testing
// NOTE by default Benchmarks are run sequentially.
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}
