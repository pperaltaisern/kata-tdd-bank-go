package main

type TransactionRepository interface {
	AddTransaction(tx Transaction)
	AllTransactions() []Transaction
}

type InMemoryTransactionRepository struct {
	mem []Transaction
}

func (r *InMemoryTransactionRepository) AddTransaction(tx Transaction) {
	r.mem = append(r.mem, tx)
}

func (r *InMemoryTransactionRepository) AllTransactions() []Transaction {
	return r.mem
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
