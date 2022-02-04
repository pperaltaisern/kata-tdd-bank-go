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