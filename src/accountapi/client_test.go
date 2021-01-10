package accountapi

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetWithMockHTTPServer(t *testing.T) {
	var mockedResponse sliceResponseBody

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/vnd.api+json")
		fmt.Fprintln(w, `{"data":[{"attributes":{"alternative_bank_account_names":null,"base_currency":"EUR","bic":"NLABNA01","country":"NL"},"created_on":"2020-11-17T21:30:10.067Z","id":"bf33e333-9605-4b4b-a0e5-3003ea9cc4dc","modified_on":"2020-11-17T21:30:10.067Z","organisation_id":"eb0bd6f5-c3f5-44b2-b677-acd23cdde73c","type":"accounts","version":0}]}`)
	}))
	defer ts.Close()

	bMockedResponse, err := makeHTTPRequest(ts.URL, GET, 200, nil, "")

	if err != nil {
		t.Errorf("Error while getting account: %v ", err)
	}

	err = json.Unmarshal(bMockedResponse, &mockedResponse)

	if err != nil {
		t.Errorf("Error while parsing retrieved account: %v ", err)
	}

}

func TestBrokenGetWithMockHTTPServer(t *testing.T) {

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/vnd.api+json")
		fmt.Fprintln(w, `{"error_message": "You encountered an expected error during testing."}`)
	}))
	defer ts.Close()

	_, err := makeHTTPRequest(ts.URL, GET, 200, nil, "")

	if err == nil {
		t.Errorf("We expected failure in test case.")
	}

	if fmt.Sprint(err) != "You encountered an expected error during testing." {
		log.Fatal("Unexpected test failure")
	}

}

func TestGetWithBrokenMockHTTPServer1(t *testing.T) {

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/vnd.api+json")
		fmt.Fprintln(w, `{"invalid_json_key": "You encountered an unexpected error during testing."}`)
	}))
	defer ts.Close()

	_, err := makeHTTPRequest(ts.URL, GET, 200, nil, "")

	// There should be an error since server threw 500.
	if err == nil {
		log.Fatal("There should be an error since server threw 500.")
	}

	if fmt.Sprint(err) != "Request silently failed with status code 500" {
		log.Fatal("Unexpected test failure")
	}
}

func TestGetWithNonexistingIdWithMockHTTPServer(t *testing.T) {

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		w.Header().Set("Content-Type", "application/vnd.api+json")
		fmt.Fprintln(w, `{"error_message": "record XXX does not exist."}`)
	}))
	defer ts.Close()

	_, err := makeHTTPRequest(ts.URL, GET, 200, nil, "")

	// There should be an error.
	if err == nil {
		log.Fatal("There should be an error since server returned 404.")
	}

	if fmt.Sprint(err) != "record XXX does not exist." {
		log.Fatal("Unexpected test failure")
	}

}

func TestPostWithMockHTTPServer(t *testing.T) {
	var createdAccount Account
	var mockedResponse responseBody

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/vnd.api+json")
		fmt.Fprintln(w, `{"data":{"attributes":{"alternative_bank_account_names":null,"base_currency":"EUR","bic":"NLABNA01","country":"NL"},"created_on":"2020-11-17T21:30:10.067Z","id":"bf33e333-9605-4b4b-a0e5-3003ea9cc4dc","modified_on":"2020-11-17T21:30:10.067Z","organisation_id":"eb0bd6f5-c3f5-44b2-b677-acd23cdde73c","type":"accounts","version":0}}`)
	}))
	defer ts.Close()

	bMockedResponse, err := makeHTTPRequest(ts.URL, POST, 201, nil, "")

	if err != nil {
		t.Errorf("Error while creating account: %v ", err)
	}

	err = json.Unmarshal(bMockedResponse, &mockedResponse)

	if err != nil {
		t.Errorf("Error while parsing created account: %v ", err)
	}

	createdAccount = mockedResponse.Data

	if validNlAccount.ID != createdAccount.ID ||
		validNlAccount.OrganisationID != createdAccount.OrganisationID ||
		validNlAccount.Attributes.BankID != createdAccount.Attributes.BankID ||
		validNlAccount.Attributes.BankIDCode != createdAccount.Attributes.BankIDCode ||
		validNlAccount.Attributes.Bic != createdAccount.Attributes.Bic ||
		validNlAccount.Attributes.BaseCurrency != createdAccount.Attributes.BaseCurrency ||
		validNlAccount.Attributes.Country != createdAccount.Attributes.Country {
		t.Errorf("Created account:\n %v \ndoes not match expected account:\n %v ", createdAccount, validNlAccount)
	}

}

func TestListWithMockHTTPServer(t *testing.T) {
	var createdAccount Account
	var mockedResponse sliceResponseBody

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/vnd.api+json")
		fmt.Fprintln(w, `{"data":[{"attributes":{"alternative_bank_account_names":null,"base_currency":"EUR","bic":"NLABNA01","country":"NL"},"created_on":"2020-11-17T21:30:10.067Z","id":"bf33e333-9605-4b4b-a0e5-3003ea9cc4dc","modified_on":"2020-11-17T21:30:10.067Z","organisation_id":"eb0bd6f5-c3f5-44b2-b677-acd23cdde73c","type":"accounts","version":0}]}`)
	}))
	defer ts.Close()

	bMockedResponse, err := makeHTTPRequest(ts.URL, GET, 200, nil, "")

	if err != nil {
		t.Errorf("Error while listing accounts: %v ", err)
	}

	err = json.Unmarshal(bMockedResponse, &mockedResponse)

	if err != nil {
		t.Errorf("Error while parsing listed accounts: %v ", err)
	}

	createdAccount = mockedResponse.Data[0]

	if validNlAccount.ID != createdAccount.ID ||
		validNlAccount.OrganisationID != createdAccount.OrganisationID ||
		validNlAccount.Attributes.BankID != createdAccount.Attributes.BankID ||
		validNlAccount.Attributes.BankIDCode != createdAccount.Attributes.BankIDCode ||
		validNlAccount.Attributes.Bic != createdAccount.Attributes.Bic ||
		validNlAccount.Attributes.BaseCurrency != createdAccount.Attributes.BaseCurrency ||
		validNlAccount.Attributes.Country != createdAccount.Attributes.Country {
		t.Errorf("Listed account:\n %v \ndoes not match expected account:\n %v ", createdAccount, validNlAccount)
	}

}

func TestDeleteWithMockHTTPServer(t *testing.T) {

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
		w.Header().Set("Content-Type", "application/vnd.api+json")
	}))
	defer ts.Close()

	_, err := makeHTTPRequest(ts.URL, DELETE, 204, nil, "")

	if err != nil {
		log.Fatal("Delete failed")
	}

}
