package main

import "testing"

func TestSearch(t *testing.T) {
	dict := map[string]string{"test": "lulz"}

	definition := Search(dict, "test")
	want := "lulz"
	assertStrings(t, definition, want)
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
