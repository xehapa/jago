package api

import (
	"github.com/xehap/jago/utils"
)

type JobAdderClient struct {
	apiKey     string
	apiSecret  string
	httpClient utils.HTTPClient
}

func NewJobadderClient(apiKey, apiSecret string) *JobAdderClient {
	return &JobAdderClient{
		apiKey:     apiKey,
		apiSecret:  apiSecret,
		httpClient: utils.NewHTTPClient(),
	}
}
