package main

type AccountService interface {
	Deposit(amount int)
	Withdraw(amount int)
	PrintStatement()
}
