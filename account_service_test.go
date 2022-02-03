package main

import (
	"fmt"
	"strings"

	"github.com/cucumber/godog"
)

type printStatementFeature struct {
	account *Account
	// buffer is injected to AccountService as a writer
	buffer *strings.Builder
}

func (f *printStatementFeature) aClientMakesADepositOfOn(amount int, date string) error {
	f.account.Deposit(amount)
	return nil
}

func (f *printStatementFeature) aDepositOfOn(amount int, date string) error {
	f.account.Deposit(amount)
	return nil
}

func (f *printStatementFeature) aWithdrawalOfOn(amount int, date string) error {
	f.account.Withdraw(amount)
	return nil
}

func (f *printStatementFeature) theyPrintTheirBankStatement() error {
	f.account.PrintStatement()
	return nil
}

func (f *printStatementFeature) theyWouldSee(expectedStatement *godog.DocString) error {
	actual := f.buffer.String()
	if actual != expectedStatement.Content {
		return fmt.Errorf("expected print statement does not match \nactual:\n %s \nexpected:\n %s", actual, expectedStatement.Content)
	}
	return nil
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	sb := &strings.Builder{}
	printStatementFeature := &printStatementFeature{
		account: NewAccount(),
		buffer:  sb,
	}

	ctx.Step(`^a client makes a deposit of (\d+) on "([^"]*)"$`, printStatementFeature.aClientMakesADepositOfOn)
	ctx.Step(`^a deposit of (\d+) on "([^"]*)"$`, printStatementFeature.aDepositOfOn)
	ctx.Step(`^a withdrawal of (\d+) on "([^"]*)"$`, printStatementFeature.aWithdrawalOfOn)
	ctx.Step(`^they print their bank statement$`, printStatementFeature.theyPrintTheirBankStatement)
	ctx.Step(`^they would see:$`, printStatementFeature.theyWouldSee)
}
