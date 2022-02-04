package main

type Account struct {
	transactionRepository TransactionRepository
	transactionFactory    TransactionFactory
	statementPrinter      StatementPrinter
}

func NewAccount(r TransactionRepository, f TransactionFactory, sp StatementPrinter) *Account {
	return &Account{
		transactionRepository: r,
		transactionFactory:    f,
		statementPrinter:      sp,
	}
}

func (acc *Account) Deposit(amount int) {
	tx := acc.transactionFactory.CreateTransactionFromDeposit(amount)
	acc.transactionRepository.AddTransaction(tx)
}

func (acc *Account) Withdraw(amount int) {
	tx := acc.transactionFactory.CreateTransactionFromWithdrawal(amount)
	acc.transactionRepository.AddTransaction(tx)
}

func (acc *Account) PrintStatement() {
	txs := acc.transactionRepository.AllTransactions()
	acc.statementPrinter.PrintStatement(txs)
}
