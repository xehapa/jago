package api

import (
	"github.com/xehap/jago/utils"
)

type JobAdderClient struct {
	apiKey     string
	apiSecret  string
	HTTPClient utils.HTTPClient // Add the HTTPClient field
}

func NewJobAdderClient(apiKey, apiSecret string) *JobAdderClient {
	return &JobAdderClient{
		apiKey:     apiKey,
		apiSecret:  apiSecret,
		HTTPClient: utils.NewHTTPClient(), // Initialize the HTTPClient field
	}
}

func (c *JobAdderClient) SetHTTPClient(httpClient utils.HTTPClient) {
	c.HTTPClient = httpClient
}
