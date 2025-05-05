package main

import "testing"


func assertCorrectMessage(got string, want string, t *testing.T) {
	t.Helper() // doing this will report correct stack trace on test failure
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestHello(t *testing.T) {
	t.Run("say hello to person", func(t *testing.T) {
		got := Hello("Hrutvik")
		want := "Hello, Hrutvik"
		assertCorrectMessage(got, want, t)
	})

	t.Run("default greeting", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"
		assertCorrectMessage(got, want, t)
	})
}
