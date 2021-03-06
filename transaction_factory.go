package main

type TransactionFactory interface {
	CreateTransactionFromDeposit(amount int) Transaction
	CreateTransactionFromWithdrawal(amount int) Transaction
}

type transactionFactory struct {
	clock Clock
}

func NewTransactionFactory() *transactionFactory {
	return NewTransactionFactoryWithClock(NewClock())
}

func NewTransactionFactoryWithClock(clock Clock) *transactionFactory {
	return &transactionFactory{
		clock: clock,
	}
}

func (f *transactionFactory) CreateTransactionFromDeposit(amount int) Transaction {
	date := f.clock.NowString()
	return NewTransaction(date, amount)
}

func (f *transactionFactory) CreateTransactionFromWithdrawal(amount int) Transaction {
	date := f.clock.NowString()
	return NewTransaction(date, -amount)
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
