package accountapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"sort"
	"testing"
)

func TestAccountCreationService(t *testing.T) {
	testRouter := http.NewServeMux()
	var testAccountResponseBody responseBody

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

	t.Run("Test to Create a Valid Uk Account", func(t *testing.T) {
		createdAcc, err := testConfig.CreateAccount(validUkAccount)
		// Validate that upon creation, create returns Account type
		if err != nil ||
			createdAcc.Type != validUkAccount.Type ||
			createdAcc.ID != validUkAccount.ID {
			t.Errorf("Error while creating account: %v", err)
		}
	})

	t.Run("Test to Create a Valid Nl Account", func(t *testing.T) {
		createdAcc, err := testConfig.CreateAccount(validNlAccount)
		// Validate that upon creation, create returns Account type
		if err != nil ||
			createdAcc.Type != validNlAccount.Type ||
			createdAcc.ID != validNlAccount.ID {
			t.Errorf("Error while creating account: %v", err)
		}
	})

	t.Run("Test to Create an Invalid Nl Account", func(t *testing.T) {
		_, err := testConfig.CreateAccount(invalidNlAccount)
		// Validate that upon creation, create returns Account type
		if err == nil {
			t.Errorf("Account creation should fail: %v", err)
		}
	})

	t.Run("Test to List Accounts", func(t *testing.T) {
		listedAccounts, err := testConfig.ListAccounts(0, 10)
		if err != nil {
			t.Errorf("Error while fetching accounts: %v", err)
		}

		expected := []Account{validNlAccount, validUkAccount}

		// We can not use reflect.DeepEqual as API returns created/modified timestamps which,
		// in this app and its tests, change upon every test run.
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
	})

	t.Run("Test Fetching Account", func(t *testing.T) {
		var accountID = "ad27e265-9605-4b4b-a0e5-3003ea9cc4dc"

		fetchedAccount, err := testConfig.FetchAccount(accountID)
		if err != nil {
			t.Errorf("Error while fetching account id %v: %v", accountID, err)
		}

		if fetchedAccount.ID != accountID {
			t.Errorf("Fetched account does not have the same id: %v and %v.", fetchedAccount.ID, accountID)
		}
	})
}
