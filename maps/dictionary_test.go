package main

import "testing"

func TestSearch(t *testing.T) {
	dict := Dictionary{"test": "lulz"}

	t.Run("search known word", func(t *testing.T) {
		definition, _ := dict.Search("test")
		want := "lulz"
		assertStrings(t, definition, want)
	})

	t.Run("search unknown word", func(t *testing.T) {
		_, err := dict.Search("testes")
		if err == nil {
			t.Fatal("expected error")
		}
		assertError(t, err, ErrWordNotFound)
	})
}

func TestAdd(t *testing.T) {
	dict := Dictionary{"test": "lulz"}
	dict.Add("testes", "ballz")

	definition, err := dict.Search("testes")
	want := "ballz"

	if err != nil {
		t.Fatal("did not expect error:", err)
	}

	assertStrings(t, definition, want)
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
