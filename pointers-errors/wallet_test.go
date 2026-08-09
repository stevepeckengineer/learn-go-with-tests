package wallet

import (
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("Deposit increases balance", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))

		got := wallet.Balance()
		want := Bitcoin(10)

		if got != want {
			t.Errorf("%#v got %s want %s", wallet, got, want)
		}
	})

	t.Run("Withdraw decreases balance", func(t *testing.T) {
		wallet := Wallet{10}
		wallet.Withdraw(5)
		got := wallet.Balance()
		want := Bitcoin(5)
		if got != want {
			t.Errorf("%#v got %s want %s", wallet, got, want)
		}
	})

	t.Run("Insufficient funds, Withdraw does nothing", func(t *testing.T) {
		wallet := Wallet{3}
		wallet.Withdraw(5)
		got := wallet.Balance()
		want := Bitcoin(4)
		if got != want {
			t.Errorf("%#v got %s want %s", wallet, got, want)
		}
	})
}
