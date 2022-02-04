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

	account := NewAccount(r, f, nil)

	// act
	account.Deposit(100)

	// assert
	require.Equal(t, 1, repoCalls)
	require.Equal(t, expectedTransaction, capturedTx)
}

func TestAccount_Withdraw_ShouldStoreATransaction(t *testing.T) {
	// arrange
	expectedTransaction := Transaction{
		Date:   "14/01/2012",
		Amount: -100,
	}

	f := &MockTransactionFactory{
		CreateTransactionFromWithdrawalFn: func(amount int) Transaction {
			return Transaction{
				Date:   expectedTransaction.Date,
				Amount: -amount,
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

	account := NewAccount(r, f, nil)

	// act
	account.Withdraw(100)

	// assert
	require.Equal(t, 1, repoCalls)
	require.Equal(t, expectedTransaction, capturedTx)
}

func TestAccount_PrintStatement_ShouldPrintABatchOfTransactions(t *testing.T) {
	// arrange
	expectedTransactions := []Transaction{
		{
			Date:   "14/01/2012",
			Amount: -100,
		},
	}

	r := &MockTransationRepository{
		AllTransactionsFn: func() []Transaction {
			return expectedTransactions
		},
	}

	printerCalls := 0
	var capturedTxs []Transaction
	sp := &MockStatementPrinter{
		PrintStatementFn: func(txs []Transaction) {
			printerCalls++
			capturedTxs = txs
		},
	}

	account := NewAccount(r, nil, sp)

	// act
	account.PrintStatement()

	// assert
	require.Equal(t, 1, printerCalls)
	require.Equal(t, expectedTransactions, capturedTxs)
}
