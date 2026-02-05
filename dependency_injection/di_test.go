package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buf_for_di := bytes.Buffer{}
	Greet(&buf_for_di, "World")

	got := buf_for_di.String()
	want := "Hello, World"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
