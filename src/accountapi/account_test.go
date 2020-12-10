package accountapi

import (
	"testing"
)

func TestNlAccountToJson(t *testing.T) {

	accountJSON, err := validNlAccount.toJSON()
	expected := `{"type":"accounts","id":"bf33e333-9605-4b4b-a0e5-3003ea9cc4dc","organisation_id":"eb0bd6f5-c3f5-44b2-b677-acd23cdde73c","attributes":{"country":"NL","base_currency":"EUR","bic":"NLABNA01"}}`
	if err != nil {
		t.Errorf("Error casting Account to Json: %v", err)
	}

	if accountJSON != expected {
		t.Errorf("Casted account does not match: %v", accountJSON)
	}
}
