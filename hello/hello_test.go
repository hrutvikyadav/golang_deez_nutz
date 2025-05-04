package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Hrutvik")
	want := "Hello, Hrutvik"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
