package main

import (
	"errors"
	"fmt"
)

// Stringer は文字列の出力のinterfaceです。
type Stringer interface {
	String() string
}

// Bitcoin はbitコインを表す
type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// Wallet は財布を表す
type Wallet struct {
	balance Bitcoin
}

// Deposit は入金
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

// Balance は残高
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

// ErrInsufficientFunds はwithdrawのエラーメッセージです。
var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

// Withdraw は預金を撤回する
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= amount
	return nil
}
