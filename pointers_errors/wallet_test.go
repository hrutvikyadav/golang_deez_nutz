package main

import "testing"

func TestWallet(t *testing.T) {
	wallet := Wallet{}

	wallet.Deposit(Bitcoin(69))

	got := wallet.Balance()
	want := Bitcoin(69)

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
