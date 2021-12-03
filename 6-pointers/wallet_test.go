package wallet

import "testing"

func TestWallet(t *testing.T) {
	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Dollar(10))

		assertBalance(t, wallet, Dollar(10))
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Dollar(20)}
		err := wallet.Withdraw(Dollar(5))

		assertNoError(t, err)
		assertBalance(t, wallet, Dollar(15))
	})

	t.Run("Withdrawal insufficient funds", func(t *testing.T) {
		wallet := Wallet{balance: Dollar(20)}
		err := wallet.Withdraw(Dollar(100))

		assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, wallet, Dollar(20))
	})
}

func assertBalance(t testing.TB, wallet Wallet, want Dollar) {
	t.Helper()

	got := wallet.Balance()

	if got != want {
		t.Errorf("expected %s but got %s", want, got)
	}
}

func assertError(t testing.TB, err error, want error) {
	t.Helper()

	if err == nil {
		t.Errorf("wanted an error but didn't get one")
	}

	if err.Error() != want.Error() {
		t.Errorf("expected %q but got %q", err, want)
	}
}

func assertNoError(t testing.TB, err error) {
	t.Helper()

	if err != nil {
		t.Fatal("got an error but didn't want one")
	}
}
