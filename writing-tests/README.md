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

## Creating Integration Service test case

The Account object will be used by the Service, which will implement create, fetch, list and delete account functionality.

Integration Service tests are ran against the real service. These tests can be ran only if real endpoint exists, which can serve our requests. If such server is not available, we need to fake (or mock) the responses so that we can accurately test our service. See below for unit tests with mocked HTTP server.

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

After that, we will call validNlAccount.Create(testConfig) method, which should create the actual account, `createdAcc`:
```
    createdAcc, err := validNlAccount.Create(testConfig)
```

When we call `Create(testConfig)`, it will create an account in Bank's API, and return back `Account` object. We have to validate that upon creation, create method returned a correct Account type. We compare the `Type` and `ID` of the created account, and compare it with the values from `validNlAccount`.

The test configuration is available in configuration_test.go and is exposed via `testConfig`.

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


When we call `FetchAccount(testConfig)` we will get back `Account` object within `fetchedAccount` and standard `Error` object, `err`.

```
    fetchedAccount, err := FetchAccount(testConfig, accountID)
```

Depending on the `err` value, we know that fetch succeeded if value of `err` is `nil`.

```
if err != nil {
    t.Errorf("Error while fetching account id %v: %v", accountID, err)
}
```
Finally we expect the IDs to be the same. If not, we will fail the test case.
```
if fetchedAccount.ID != accountID {
    t.Errorf("Fetched account does not have the same id: %v and %v.", fetchedAccount.ID, accountID)
}
```

### Test list accounts

Testing of list accounts is a complex test. It requires us to call `ListAccounts` with `testConfig` that contains test configuration.

func TestListAccounts(t *testing.T) {
	listedAccounts, err := ListAccounts(testConfig, 0, 10)
	if err != nil {
		t.Errorf("Error while fetching accounts: %v", err)
	}
```

We expect to get a list of accounts, represented by `[]Account` type. To construct list of accounts, we use syntax `[]Account{}` where in `{}` we place Account objects we have already created.

```
	expected := []Account{validNlAccount, validUkAccount}
```

Comparison of the expected and actual list is tricky. We can not use reflect.DeepEqual as API returns created/modified timestamps which, in this app and its tests, change upon every test run.

```
	var matched []Account
	for i := 0; i < len(listedAccounts); i++ {
		for j := 0; j < len(expected); j++ {
			if expected[j].ID == listedAccounts[i].ID &&
				expected[j].OrganisationID == listedAccounts[i].OrganisationID &&
				expected[j].Attributes.BankID == listedAccounts[i].Attributes.BankID &&
				expected[j].Attributes.BankIDCode == listedAccounts[i].Attributes.BankIDCode &&
				expected[j].Attributes.Bic == listedAccounts[i].Attributes.Bic &&
				expected[j].Attributes.BaseCurrency == listedAccounts[i].Attributes.BaseCurrency &&
				expected[j].Attributes.Country == listedAccounts[i].Attributes.Country {
				matched = append(matched, expected[j])
			}
		}
	}

	// Try to sort matched accounts using custom comparator to
	// avoid test failure due to mismatched sorting
	sort.SliceStable(matched, func(i, j int) bool {
		return matched[i].Attributes.BankIDCode < matched[j].Attributes.BankIDCode
	})

	if !reflect.DeepEqual(matched, expected) {
		t.Errorf("Listed accounts do not match the expected ones: %v ", fmt.Sprint(listedAccounts))
	}
}

## Creating unit test cases for Service

Unit tests depend on mocked HTTP server. It is a bit more tricky to set it up.

We will cover all test cases for the account creation service within a single test. Within that test, we will first start mock http service, register handlers that will respond to our HTTP requests, and then run tests.

`http.NewServeMux()` creates a multiplexed HTTP server as `testRouter`. We will record response body in `testAccountResponseBody`.

```
func TestAccountCreationService(t *testing.T) {
	testRouter := http.NewServeMux()
	var testAccountResponseBody responseBody
```

Now we register handlers for HTTP URIs and generate the response. We cover for `/v1/organisation/accounts/` path for both `POST` and `GET`.

In `POST` requests, first we have to parse request body from the test, so that we know what kind of `testAccountRequestBody` to generate and return to calling test. 

