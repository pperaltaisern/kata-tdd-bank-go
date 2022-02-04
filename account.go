package main

type Account struct {
	transactionRepository TransactionRepository
	transactionFactory    TransactionFactory
}

func NewAccount(r TransactionRepository, f TransactionFactory) *Account {
	return &Account{
		transactionRepository: r,
		transactionFactory:    f,
	}
}

func (acc *Account) Deposit(amount int) {
	tx := acc.transactionFactory.CreateTransactionFromDeposit(amount)
	acc.transactionRepository.AddTransaction(tx)
}

func (acc *Account) Withdraw(amount int) {

}

func (acc *Account) PrintStatement() {

}
