package utils

import (
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
