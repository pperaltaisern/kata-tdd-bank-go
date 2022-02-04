package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/cucumber/godog"
)

type printStatementFeature struct {
	account *Account
	// buffer is injected to Account as a writer
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
	r := &InMemoryTransactionRepository{}
	clock := acceptanceClockTest()
	f := NewTransactionFactoryWithClock(clock)
	sp := NewStatementPrinter(sb)
	printStatementFeature := &printStatementFeature{
		account: NewAccount(r, f, sp),
		buffer:  sb,
	}

	ctx.Step(`^a client makes a deposit of (\d+) on "([^"]*)"$`, printStatementFeature.aClientMakesADepositOfOn)
	ctx.Step(`^a deposit of (\d+) on "([^"]*)"$`, printStatementFeature.aDepositOfOn)
	ctx.Step(`^a withdrawal of (\d+) on "([^"]*)"$`, printStatementFeature.aWithdrawalOfOn)
	ctx.Step(`^they print their bank statement$`, printStatementFeature.theyPrintTheirBankStatement)
	ctx.Step(`^they would see:$`, printStatementFeature.theyWouldSee)
}

func acceptanceClockTest() Clock {
	calls := 1
	now := func() time.Time {
		var d string
		switch calls {
		case 1:
			d = "10/01/2012"
		case 2:
			d = "13/01/2012"
		case 3:
			d = "14/01/2012"
		}
		t, err := time.Parse("02/01/2006", d)
		if err != nil {
			panic(err)
		}
		calls++
		return t
	}
	return NewClockWithNowFunc(now)
}
