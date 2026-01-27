package main

import "testing"

func TestWallet(t *testing.T) {
	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(69))

		assertBalance(t, wallet, Bitcoin(69))
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{Bitcoin(69)}
		wallet.Withdraw(Bitcoin(9))

		assertBalance(t, wallet,Bitcoin(60))
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		wallet := Wallet{Bitcoin(69)}
		err := wallet.Withdraw(Bitcoin(420))

		assertError(t, err, "insufficient balance")
		assertBalance(t, wallet,Bitcoin(69))
	})
}

func assertBalance (t testing.TB, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertError (t testing.TB, gotErr error, want string) {
	t.Helper()
	if gotErr == nil {
		t.Fatal("expected error but did not get one")
	}

	if gotErr.Error() != want {
		t.Errorf("got %q want %q", gotErr, want)
	}
}
