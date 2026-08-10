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
		err := wallet.Withdraw(Bitcoin(5))

		assertNoError(t, err)
		assertBalance(t, wallet, Bitcoin(5))
	})

	t.Run("Insufficient funds, Withdraw errors with no balance change", func(t *testing.T) {
		wallet := Wallet{Bitcoin(3)}
		err := wallet.Withdraw(Bitcoin(5))

		assertError(t, err, ErrInsufficientFunds)
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

// helper to assert absence of error
func assertNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("got an error but didn't want one")
	}
}

// helper to assert presence of error
func assertError(t testing.TB, err, want error) {
	t.Helper()
	if err == nil {
		t.Fatal("Expected error, got none")
	}

	if err != want {
		t.Errorf("got %q, want %q", err, want)
	}
}
