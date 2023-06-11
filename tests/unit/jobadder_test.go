package unit

import (
	"bytes"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/xehapa/jago/api"
)

type mockHTTPClient struct {
	mock.Mock
}

func (m *mockHTTPClient) Do(req *http.Request) (*http.Response, error) {
	args := m.Called(req)
	return args.Get(0).(*http.Response), args.Error(1)
}

func TestJobAdderClient_Do(t *testing.T) {
	// Create a mock HTTP client
	mockClient := new(mockHTTPClient)
	mockResponseBody := "Mock response"
	mockResponse := &http.Response{
		StatusCode: http.StatusOK,
		Body:       newReadCloser(mockResponseBody),
	}
	mockClient.On("Do", mock.Anything).Return(mockResponse, nil)

	// Create a JobAdderClient instance with the mock HTTP client
	client := api.NewJobAdderClient("apiKey", "apiSecret")
	client.HTTPClient = mockClient

	// Create an HTTP request
	req, err := http.NewRequest("GET", "http://example.com", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Call the Do method on the JobAdderClient
	resp, err := client.HTTPClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	// Read the response body
	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	body := buf.String()

	// Assert the response status code
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
	}

	// Assert the response body
	expectedBody := mockResponseBody
	if body != expectedBody {
		t.Errorf("Expected response body '%s', got '%s'", expectedBody, body)
	}

	// Assert that the mock HTTP client's Do method was called
	mockClient.AssertCalled(t, "Do", mock.Anything)
}

type readCloser struct {
	io.Reader
}

func (rc *readCloser) Close() error {
	return nil
}

func newReadCloser(s string) io.ReadCloser {
	return &readCloser{strings.NewReader(s)}
}
