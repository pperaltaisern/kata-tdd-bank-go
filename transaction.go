package main

type Transaction struct {
	Date   string
	Amount int
}

func NewTransaction(date string, amount int) Transaction {
	return Transaction{
		Date:   date,
		Amount: amount,
	}
}
