package main

type StatementPrinter interface {
	PrintStatement([]Transaction)
}

type statementPrinter struct{}

func (sp *statementPrinter) PrintStatement(txs []Transaction) {
	panic("not implemented")
}

type MockStatementPrinter struct {
	PrintStatementFn func([]Transaction)
}

func (f *MockStatementPrinter) PrintStatement(txs []Transaction) {
	f.PrintStatementFn(txs)
}
