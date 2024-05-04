package err

type Wallet struct {
	balance Bitcoin
}
type Bitcoin int

func (w *Wallet) Deposit(i Bitcoin) {
	w.balance += i
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
