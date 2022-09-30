package pointer

import "testing"

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

	got := wallet.Ballance()
	want := Bitcoin(10)

	assertCorrectBitCoinMessage(t, got, want)
	assertCorrectBitCoinStringMessage(t, got, want)
}

func TestWalletWithdraw(t *testing.T) {
	wallet := Wallet{balance: Bitcoin(100)}
	wallet.Withdraw(9)
	got := wallet.Ballance()
	want := Bitcoin(91)
	assertCorrectBitCoinMessage(t, got, want)
	assertCorrectBitCoinStringMessage(t, got, want)
}
