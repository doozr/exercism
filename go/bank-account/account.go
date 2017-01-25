// Package account contains implementation of an account
package account

import "sync"

// Account is a representation of an account
type Account struct {
	mutex   sync.Mutex
	closed  bool
	balance int64
}

type adjustment func() bool

// Open a new account
func Open(initalDeposit int64) *Account {
	if initalDeposit < 0 {
		return nil
	}
	return &Account{
		balance: initalDeposit,
	}
}

// access the account and optionally make an adjustment
func (a *Account) access(fn adjustment) (balance int64, ok bool) {
	a.mutex.Lock()
	defer a.mutex.Unlock()

	if a.closed {
		return
	}

	if fn != nil && !fn() {
		return
	}

	return a.balance, true
}

// Close the account
func (a *Account) Close() (payout int64, ok bool) {
	return a.access(func() bool {
		a.closed = true
		return true
	})
}

// Balance of the account
func (a *Account) Balance() (balance int64, ok bool) {
	return a.access(nil)
}

// Deposit or withdraw from the account
func (a *Account) Deposit(amount int64) (newBalance int64, ok bool) {
	return a.access(func() bool {
		if amount < 0 && a.balance+amount < 0 {
			return false
		}

		a.balance += amount
		return true
	})
}
