package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {

	//Given, When
	expected := "aaaaa"
	repeated := Repeat("a", 5)

	//Asserting
	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

// benchmark testing
// NOTE by default Benchmarks are run sequentially.
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

// Repeat Example
func ExampleRepeat() {
	repeated := Repeat("a", 10)
	fmt.Println(repeated)
	//Output: aaaaaaaaaa
}
