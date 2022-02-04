package main

import (
	"io"
	"strconv"
)

type StatementPrinter interface {
	PrintStatement([]Transaction)
}

type statementPrinter struct {
	w io.Writer
}

func NewStatementPrinter(w io.Writer) *statementPrinter {
	return &statementPrinter{
		w: w,
	}
}

func (sp *statementPrinter) PrintStatement(txs []Transaction) {
	lines := make([]string, 0, len(txs))
	balance := 0
	for _, tx := range txs {
		balance += tx.Amount
		l := transactionStatement(tx, balance)
		lines = append(lines, l)
	}

	sp.w.Write([]byte("Date || Amount || Balance\n"))
	for i := len(lines) - 1; 0 <= i; i-- {
		sp.w.Write([]byte(lines[i]))
		if i != 0 {
			sp.w.Write([]byte("\n"))
		}
	}
}

func transactionStatement(tx Transaction, balance int) string {
	return tx.Date + " || " + strconv.FormatInt(int64(tx.Amount), 10) + " || " + strconv.FormatInt(int64(balance), 10)
}

type MockStatementPrinter struct {
	PrintStatementFn func([]Transaction)
}

func (f *MockStatementPrinter) PrintStatement(txs []Transaction) {
	f.PrintStatementFn(txs)
}
