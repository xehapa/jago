package utils

import (
	"net/http"
	"time"
)

// HTTPClient interface represents an HTTP client
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client represents the HTTP client implementatio
type Client struct {
	client *http.Client
}

// NewHTTPClient creates a new instance of the HTTP client
func NewHTTPClient() HTTPClient {
	return &Client{
		client: &http.Client{
			Timeout: time.Second * 10,
		},
	}
}

// Do sends an HTTP request and returns the response
func (c *Client) Do(req *http.Request) (*http.Response, error) {
	return c.client.Do(req)
}
