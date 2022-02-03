# kata-tdd-bank-go
https://katalyst.codurance.com/bank


## Creating a failing acceptance test 

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