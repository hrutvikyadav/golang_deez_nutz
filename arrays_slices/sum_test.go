package main

import "testing"

func TestSum(t *testing.T) {
	t.Run("Array of 5 integers", func(t *testing.T) {
		numbers := []int{1,2,3,4,5} // Changed to be a slice

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("Got: %d, want: %d; given: %v", got, want, numbers)
		}
	})

	t.Run("Slice of any number of integers", func(t *testing.T) {
		numbers := []int{1,2,3,4,5,5} // Slices can have any capacity

		got := Sum(numbers)
		want := 20

		if got != want {
			t.Errorf("Got: %d, want: %d; given: %v", got, want, numbers)
		}
	})
}
