package bank

import "errors"

type Account struct {
    balance float64
}

func (a *Account) Deposit(amount float64) error {
    if amount <= 0 {
        return errors.New("invalid deposit amount")
    }
    a.balance += amount
    return nil
}

func (a *Account) Withdraw(amount float64) error {
    if amount <= 0 {
        return errors.New("invalid withdrawal amount")
    }
    if amount > a.balance {
        return errors.New("insufficient funds")
    }
    a.balance -= amount
    return nil
}

func (a *Account) GetBalance() float64 {
    return a.balance
}