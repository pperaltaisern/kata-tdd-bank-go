package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestTransactionFactory_CreateTransactionFromDeposit(t *testing.T) {
	// Arrange
	expectedTx := Transaction{
		Amount: 100,
		Date:   "13/01/2012",
	}
	nowFn := testNowFunc(expectedTx.Date)

	f := NewTransactionFactoryWithNowFunc(nowFn)

	// Act
	tx := f.CreateTransactionFromDeposit(expectedTx.Amount)
	// Assert
	require.Equal(t, expectedTx, tx)
}

func TestTransactionFactory_CreateTransactionFromWithdrawal(t *testing.T) {
	// Arrange
	expectedTx := Transaction{
		Amount: -100,
		Date:   "13/01/2012",
	}
	nowFn := testNowFunc(expectedTx.Date)

	f := NewTransactionFactoryWithNowFunc(nowFn)

	// Act
	tx := f.CreateTransactionFromWithdrawal(expectedTx.Amount)
	// Assert
	require.Equal(t, expectedTx, tx)
}

func testNowFunc(date string) func() time.Time {
	return func() time.Time {
		t, err := time.Parse("01/02/2006", date)
		if err != nil {
			panic(err)
		}
		return t
	}
}
