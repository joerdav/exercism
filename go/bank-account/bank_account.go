package account

import "sync"

// Define the Account type here.
type Account struct {
	balance int64
	open    bool
	lock    sync.RWMutex
}

func Open(amount int64) *Account {
	if amount < 0 {
		return nil
	}
	return &Account{
		balance: amount,
		open:    true,
	}
}

func (a *Account) Balance() (int64, bool) {
	a.lock.RLock()
	defer a.lock.RUnlock()
	if !a.open {
		return 0, false
	}
	return a.balance, true
}

func (a *Account) Deposit(amount int64) (int64, bool) {
	a.lock.Lock()
	defer a.lock.Unlock()
	if !a.open || a.balance+amount < 0 {
		return 0, false
	}
	a.balance += amount
	return a.balance, true
}

func (a *Account) Close() (int64, bool) {
	a.lock.Lock()
	defer a.lock.Unlock()
	prev := Account{balance: a.balance, open: a.open}
	a.open = false
	a.balance = 0
	return prev.balance, prev.open
}
