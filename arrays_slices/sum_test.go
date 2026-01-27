package main

import (
	"reflect"
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

func TestSumAllTails(t *testing.T) {
	assertSums := func (t testing.TB, got []int, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Got: %d, want: %d", got, want)
		}
	}

	t.Run("Sum tails of multiple slices", func(t *testing.T) {
		s1 := []int{1,2,3}
		s2 := []int{2,3,4}

		got := SumAllTails(s1, s2)
		want := []int{5, 7}

		assertSums(t, got, want)
	})

	t.Run("Safely sum tails of EMPTY slices", func(t *testing.T) {
		s1 := []int{}
		s2 := []int{2,3,4}

		got := SumAllTails(s1, s2)
		want := []int{0, 7}

		assertSums(t, got, want)
	})
}
