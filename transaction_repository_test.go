package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTransactionFactory_StoreAndRetrieveTransaction(t *testing.T) {
	// Arrange
	expectedTxs := []Transaction{
		NewTransaction("13/01/2012", 100),
		NewTransaction("14/01/2012", 200),
	}
	r := &InMemoryTransactionRepository{}

	// Act
	r.AddTransaction(expectedTxs[0])
	r.AddTransaction(expectedTxs[1])
	txs := r.AllTransactions()

	// Assert
	require.Equal(t, expectedTxs, txs)
}
