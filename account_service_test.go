package main

import (
	"github.com/cucumber/godog"
)

type printStatementFeature struct {
	accountService AccountService
}

func (f *printStatementFeature) aClientMakesADepositOfOn(deposit int, date string) error {
	return godog.ErrPending
}

func (f *printStatementFeature) aDepositOfOn(deposit int, date string) error {
	return godog.ErrPending
}

func (f *printStatementFeature) aWithdrawalOfOn(withdrawal int, date string) error {
	return godog.ErrPending
}

func (f *printStatementFeature) theyPrintTheirBankStatement() error {
	return godog.ErrPending
}

func (f *printStatementFeature) theyWouldSee(arg1 *godog.DocString) error {
	return godog.ErrPending
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	printStatementFeature := &printStatementFeature{}
	ctx.Step(`^a client makes a deposit of (\d+) on "([^"]*)"$`, printStatementFeature.aClientMakesADepositOfOn)
	ctx.Step(`^a deposit of (\d+) on "([^"]*)"$`, printStatementFeature.aDepositOfOn)
	ctx.Step(`^a withdrawal of (\d+) on "([^"]*)"$`, printStatementFeature.aWithdrawalOfOn)
	ctx.Step(`^they print their bank statement$`, printStatementFeature.theyPrintTheirBankStatement)
	ctx.Step(`^they would see:$`, printStatementFeature.theyWouldSee)
}
