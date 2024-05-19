package wallet

import (
	"testing"
)

func TestWallet(t *testing.T) {
	assertBalance := func(t testing.TB, wallet Wallet, want string) {
		t.Helper()
		got := wallet.Balance()

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	assertError := func(t testing.TB, got, want error) {
		t.Helper()
		if got == nil {
			t.Fatal("didn't get an error but want one")
		}
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}
	assertNoError := func(t testing.TB, got error) {
		t.Helper()
		if got != nil {
			t.Fatal("got an error but didn't want one")
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		// print memory value: fmt.Println("address of balance in test is", &wallet.balance)
		assertBalance(t, wallet, "10 BTC")
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		err := wallet.Withdraw(Bitcoin(5))
		// print memory value: fmt.Println("address of balance in test is", &wallet.balance)
		assertNoError(t, err)
		assertBalance(t, wallet, "5 BTC")
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(100))

		assertBalance(t, wallet, startingBalance.String())
		assertError(t, err, ErrInsufficientFunds)
	})
}
