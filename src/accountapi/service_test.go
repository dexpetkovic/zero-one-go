package accountapi

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

func TestCreateValidUkAccount(t *testing.T) {
	createdAcc, err := validUkAccount.Create()
	// Validate that upon creation, create returns Account type
	if err != nil ||
		createdAcc.Type != validUkAccount.Type ||
		createdAcc.ID != validUkAccount.ID {
		t.Errorf("Error while creating account: %v", err)
	}
}

func TestCreateValidNlAccount(t *testing.T) {
	createdAcc, err := validNlAccount.Create()
	// Validate that upon creation, create returns Account type
	if err != nil ||
		createdAcc.Type != validNlAccount.Type ||
		createdAcc.ID != validNlAccount.ID {
		t.Errorf("Error while creating account: %v", err)
	}
}

func TestCreateInvalidNlAccount(t *testing.T) {
	_, err := invalidNlAccount.Create()
	// Validate that upon creation, create returns Account type
	if err == nil {
		t.Errorf("Account creation should fail: %v", err)
	}
}

func TestFetchAccount(t *testing.T) {
	var accountID = "ad27e265-9605-4b4b-a0e5-3003ea9cc4dc"

	fetchedAccount, err := FetchAccount(accountID)
	if err != nil {
		t.Errorf("Error while fetching account id %v: %v", accountID, err)
	}

	if fetchedAccount.ID != accountID {
		t.Errorf("Fetched account does not have the same id: %v and %v.", fetchedAccount.ID, accountID)
	}
}

func TestListAccounts(t *testing.T) {
	listedAccounts, err := ListAccounts(0, 10)
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
}

func TestDeleteCreatedAccounts(t *testing.T) {
	listedAccounts, _ := ListAccounts(0, 10)

	for i := 0; i < len(listedAccounts); i++ {
		var removeAccount = listedAccounts[i].ID
		err := DeleteAccount(removeAccount, 0)

		if err != nil {
			t.Errorf("Error while deleting account: %v", err)
		}
	}
}
