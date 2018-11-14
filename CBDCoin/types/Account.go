package types

import ()

type Account struct {
	Address string
	Balance int
	Nonce   int
}

func NewAccount(addr string) *Account {
	ac := &Account{
		Address: addr,
		Balance: 0,
		Nonce:   0,
	}
	return ac
}

func (ac *Account) Add(receive int) *Account {
	ac.Balance += receive
	return ac
}

func (ac *Account) Checkpay(payment int, newnonce int) bool {
	if payment > ac.Balance || newnonce <= ac.Nonce {
		return false
	}
	return true
}

func (ac *Account) Pay(payment int, newnonce int) *Account {
	ac.Balance = ac.Balance - payment
	ac.Nonce = newnonce
	return ac
}
