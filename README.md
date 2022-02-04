# kata-tdd-bank-go
https://katalyst.codurance.com/bank


## Creating a failing acceptance test 
We write the acceptance criteria for our "print statement" feature, which is used to generate the code for our acceptance test. Then we create the Account struct with empty methods so we can write the test. When running the test, we see that it is failing:

``` gherkin
$ godog
Feature: print statement
  In order to see an account state
  As a client
  I need to be able to see my transactions

  Scenario: Should get deposit and withdrawal transactions # features\print_statement.feature:6
    Given a client makes a deposit of 1000 on "10-01-2012" # account_service_test.go:16 -> *printStatementFeature
    And a deposit of 2000 on "13-01-2012"                  # account_service_test.go:21 -> *printStatementFeature
    And a withdrawal of 500 on "14-01-2012"                # account_service_test.go:26 -> *printStatementFeature
    When they print their bank statement                   # account_service_test.go:31 -> *printStatementFeature
    Then they would see:                                   # account_service_test.go:36 -> *printStatementFeature
      """
                Date       || Amount || Balance
                14/01/2012 || -500   || 2500
                13/01/2012 || 2000   || 3000
                10/01/2012 || 1000   || 1000
      """
    expected print statement does not match
actual:

expected:
                Date       || Amount || Balance
                14/01/2012 || -500   || 2500
                13/01/2012 || 2000   || 3000
                10/01/2012 || 1000   || 1000
```


## Failing Unit tests for Account
The acceptance test will guide us, now we have empty methods for Account so the next step is to create unit tests for it. Starting unit testing the "Deposit" method, we see that there are this method has to do two different things:
1. Store transactions (date, amount)
2. Create transactions from an amount

Therefore, we create two different objects, one per responsability.
1. For storing transactions, we create a TransactionsRepository
2. For creating transactions, we create a TransactionsFactory

We mock them and build our test, we expect the repository to be called with a expected transaction. When running it, it fails since it's not already implemented:

``` gherkin
--- FAIL: TestAccount_Deposit_ShouldStoreATransaction (0.00s)
    account_test.go:40:
                Error Trace:    account_test.go:40
                Error:          Not equal:
                                expected: 1
                                actual  : 0
                Test:           TestAccount_Deposit_ShouldStoreATransaction
```

## Make unit tests for Account pass

We implement the Account's Deposit method:

```go
func (acc *Account) Deposit(amount int) {
	tx := acc.transactionFactory.CreateTransactionFromDeposit(amount)
	acc.transactionRepository.AddTransaction(tx)
}
```

Now, the test is passing:

``` gherkin
go test ./... -v
=== RUN   TestAccount_Deposit_ShouldStoreATransaction
--- PASS: TestAccount_Deposit_ShouldStoreATransaction (0.00s)
```


The result is a clean Account struct that has only high level logic. Each of his responsabilities are delegated to smaller objects.

We do the same for the rest of the Account's methods, the Withdraw is similar to the Deposit. The PrintStatement gives the Action class a third responsability, that is printing an array of transactions, and we are creating another interface for it.

At the end, all unit tests for Account pass, so we are done with it.
```
go test ./... -v
=== RUN   TestAccount_Deposit_ShouldStoreATransaction
--- PASS: TestAccount_Deposit_ShouldStoreATransaction (0.00s)
=== RUN   TestAccount_Withdraw_ShouldStoreATransaction
--- PASS: TestAccount_Withdraw_ShouldStoreATransaction (0.00s)
=== RUN   TestAccount_PrintStatement_ShouldPrintABatchOfTransactions
--- PASS: TestAccount_PrintStatement_ShouldPrintABatchOfTransactions (0.00s)
```

## Let the acceptance test guide us

After having a valid Account implementation, the acceptance test is panicking because there are still classes that have not been implemented. In this case, it's the transactionFactory

``` gherkin
godog
Feature: print statement
  In order to see an account state
  As a client
  I need to be able to see my transactions

  Scenario: Should get deposit and withdrawal transactions # features\print_statement.feature:6
    Given a client makes a deposit of 1000 on "10-01-2012" # account_acceptance_test.go:16 -> *printStatementFeature
    not implemented
runtime.gopanic
        c:/go/src/runtime/panic.go:1038
github.com/pperaltaisern/kata-tdd-bank-go.(*transactionFactory).CreateTransactionFromDeposit
        C:/Users/PepPeralta/kata-tdd-bank-go/transaction_factory.go:12
```

The next step is to do the same that we did for the Account, create failing unit test for the factory, then implement the methods. Repeat this process until the acceptance test passes.

## Acceptance test passed

``` gherkin
godog
Feature: print statement
  In order to see an account state
  As a client
  I need to be able to see my transactions

  Scenario: Should get deposit and withdrawal transactions # features\print_statement.feature:6
    Given a client makes a deposit of 1000 on "10-01-2012" # account_acceptance_test.go:17 -> *printStatementFeature
    And a deposit of 2000 on "13-01-2012"                  # account_acceptance_test.go:22 -> *printStatementFeature
    And a withdrawal of 500 on "14-01-2012"                # account_acceptance_test.go:27 -> *printStatementFeature
    When they print their bank statement                   # account_acceptance_test.go:32 -> *printStatementFeature
    Then they would see:                                   # account_acceptance_test.go:37 -> *printStatementFeature
      """
      Date || Amount || Balance
      14/01/2012 || -500 || 2500
      13/01/2012 || 2000 || 3000
      10/01/2012 || 1000 || 1000
      """

1 scenarios (1 passed)
5 steps (5 passed)
```

All unit tests passing with a 94.7% of coverage

```
go test ./... -v -cover
=== RUN   TestAccount_Deposit_ShouldStoreATransaction
--- PASS: TestAccount_Deposit_ShouldStoreATransaction (0.00s)
=== RUN   TestAccount_Withdraw_ShouldStoreATransaction
--- PASS: TestAccount_Withdraw_ShouldStoreATransaction (0.00s)
=== RUN   TestAccount_PrintStatement_ShouldPrintABatchOfTransactions
--- PASS: TestAccount_PrintStatement_ShouldPrintABatchOfTransactions (0.00s)
=== RUN   TestClock_NowString_ShouldFormatAsExpected
--- PASS: TestClock_NowString_ShouldFormatAsExpected (0.00s)
=== RUN   TestStatementPrinter_PrintStatement_printsInExpectedFormat
--- PASS: TestStatementPrinter_PrintStatement_printsInExpectedFormat (0.00s)
=== RUN   TestTransactionFactory_CreateTransactionFromDeposit
--- PASS: TestTransactionFactory_CreateTransactionFromDeposit (0.00s)
=== RUN   TestTransactionFactory_CreateTransactionFromWithdrawal
--- PASS: TestTransactionFactory_CreateTransactionFromWithdrawal (0.00s)
=== RUN   TestTransactionFactory_StoreAndRetrieveTransaction
--- PASS: TestTransactionFactory_StoreAndRetrieveTransaction (0.00s)
PASS
coverage: 94.7% of statements
ok      github.com/pperaltaisern/kata-tdd-bank-go       1.408s  coverage: 94.7% of statements
```