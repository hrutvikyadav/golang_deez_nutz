package main

import "testing"

func TestWallet(t *testing.T) {
	wallet := Wallet{}

	wallet.Deposit(69)

	got := wallet.Balance()
	want := 69

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
