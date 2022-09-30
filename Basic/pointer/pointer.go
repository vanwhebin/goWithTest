package pointer

import "fmt"

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Ballance() Bitcoin {
	return w.balance
}

// 存钱
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

// 取钱
func (w *Wallet) Withdraw(amount Bitcoin) {
	w.balance -= amount
}

//  类型别名
type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
