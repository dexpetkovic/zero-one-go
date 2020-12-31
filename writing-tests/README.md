# Zero, One, GO!
-----
# Writing Tests - An introduction to test driven development

Test driven development is as simple as it sounds - we will first write tests and only then try to implement a method that behaves as test case specifies.

It often happens that methods that we write can not be tested against the real API. Also, testing against a real API would mean integration tests, and not unit tests. The real API can have an outage, and if we rely on it for our development, we will not be able to perform work in the period when API is not available. 

In other words, the unit tests will enable us to precisely define behaviour of our application and it will act as a contract for the methods we implement.

# What are we testing?

We will try to cover all required functionality of the AccountAPI library with the test cases. In other words, we will test the functionality of the library using unit test cases. We will mock Bank's API in these tests.

# Our first test case

If you look at [Bank API specifications](../getting-started/api-specs.md) you will notice that the main object we work with is an Account object.

To create an account in Bank's API, we need to send an Account object, serialized as JSON, to the API. And on our queries, Bank's API returns back Account object(s) which we have to return back to the caller.

This implies the need to create an explicit Account object.

## Creating Account test case

First step in creating a test case for Account is to create a file named `account_test.go`, which will contain tests for code in `account.go`.

The test will be very simple since we will just try to create an Account object and serialize it to JSON.

Code in Go is divided into packages, and multiple packages form a library. This helps to better organise the code that we work with. We will begin our `account_test.go` test case with a package declaration.

```
package accountapi
```

To enable functions such as failing test cases and logging errors, we will include Go's package `testing` in our code. This is a standard syntax to import dependencies into your Go files.

```
import (
    "testing"
)
```

The test case should:
* Define an expected Account JSON
* Define Account object
* Serialize Account object into JSON and test if expected and actual match

To accomplish this, we first need to define our test case. The below snippet defines a test case. Test case is a function named `TestNlAccountToJson` which accepts as an argument `t` which is a pointer (defined by `*`) to `testing.T` package.

```
func TestNlAccountToJson(t *testing.T) {
    ...
}
```

Within `TestNlAccountToJson` test, we want to create an account and then serialize it to JSON.

We still do not have an Account object defined, but the test case already shows how it will look like :)

* `var` identifies a variable. It's type is inferred to be `Account`
* `Account` identifies an Account object. Objects are always capitalized.
* All attributes between curly braces `Account {}` are capitalized (`Type`, `ID`...). This is necessary because otherwise, the attributes would not be serialized into JSON. Reason is that only capitalized attributes in Go are `public` and accessible outside of the scope of object or method. 

```
var validNlAccount = Account {
    Type:           "accounts",
    ID:             "DBC1D4F4-82E7-4286-B169-1D1B7D0D3989",
    OrganisationID: "934AD475-6612-44A1-A0E8-6D74408780BC",
    Attributes: countryAttributes{
        Country:      "NL",
        BaseCurrency: "EUR",
    },
}
```

Next step is to serialize this Account with method `toJSON()`. We haven't written this method - yet :)

```
accountJSON, err := validNlAccount.toJSON()
```

We see here a new operator which is `;=`. This operator initialises variables and immediatelly assigns value to it. Values comes from the right statement which is `validNlAccount.toJSON()`.

Method `toJSON` will return something that looks like a tuple, but is essentially a pair of values. Tuples in Go are not first-class citizens. First-class citizens are functionalities whose support is built in a language.

So now we have account JSON which is of type account and we have error variable. The error variable should by convention contain error object with a message on what exactly failed. This is a solution in Go on how to handle errors in your application. Instead of using exceptions, Go applications (or actually Go methods) return error object to caller, so that caller can decide what to do with it.

We will define a expected value for our test case.

```
expected := `{"type":"accounts","id":"DBC1D4F4-82E7-4286-B169-1D1B7D0D3989","organisation_id":"934AD475-6612-44A1-A0E8-6D74408780BC","attributes":{"country":"NL","base_currency":"EUR"}}`
```

Note in the above expected value (of inferred type String, by `:=` operator) that not all attributes that are mentioned in the API specifications are present in the response. This is because `nil` attributes are not serialised in JSON. `Nil` is a `Null` in Go.

Every time we call a method that in signature returns a value pair, we will check if `err` is `nil`.

```
if err != nil {
    t.Errorf("Error casting Account to Json: %v", err)
}

if accountJSON != expected {
    t.Errorf("Casted account does not match: %v", accountJSON)
}
```

In the first check above we will check if we managed to cast account object to JSON.

In the second check, we check whether the expected Account JSON is the same as the actual one (the one returned by our `toJSON()` method). The compared types here are String.

Final test case for Account `toJSON()` can be seen [here](../src/accountapi/account_test.go).

## Creating Service test case

The Account object will be used by the Service, which will implement create, fetch, list and delete account functionality.

### Test create account

Create Account test case will begin with package declaration and importing of required dependencies. Fmt is used for formatting, reflect for object reflection, sort to perform sorting of results and testing, a standard Go test library. 

```
package accountapi

import (
    "fmt"
    "reflect"
    "sort"
    "testing"
)
```

Test cases for [service](../src/accountapi/service_test.go) cover various scenarios for testing account creation. We will focus on creating a valid NL bank account. First we will define the test case itself:

```
func TestCreateValidNlAccount(t *testing.T) {
    ...
}
```

After that, we will call validNlAccount.Create() method, which should create the actual account, `createdAcc`:
```
    createdAcc, err := validNlAccount.Create()
```

When we call `Create()`, it will create an account in Bank's API, and return back `Account` object. We have to validate that upon creation, create method returned a correct Account type. We compare the `Type` and `ID` of the created account, and compare it with the values from `validNlAccount`.

```
    if err != nil ||
        createdAcc.Type != validNlAccount.Type ||
        createdAcc.ID != validNlAccount.ID {
        t.Errorf("Error while creating account: %v", err)
    }
```

### Test fetch accounts

To test fetching accounts, we will use ID of account we created in tests above.

func TestFetchAccount(t *testing.T) {
    var accountID = "ad27e265-9605-4b4b-a0e5-3003ea9cc4dc"
}


When we call `FetchAccount()` we will get back `Account` object within `fetchedAccount` and standard `Error` object, `err`.

```
    fetchedAccount, err := FetchAccount(accountID)
```

Depending on the `err` value, we know that fetch succeeded if value of `err` is `nil`.

```
if err != nil {
    t.Errorf("Error while fetching account id %v: %v", accountID, err)
}

Finally we expect the IDs to be the same. If not, we will fail the test case.

if fetchedAccount.ID != accountID {
    t.Errorf("Fetched account does not have the same id: %v and %v.", fetchedAccount.ID, accountID)
}

### Test list accounts





Final test case for account creation service can be seen [here](../src/accountapi/service_test.go).