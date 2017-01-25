// Package account contains implementation of an account
package account

import "sync"

// Account is a representation of an account
type Account struct {
	mutex   sync.Mutex
	closed  bool
	balance int64
}

// Open a new account
func Open(initalDeposit int64) *Account {
	if initalDeposit < 0 {
		return nil
	}
	return &Account{
		balance: initalDeposit,
	}
}

// Close the account
func (a *Account) Close() (payout int64, ok bool) {
	a.mutex.Lock()
	defer a.mutex.Unlock()

	if a.closed {
		return
	}

	a.closed = true
	payout = a.balance
	ok = true
	return
}

// Balance of the account
func (a *Account) Balance() (balance int64, ok bool) {
	a.mutex.Lock()
	defer a.mutex.Unlock()

	if a.closed {
		return
	}

	ok = true
	balance = a.balance
	return
}

// Deposit or withdraw from the account
func (a *Account) Deposit(amount int64) (newBalance int64, ok bool) {
	a.mutex.Lock()
	defer a.mutex.Unlock()

	if a.closed {
		return
	}

	if amount < 0 && a.balance+amount < 0 {
		return
	}

	ok = true
	a.balance += amount
	newBalance = a.balance
	return
}
