package integers

import (
	"fmt"
	"testing"
)

func TestAdder (t  *testing.T) {
	sum := Add(1, 1)
	expected := 2

	if sum != expected {
		t.Errorf("got %d, want %d", sum, expected)
	}
}

func ExampleAdd () {
	sum := Add(2, 5)
	fmt.Println(sum)
	// Output: 7
}
