package accountapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

// Creating custom Transport to prevent connection hanging due to
// default values being set to infinite
var transport = &http.Transport{
	// Limit the maximum number of idle (keep-alive) connections
	MaxIdleConns: 10,
	// Limit the maximum amount of time an idle
	// (keep-alive) connection will remain idle before closing
	IdleConnTimeout: 30 * time.Second,
	// Limit the amount of time to wait for a server's response headers
	// after fully writing the request
	ResponseHeaderTimeout: 10 * time.Second,
	// Limit the maximum amount of time to wait for a TLS handshake.
	TLSHandshakeTimeout: 10 * time.Second,
}

// Create global http client for shared use.
// Timeout is set to 10 seconds to prevent deadlocks
var netClient = &http.Client{
	Transport: transport,
	Timeout:   time.Second * 10,
}

// HTTPMethod encapsulates HTTP verbs
type HTTPMethod string

// Const that lists possible HTTP methods
const (
	POST   HTTPMethod = "POST"
	GET               = "GET"
	DELETE            = "DELETE"
)

func doPost(baseURL string, bRequestBody []byte) ([]byte, error) {
	return makeHTTPRequest(baseURL, POST, 201, bRequestBody, "")
}

func doGet(baseURL string, queryParams string) ([]byte, error) {
	return makeHTTPRequest(baseURL, GET, 200, nil, queryParams)
}

func doDelete(baseURL string, queryParams string) ([]byte, error) {
	return makeHTTPRequest(baseURL, DELETE, 204, nil, queryParams)
}

func makeHTTPRequest(baseURL string, method HTTPMethod, successStatusCode int, bRequestBody []byte, queryParams string) ([]byte, error) {
	var resp *http.Response
	var err error
	var bResponseBody []byte

	// Prepare request body in case it's a Post.
	reqBody := strings.NewReader(string(bRequestBody))

	// Append query parameters to the baseURL
	if queryParams != "" {
		baseURL = baseURL + queryParams
	}

	switch method {
	case POST:
		resp, err = netClient.Post(baseURL, "application/vnd.api+json", reqBody)
	case GET:
		resp, err = netClient.Get(baseURL)
	case DELETE:
		req, err := http.NewRequest("DELETE", baseURL, nil)
		if err != nil {
			return bResponseBody, err
		}
		resp, err = netClient.Do(req)
	}

	if err != nil {
		//log.Printf("Failed to access AccountAPI endpoint: %v", err)
		return bResponseBody, err
	}

	// Defer closing the stream
	defer resp.Body.Close()
	// Try to read response body (if not nil) to see the response
	bResponseBody, err = ioutil.ReadAll(resp.Body)

	if err != nil {
		//log.Printf("Failed to parse AccountAPI response: %v", err)
		return bResponseBody, err
	}

	if resp.StatusCode != successStatusCode {
		// If not success, then get verbose response error from the body
		var resErr responseErr
		// With side effect, probably not ideal
		err = json.Unmarshal(bResponseBody, &resErr)

		// If parsing of error response succeeds, and is non-empty, we return that.
		// Else, it was a silent failure, assuming that API will always return error in same format.
		if err == nil {
			if resErr.ErrorMessage != "" {
				err = fmt.Errorf(resErr.ErrorMessage)
			} else {
				err = fmt.Errorf("Request silently failed with status code %v", resp.StatusCode)
			}
		}
	}

	return bResponseBody, err
}
