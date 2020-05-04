package pointers_errors

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {
	//Given
	wallet := Wallet{}

	wallet.Deposit(10)

	got := wallet.Balance()
	want := 10

	fmt.Printf("address of balance in test is %v \n", &wallet.balance)
	//Asserting
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
	const name = iota
}
