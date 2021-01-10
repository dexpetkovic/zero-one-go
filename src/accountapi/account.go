package accountapi

import (
	"encoding/json"
)

// Account is object that holds account details. Optional attributes can be omitted.
type Account struct {
	Type           string            `json:"type"`
	ID             string            `json:"id"`
	OrganisationID string            `json:"organisation_id"`
	Version        int               `json:"version,omitempty"`
	CreatedOn      string            `json:"created_on,omitempty"`
	ModifiedOn     string            `json:"modified_on,omitempty"`
	Attributes     countryAttributes `json:"attributes"`
}

type countryAttributes struct {
	Country                 string   `json:"country"`
	BaseCurrency            string   `json:"base_currency,omitempty"`
	AccountNumber           string   `json:"account_number,omitempty"`
	BankID                  string   `json:"bank_id,omitempty"`
	BankIDCode              string   `json:"bank_id_code,omitempty"`
	Bic                     string   `json:"bic,omitempty"`
	Iban                    string   `json:"iban,omitempty"`
	Name                    []string `json:"name,omitempty"`
	AlternativeNames        []string `json:"alternative_names,omitempty"`
	AccountClassification   string   `json:"account_classification,omitempty"`
	JointAccount            bool     `json:"joint_account,omitempty"`
	AccountMatchingOptOut   bool     `json:"account_matching_opt_out,omitempty"`
	SecondaryIdentification string   `json:"secondary_identification,omitempty"`
	Switched                bool     `json:"switched,omitempty"`
	Status                  string   `json:"status,omitempty"`
}

// Marshal input object (of Account type) to JSON
// TODO: Apply inheritance to accept any inherited type
func (input Account) toJSON() (string, error) {
	JSONBytes, err := json.Marshal(input)
	if err != nil {
		//log.Printf("Could not marshal input to JSON string, returning empty string: %v", err)
		return "", err
	}
	return string(JSONBytes), err
}
