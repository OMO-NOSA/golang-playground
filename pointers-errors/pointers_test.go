package pointerserrors

import "testing"

func TestWallet(t *testing.T) {
	wallet := Wallet{}
	wallet.Deposit(10)

	got := wallet.Balance()
	want := 10.0

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
