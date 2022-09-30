package pointer

import (
	"errors"
	"fmt"
)

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

var ErrInsufficientFunds = errors.New("cantnot withdraw, insufficient funds")

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

// 存钱
func (w *Wallet) Deposit(amount Bitcoin) error {
	w.balance += amount
	return nil
}

// 取钱
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}
	w.balance -= amount
	return nil
}

//  类型别名
type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
