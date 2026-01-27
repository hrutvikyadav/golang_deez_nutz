package main

import (
	"slices"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("Slice of any number of integers", func(t *testing.T) {
		numbers := []int{1,2,3,4,5,5} // Slices can have any capacity

		got := Sum(numbers)
		want := 20

		if got != want {
			t.Errorf("Got: %d, want: %d; given: %v", got, want, numbers)
		}
	})
}

// We need a new function called SumAll which will take a varying number of slices, returning a new slice containing the totals for each slice passed in.
func TestSumAll(t *testing.T) {
	s1 := []int{1,2,3}
	s2 := []int{2,3,4}

	got := SumAll(s1, s2)
	want := []int{6, 9}

	// if got != want { WARN: slices annot be compared with `=` operator, use slices standard lib package
	if !slices.Equal(got, want) {
		t.Errorf("Got: %d, want: %d", got, want)
	}
}
