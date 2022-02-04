package main

import "time"

type TransactionFactory interface {
	CreateTransactionFromDeposit(amount int) Transaction
	CreateTransactionFromWithdrawal(amount int) Transaction
}

type transactionFactory struct {
	now func() time.Time
}

func NewTransactionFactory() *transactionFactory {
	return NewTransactionFactoryWithNowFunc(time.Now)
}

func NewTransactionFactoryWithNowFunc(now func() time.Time) *transactionFactory {
	return &transactionFactory{
		now: now,
	}
}

func (f *transactionFactory) CreateTransactionFromDeposit(amount int) Transaction {
	panic("not implemented")
}

func (f *transactionFactory) CreateTransactionFromWithdrawal(amount int) Transaction {
	panic("not implemented")
}

type MockTransactionFactory struct {
	CreateTransactionFromDepositFn    func(amount int) Transaction
	CreateTransactionFromWithdrawalFn func(amount int) Transaction
}

func (f *MockTransactionFactory) CreateTransactionFromDeposit(amount int) Transaction {
	return f.CreateTransactionFromDepositFn(amount)
}

func (f *MockTransactionFactory) CreateTransactionFromWithdrawal(amount int) Transaction {
	return f.CreateTransactionFromWithdrawalFn(amount)
}
