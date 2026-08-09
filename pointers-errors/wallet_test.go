package wallet

import (
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("Deposit increases balance", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))

		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw decreases balance", func(t *testing.T) {
		wallet := Wallet{Bitcoin(10)}
		wallet.Withdraw(Bitcoin(5))

		assertBalance(t, wallet, Bitcoin(5))
	})

	t.Run("Insufficient funds, Withdraw errors with no balance change", func(t *testing.T) {
		wallet := Wallet{Bitcoin(3)}
		err := wallet.Withdraw(Bitcoin(5))
		want := ErrInsufficientFunds

		assertErr(t, err, want)
		assertBalance(t, wallet, Bitcoin(3))
	})
}

// helper to assert balance equality
func assertBalance(t testing.TB, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()

	if got != want {
		t.Errorf("%#v got %s want %s", wallet, got, want)
	}
}

// helper to assert presence of error
func assertErr(t testing.TB, err, want error) {
	if err == nil {
		t.Fatal("Expected error, got none")
	}

	if err != want {
		t.Errorf("got %q, want %q", err, want)
	}
}
