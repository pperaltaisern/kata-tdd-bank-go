package main

type Account struct {
	transactionRepository TransactionRepository
	transactionFactory    TransactionFactory
}

func NewAccount(r TransactionRepository, f TransactionFactory) *Account {
	return &Account{
		transactionRepository: r,
	}
}

func (acc *Account) Deposit(amount int) {

}

func (acc *Account) Withdraw(amount int) {

}

func (acc *Account) PrintStatement() {

}
