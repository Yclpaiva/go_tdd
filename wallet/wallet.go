package wallet

import (
	"errors"
	"fmt"
)

type Bitcoin int

type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

// * is the pointer of the memory
func (w *Wallet) Deposit(amount Bitcoin) {
	// print memory value fmt.Println("address of balance in test is", &w.balance)
	(*w).balance += amount
}

var ErrInsufficientFunds = errors.New("cannot withdraw, balance lower than amount")

func (w *Wallet) Withdraw(amount Bitcoin) error {
	// print memory value fmt.Println("address of balance in test is", &w.balance)
	if amount > w.balance {
		return ErrInsufficientFunds
	}
	(*w).balance -= amount
	return nil
}

func (w *Wallet) Balance() string {
	// print memory value fmt.Println("address of balance in test is", &w.balance)
	return w.balance.String() // is the same of (*w).balance
}
