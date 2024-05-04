package err

import (
	"errors"
	"fmt"
)

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

func (w *Wallet) Withdraw(bitcoin Bitcoin) error {
	if w.balance < bitcoin {
		return errors.New("oh no")
	}
	w.balance -= bitcoin
	return nil
}
