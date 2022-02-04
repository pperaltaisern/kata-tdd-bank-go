package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAccount_Deposit_ShouldStoreATransaction(t *testing.T) {
	// arrange
	expectedTransaction := Transaction{
		Date:   "14/01/2012",
		Amount: 100,
	}

	f := &MockTransactionFactory{
		CreateTransactionFromDepositFn: func(amount int) Transaction {
			return Transaction{
				Date:   expectedTransaction.Date,
				Amount: amount,
			}
		},
	}

	repoCalls := 0
	var capturedTx Transaction
	r := &MockTransationRepository{
		AddTransactionFn: func(tx Transaction) {
			repoCalls++
			capturedTx = tx
		},
	}

	account := NewAccount(r, f)

	// act
	account.Deposit(100)

	// assert
	require.Equal(t, 1, repoCalls)
	require.Equal(t, expectedTransaction, capturedTx)
}

func TestAccount_Withdraw(t *testing.T) {

}

func TestAccount_PrintStatement(t *testing.T) {

}
