package main

type Transaction struct {
	Date   string
	Amount int
}

type TransactionRepository interface {
	AddTransaction(tx Transaction)
}

type InMemoryTransactionRepository struct {
}

func (r *InMemoryTransactionRepository) AddTransaction(tx Transaction) {
	panic("not implemented")
}

type MockTransationRepository struct {
	AddTransactionFn func(Transaction)
}

func (r *MockTransationRepository) AddTransaction(tx Transaction) {
	r.AddTransactionFn(tx)
}
