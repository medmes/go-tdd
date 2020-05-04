package pointers_errors

import "fmt"

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	fmt.Printf("address of balance in Deposit is %v \n", &w.balance)
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance // Normally it should be (*w).balance but makers of Go made without explicit dereference
}
