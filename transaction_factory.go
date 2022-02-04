package main

type TransactionFactory interface {
	CreateTransactionFromDeposit(amount int) Transaction
	CreateTransactionFromWithdrawal(amount int) Transaction
}

type transactionFactory struct {
}

func (f *transactionFactory) CreateTransactionFromDeposit(amount int) Transaction {
	return Transaction{}
}

func (f *transactionFactory) CreateTransactionFromWithdrawal(amount int) Transaction {
	return Transaction{}
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
