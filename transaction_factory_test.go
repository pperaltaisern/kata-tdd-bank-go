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
	clock := testClock(expectedTx.Date)

	f := NewTransactionFactoryWithClock(clock)

	// Act
	tx := f.CreateTransactionFromDeposit(expectedTx.Amount)
	// Assert
	require.Equal(t, expectedTx, tx)
}

func TestTransactionFactory_CreateTransactionFromWithdrawal(t *testing.T) {
	// Arrange
	amount := 100
	expectedTx := Transaction{
		Amount: -100,
		Date:   "13/01/2012",
	}
	clock := testClock(expectedTx.Date)

	f := NewTransactionFactoryWithClock(clock)

	// Act
	tx := f.CreateTransactionFromWithdrawal(amount)
	// Assert
	require.Equal(t, expectedTx, tx)
}

func testClock(date string) Clock {
	now := func() time.Time {
		t, err := time.Parse("02/01/2006", date)
		if err != nil {
			panic(err)
		}
		return t
	}
	return NewClockWithNowFunc(now)
}
