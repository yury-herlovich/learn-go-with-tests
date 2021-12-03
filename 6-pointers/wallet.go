package wallet

import (
	"errors"
	"fmt"
)

var ErrInsufficientFunds = errors.New("insufficient funds")

type Stringer interface {
	String() string
}

type Dollar int

func (b Dollar) String() string {
	return fmt.Sprintf("%d USD", b)
}

type Wallet struct {
	balance Dollar
}

func (w *Wallet) Deposit(amount Dollar) {
	w.balance += amount
}

func (w *Wallet) Withdraw(amount Dollar) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= amount
	return nil
}

func (w *Wallet) Balance() Dollar {
	return w.balance
}
