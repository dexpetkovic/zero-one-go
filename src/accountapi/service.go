package accountapi

import (
	"encoding/json"
	"fmt"
)

// RequestBody wraps the Account object
type requestBody struct {
	Data Account `json:"data"`
}

// ResponseBody of single returned account
type responseBody struct {
	Data Account `json:"data"`
}

// ResponseBody of returned slice of accounts
type sliceResponseBody struct {
	Data []Account `json:"data"`
}

// Load configuration from simple config struct
var uri = config.AccountAPIUrl

// Create method instantiates an Account object and creates an account resource via API
func Create(
	Type string,
	ID string,
	OrganisationID string,
	Version int,
	CreatedOn string,
	ModifiedOn string,
	Country string,
	BaseCurrency string,
	AccountNumber string,
	BankID string,
	BankIDCode string,
	Bic string,
	Iban string,
	Name []string,
	AlternativeNames []string,
	AccountClassification string,
	JointAccount bool,
	AccountMatchingOptOut bool,
	SecondaryIdentification string,
	Switched bool,
	Status string) (Account, error) {

	account := Account{
		Type,
		ID,
		OrganisationID,
		Version,
		CreatedOn,
		ModifiedOn,
		countryAttributes{
			Country,
			BaseCurrency,
			AccountNumber,
			BankID,
			BankIDCode,
			Bic,
			Iban,
			Name,
			AlternativeNames,
			AccountClassification,
			JointAccount,
			AccountMatchingOptOut,
			SecondaryIdentification,
			Switched,
			Status,
		},
	}

	return account.Create()

}

// Create creates an account resource via API from given Account object
func (account Account) Create() (Account, error) {

	var createdAccount Account
	var createAccountResponseBody responseBody
	var createAccountResponse []byte
	var err error

	// Create Request body
	accountJSONReq, err := json.Marshal(requestBody{account})

	// Handle Marshalling errors
	if err != nil {
		err := fmt.Errorf("Marshalling error: %v", err)
		return createdAccount, err
	}

	// Create account resource
	createAccountResponse, err = doPost(uri, accountJSONReq)

	if err != nil {
		//log.Printf("Request failed: %v", err)
		return createdAccount, err
	}

	// If success, unmarshal the response into desired type and return to the caller
	err = json.Unmarshal(createAccountResponse, &createAccountResponseBody)

	if err != nil {
		//log.Printf("Unmarshalling response failed: %v", err)
		return createdAccount, err
	}

	createdAccount = createAccountResponseBody.Data

	return createdAccount, err
}

// FetchAccount fetches an account resource from Account API with given Account ID
func FetchAccount(accountID string) (Account, error) {
	var fetchedAccount Account
	var fetchAccountResponseBody responseBody
	var fetchAccountResponse []byte
	var err error
	var fetchURI = uri + accountID
	var queryParams = ""

	// Fetch account resource
	fetchAccountResponse, err = doGet(fetchURI, queryParams)

	if err != nil {
		//log.Printf("Request failed: %v", err)
		return fetchedAccount, err
	}

	// If success, unmarshal the response into desired type and return to the caller
	err = json.Unmarshal(fetchAccountResponse, &fetchAccountResponseBody)

	if err != nil {
		//log.Printf("Unmarshalling response failed: %v", err)
		return fetchedAccount, err
	}

	fetchedAccount = fetchAccountResponseBody.Data

	return fetchedAccount, err
}

// ListAccounts fetches paged account resources that match given filter
func ListAccounts(pageNumber int, pageSize int) ([]Account, error) {
	var listedAccounts []Account
	var listAccountsResponseBody sliceResponseBody
	var listAccountsResponse []byte
	var err error
	var queryParams = "?page[number]=" + fmt.Sprint(pageNumber) + "&page[size]=" + fmt.Sprint(pageSize)

	// List account resource
	listAccountsResponse, err = doGet(uri, queryParams)

	if err != nil {
		//log.Printf("Request failed: %v", err)
		return listedAccounts, err
	}
	// If success, unmarshal the response into desired type and return to the caller
	err = json.Unmarshal(listAccountsResponse, &listAccountsResponseBody)

	if err != nil {
		//log.Printf("Unmarshalling response failed: %v", err)
		return listedAccounts, err
	}

	listedAccounts = listAccountsResponseBody.Data

	return listedAccounts, err
}

// DeleteAccount deletes account resource with given ID
func DeleteAccount(accountID string, version int) error {
	var deleteURI = uri + accountID
	var queryParams = "?version=" + fmt.Sprint(version)

	// DELETE, when successful, does not return content.
	_, err := doDelete(deleteURI, queryParams)
	return err
}
