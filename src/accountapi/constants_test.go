package accountapi

// Constants to use for testing
var validUkAccount = Account{
	Type:           "accounts",
	ID:             "ad27e265-9605-4b4b-a0e5-3003ea9cc4dc",
	OrganisationID: "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c",
	Attributes: countryAttributes{
		Country:      "GB",
		BaseCurrency: "GBP",
		BankID:       "400300",
		BankIDCode:   "GBDSC",
		Bic:          "NWBKGB22",
	},
}

var validNlAccount = Account{
	Type:           "accounts",
	ID:             "bf33e333-9605-4b4b-a0e5-3003ea9cc4dc",
	OrganisationID: "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c",
	Attributes: countryAttributes{
		Country:      "NL",
		BaseCurrency: "EUR",
		BankID:       "",
		BankIDCode:   "",
		Bic:          "NLABNA01",
	},
}

var invalidUkAccount = Account{
	Type:           "accounts",
	ID:             "myaccount",
	OrganisationID: "invalidval",
	Attributes: countryAttributes{
		Country:      "GB",
		BaseCurrency: "USD",
		BankID:       "400300",
		BankIDCode:   "GBDSC",
		Bic:          "NWBKGB22",
	},
}

var invalidNlAccount = Account{
	Type:           "accounts",
	ID:             "1234-abcd",
	OrganisationID: "org-id",
	Attributes: countryAttributes{
		Country:      "NL",
		BaseCurrency: "RSD",
		BankID:       "",
		BankIDCode:   "",
		Bic:          "ABNABIC",
	},
}
