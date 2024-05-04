package err

type Wallet struct {
	balance int
}

func (w *Wallet) Deposit(i int) {
	w.balance += i
}

func (w *Wallet) Balance() int {
	return w.balance
}
