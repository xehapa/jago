package unit

import (
	"net/http"
	"testing"
	"time"

	"github.com/xehap/jago/utils"
)

type MockHTTPClient struct {
	response *http.Response
	err      error
}

func (m *MockHTTPClient) Do(req *http.Request) (*http.Response, error) {
	return m.response, m.err
}

func TestClientDo(t *testing.T) {
	// Create a mock HTTP response
	response := &http.Response{
		StatusCode: http.StatusOK,
		Body:       nil,
	}

	// Create a client instance with the mock HTTP client
	client := &utils.Client{
		HTTPClient: &MockHTTPClient{
			response: response,
			err:      nil,
		},
	}

	// Create a mock HTTP request
	req, _ := http.NewRequest(http.MethodGet, "http://example.com", nil)

	// Perform the request
	resp, err := client.Do(req)

	// Assert the response and error
	if err != nil {
		t.Fatalf("Expected no error, but got: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, but got: %d", http.StatusOK, resp.StatusCode)
	}
}

func TestNewHTTPClient(t *testing.T) {
	// Create a new instance of the HTTP client
	client := utils.NewHTTPClient()

	// Assert the client type
	_, ok := client.HTTPClient.(*http.Client)
	if !ok {
		t.Error("Expected the client to be an instance of *http.Client")
	}

	// Assert the client timeout value
	timeout := client.HTTPClient.(*http.Client).Timeout
	expectedTimeout := time.Second * 10
	if timeout != expectedTimeout {
		t.Errorf("Expected client timeout to be %v, but got: %v", expectedTimeout, timeout)
	}
}