```
	testRouter.HandleFunc("/v1/organisation/accounts/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {

			// Get request body, and based on account ID, return fake response body
			bReqBody, err := ioutil.ReadAll(r.Body)
			if err != nil {
				t.Errorf("Error while parsing test request body: %v", err)
			}
			err = json.Unmarshal(bReqBody, &testAccountResponseBody)

			w.Header().Set("Content-Type", "application/vnd.api+json")

			if testAccountResponseBody.Data.ID == validUkAccount.ID {
				w.WriteHeader(http.StatusCreated)
				w.Write([]byte(`{"data":{"attributes":{"alternative_bank_account_names":null,"bank_id":"400300","bank_id_code":"GBDSC","base_currency":"GBP","bic":"NWBKGB22","country":"GB"},"created_on":"2021-01-09T13:24:54.567Z","id":"ad27e265-9605-4b4b-a0e5-3003ea9cc4dc","modified_on":"2021-01-09T13:24:54.567Z","organisation_id":"eb0bd6f5-c3f5-44b2-b677-acd23cdde73c","type":"accounts","version":0},"links":{"self":"/v1/organisation/accounts/ad27e265-9605-4b4b-a0e5-3003ea9cc4dc"}}`))
			} else if testAccountResponseBody.Data.ID == validNlAccount.ID {
				w.WriteHeader(http.StatusCreated)
				w.Write([]byte(`{"data":{"attributes":{"alternative_bank_account_names":null,"base_currency":"EUR","bic":"NLABNA01","country":"NL"},"created_on":"2021-01-09T13:24:54.574Z","id":"bf33e333-9605-4b4b-a0e5-3003ea9cc4dc","modified_on":"2021-01-09T13:24:54.574Z","organisation_id":"eb0bd6f5-c3f5-44b2-b677-acd23cdde73c","type":"accounts","version":0},"links":{"self":"/v1/organisation/accounts/bf33e333-9605-4b4b-a0e5-3003ea9cc4dc"}}`))
			} else if testAccountResponseBody.Data.ID == invalidNlAccount.ID {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(`{"error_message":"validation failure list:\nvalidation failure list:\nvalidation failure list:\nbic in body should match '^([A-Z]{6}[A-Z0-9]{2}|[A-Z]{6}[A-Z0-9]{5})$'\nid in body must be of type uuid: \"1234-abcd\"\norganisation_id in body must be of type uuid: \"org-id\""}`))
			}
		} else if r.Method == "GET" {
			w.Header().Set("Content-Type", "application/vnd.api+json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"data":[{"attributes":{"alternative_bank_account_names":null,"bank_id":"400300","bank_id_code":"GBDSC","base_currency":"GBP","bic":"NWBKGB22","country":"GB"},"created_on":"2021-01-09T16:34:25.719Z","id":"ad27e265-9605-4b4b-a0e5-3003ea9cc4dc","modified_on":"2021-01-09T16:34:25.719Z","organisation_id":"eb0bd6f5-c3f5-44b2-b677-acd23cdde73c","type":"accounts","version":0},{"attributes":{"alternative_bank_account_names":null,"base_currency":"EUR","bic":"NLABNA01","country":"NL"},"created_on":"2021-01-09T16:34:25.725Z","id":"bf33e333-9605-4b4b-a0e5-3003ea9cc4dc","modified_on":"2021-01-09T16:34:25.725Z","organisation_id":"eb0bd6f5-c3f5-44b2-b677-acd23cdde73c","type":"accounts","version":0}],"links":{"first":"/v1/organisation/accounts?page%5Bnumber%5D=first\u0026page%5Bsize%5D=10","last":"/v1/organisation/accounts?page%5Bnumber%5D=last\u0026page%5Bsize%5D=10","self":"/v1/organisation/accounts?page%5Bnumber%5D=0\u0026page%5Bsize%5D=10"}}`))
		} else {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	})
```

Below are further handlers:

```
	testRouter.HandleFunc("/v1/organisation/accounts/ad27e265-9605-4b4b-a0e5-3003ea9cc4dc", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "application/vnd.api+json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"data":{"attributes":{"alternative_bank_account_names":null,"bank_id":"400300","bank_id_code":"GBDSC","base_currency":"GBP","bic":"NWBKGB22","country":"GB"},"created_on":"2021-01-09T16:39:15.962Z","id":"ad27e265-9605-4b4b-a0e5-3003ea9cc4dc","modified_on":"2021-01-09T16:39:15.962Z","organisation_id":"eb0bd6f5-c3f5-44b2-b677-acd23cdde73c","type":"accounts","version":0},"links":{"self":"/v1/organisation/accounts/ad27e265-9605-4b4b-a0e5-3003ea9cc4dc"}}`))
	})

	srv := httptest.NewServer(testRouter)
	defer srv.Close()

	var testConfig = Configuration{"", "", "", srv.URL + "/v1/organisation/accounts/"}
    ...
```

The other tests are ran within, using the below syntax:

```
	t.Run("Test to Create a Valid Uk Account", func(t *testing.T) {
		createdAcc, err := testConfig.CreateAccount(validUkAccount)
		// Validate that upon creation, create returns Account type
		if err != nil ||
			createdAcc.Type != validUkAccount.Type ||
			createdAcc.ID != validUkAccount.ID {
			t.Errorf("Error while creating account: %v", err)
		}
	})
```

# Resources

You can find here complete test case for [account creation service](../src/accountapi/service_test.go) and [account creation integration service](../src/accountapi/service_integration_test.go) 