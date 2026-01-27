package main

import "testing"

func TestSum(t *testing.T) {
	numbers := [5]int{1,2,3,4,5} // Arrays have fixed capacity
	// numbers := [...]int{1,2,3,4,5} NOTE: capacity can also be infered

	got := Sum(numbers)
	want := 15

	if got != want {
		t.Errorf("Got: %d, want: %d; given: %v", got, want, numbers)
	}
}
