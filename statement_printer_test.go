package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStatementPrinter_PrintStatement_printsInExpectedFormat(t *testing.T) {
	// Arrange
	expected := `Date || Amount || Balance
14/01/2012 || -500 || 2500
13/01/2012 || 2000 || 3000
10/01/2012 || 1000 || 1000`

	txs := []Transaction{
		NewTransaction("10/01/2012", 1000),
		NewTransaction("13/01/2012", 2000),
		NewTransaction("14/01/2012", -500),
	}
	b := &strings.Builder{}
	sp := NewStatementPrinter(b)

	//Act
	sp.PrintStatement(txs)

	// Asset
	require.Equal(t, expected, b.String())
}
