package pointer

import (
	"testing"
)

func assertCorrectBitCoinMessage(t *testing.T, got, want Bitcoin) {
	t.Helper()
	if got != want {
		t.Errorf("expect %d, while got %d", want, got)
	}
}

func assertCorrectBitCoinStringMessage(t *testing.T, got, want Bitcoin) {
	t.Helper()
	if got != want {
		t.Errorf("expect %s, while got %s", want, got)
	}
}

func assertBalance(t *testing.T, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertError(t *testing.T, err error) {
	t.Helper()
	if err == nil {
		t.Error("wanted an error but didn't get one")
		return
	}

	t.Errorf(err.Error())
}

func TestWallet(t *testing.T) {
	t.Run("test deposit", func(t *testing.T) {
		TestWalletDeposit(t)
	})
	t.Run("test withdraw", func(t *testing.T) {
		TestWalletWithdraw(t)
	})
}

func TestWalletDeposit(t *testing.T) {
	wallet := Wallet{}
	wallet.Deposit(Bitcoin(10))

	got := wallet.Balance()
	want := Bitcoin(10)

	assertCorrectBitCoinMessage(t, got, want)
	assertCorrectBitCoinStringMessage(t, got, want)
}

func TestWalletWithdraw(t *testing.T) {
	wallet := Wallet{balance: Bitcoin(100)}
	wallet.Withdraw(9)
	got := wallet.Balance()
	want := Bitcoin(91)
	assertCorrectBitCoinMessage(t, got, want)
	assertCorrectBitCoinStringMessage(t, got, want)
}

func TestTableDrivenWallet(t *testing.T) {

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		wallet.Withdraw(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(100))
		assertBalance(t, wallet, startingBalance)
		assertError(t, err)
	})

}
