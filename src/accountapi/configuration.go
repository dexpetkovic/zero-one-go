package accountapi

import "os"

// Configuration struct holds application configuration
type Configuration struct {
	AccountAPIProtocol string
	AccountAPISocket   string
	AccountAPIUri      string
	AccountAPIUrl      string
}

// For easier running tests from IDE and CLI, we set defaults
func setDefaults(val string, defaultVal string) string {
	if val == "" {
		return defaultVal
	}
	return val
}

var accountAPIProtocol = setDefaults(os.Getenv("AccountAPIProtocol"), "http://")
var accountAPISocket = setDefaults(os.Getenv("AccountAPISocket"), "localhost:8080")
var accountAPIUri = setDefaults(os.Getenv("AccountAPIUri"), "/v1/organisation/accounts/")
var accountAPIUrl = accountAPIProtocol + accountAPISocket + accountAPIUri

var config = Configuration{accountAPIUri, accountAPISocket, accountAPIUri, accountAPIUrl}
