# Zero, One, GO!
-----
# Writing Tests - An introduction to test driven development

Test driven development is as simple as it sounds - we will first write tests and only then try to implement a method that behaves as test case specifies.

It often happens that methods that we write can not be tested against the real API. Also, testing against a real API would mean integration tests, and not unit tests. The real API can have an outage, and if we rely on it for our development, we will not be able to perform work in the period when API is not available. 

In other words, the unit tests will enable us to precisely define behaviour of our application and it will act as a contract for the methods we implement.

# Our first test case

If you look at [API specifications](../getting-started/api-specs.md) you will notice that the main object we work with is an Account object.

To create an account, we send an Account object, serialized as JSON, to the API. On our queries, API returns back an Account object(s).

This implies the need to create an explicit Account object.

## How to create test case

First step in creating test case is to create a file named `account_test.go`, which will contain tests for code in `account.go`.

We will try to create an Account object and serialize it to JSON.

Code in Go is divided into packages, and multiple packages form a library. We will begin our `account_test.go` test case with a package declaration.

```
package accountapi
```

To enable functions such as failing test cases and logging errors, we will include Go's package `testing` in our code. This is a standard syntax to import dependencies into your Go files.

```
import (
	"testing"
)
```

The actual test case should:
* Define expected account JSON
* Define Account object
* Serialize Account object into JSON and test if expected and actual match

To accomplish this, we first need to define our test case. The below snippet defines a test case. Test case is a function named `TestNlAccountToJson` which accepts as an argument `t` which is a pointer (defined by `*`) to `testing.T` package.

```
func TestNlAccountToJson(t *testing.T) {
    ...
}
```

Within `TestNlAccountToJson` we want to create an account and then serialize it to JSON. We still do not have an Account object defined, but the test case already shows how it will look like :)
* `var` identifies a variable. It's type is inferred to be `Account`
* `Account` identifies an Account object. Objects are always capitalized.
* All attributes between curly braces `Account {}` are capitalized (`Type`, `ID`...). This is necessary because otherwise, the attributes would not be serialized into JSON. Reason is that capitalized attributes in Go are `public` and accessible outside of the scope of object or method. 

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

We see here the new operator which is `;=`. This operator initialises variables and immediatelly assigns value to it. Values comes from the right statement which is `validNlAccount.toJSON()`.

Method `toJSON` will return something that looks like a tuple, but is essentially a pair of values. Tuples in Go are not first-class citizens. First-class citizens are functionalities whose support is built in a language.

So now we have account Jason which is of type account and we have error variable. The error variable shby convention should contain error object with a message on what exactly failed.

This is a solution in goal on how to handle errors in your application. Instead of using exceptions, Go applications (or actually go methods) return error object to caller so that caller can decide what to do with it.

```
expected := `{"type":"accounts","id":"DBC1D4F4-82E7-4286-B169-1D1B7D0D3989","organisation_id":"934AD475-6612-44A1-A0E8-6D74408780BC","attributes":{"country":"NL","base_currency":"EUR"}}`
```

Note in the above expected value (of inferred type String) that not all attributes that are mentioned in the API specifications are present in the response. This is because nil attributes are not serialised in JSON. Nil is a Null in Go.

Every time we call a method that in signature returns a value pair, we will check if `err` is nil.

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

Final test case can be seen [here](../src/accountapi/account_test.go)