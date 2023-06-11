package utils

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

// HTTPClient interface represents an HTTP client
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client represents the HTTP client implementation
type Client struct {
	HTTPClient HTTPClient
}

// NewHTTPClient creates a new instance of the HTTP client
func NewHTTPClient() *Client {
	return &Client{
		HTTPClient: &http.Client{
			Timeout: time.Second * 10,
		},
	}
}

// Do sends an HTTP request and returns the response
func (c *Client) Do(req *http.Request) (*http.Response, error) {
	return c.HTTPClient.Do(req)
}

func (c *Client) SendRequest(method, url string, body []byte, headers map[string]string) ([]byte, error) {
	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))

	if err != nil {
		log.Fatal(err)
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	resp, err := c.HTTPClient.Do(req)

	if err != nil {
		log.Fatalf("Request failed: %v", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		errMsg := fmt.Sprintf("Request failed with status code %d", resp.StatusCode)
		log.Fatal(errMsg)
	}

	var responseBody bytes.Buffer
	_, err = io.Copy(&responseBody, resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	return responseBody.Bytes(), nil
}
