package main

type Transaction struct {
	Date   string
	Amount int
}

type TransactionRepository interface {
	AddTransaction(tx Transaction)
	AllTransactions() []Transaction
}

type InMemoryTransactionRepository struct {
}

func (r *InMemoryTransactionRepository) AddTransaction(tx Transaction) {
	panic("not implemented")
}

func (r *InMemoryTransactionRepository) AllTransactions() []Transaction {
	panic("not implemented")
}

type MockTransationRepository struct {
	AddTransactionFn  func(Transaction)
	AllTransactionsFn func() []Transaction
}

func (r *MockTransationRepository) AddTransaction(tx Transaction) {
	r.AddTransactionFn(tx)
}

func (r *MockTransationRepository) AllTransactions() []Transaction {
	return r.AllTransactionsFn()
}
