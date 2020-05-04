package pointers_errors

import "testing"

func TestWallet(t *testing.T) {
	//Given
	wallet := Wallet{}

	wallet.Deposit(10)

	got := wallet.Balance()
	want := 10

	//Asserting
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
	const name = iota
}
