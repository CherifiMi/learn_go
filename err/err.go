package err

import "fmt"

type Wallet struct {
	balance Bitcoin
}
type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

func (w *Wallet) Deposit(i Bitcoin) {
	w.balance += i
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) Withdraw(bitcoin Bitcoin) {
	w.balance -= bitcoin
}
