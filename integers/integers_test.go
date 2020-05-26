package integers

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	//Given, When
	expected := 12
	got := Add(6, 6)

	//Asserting
	if got != expected {
		t.Errorf("got: %d - expected: %d", got, expected)
	}

}

// it will not execute this function if we didn t mention // Output keyword
// because it's a funnction which is placed in the *_test package.
func ExampleAdd() {
	sum := Add(10, 30)
	fmt.Println(sum)
	//Output: 40
}
