package api

import (
	"github.com/xehapa/jago/models"
	"github.com/xehapa/jago/utils"
)

type JobAdderClient struct {
	ClientID     string
	ClientSecret string
	HTTPClient   utils.HTTPClient // Add the HTTPClient field
}

func NewJobAdderClient(clientId, clientSecret string) *JobAdderClient {
	return &JobAdderClient{
		ClientID:     clientId,
		ClientSecret: clientSecret,
		HTTPClient:   utils.NewHTTPClient(), // Initialize the HTTPClient field
	}
}

func (c *JobAdderClient) SetHTTPClient(httpClient utils.HTTPClient) {
	c.HTTPClient = httpClient
}

// GetAccessToken retrieves the access token using the client ID and secret
func (j *JobAdderClient) GetAccessToken(refreshToken string) (*models.RefreshTokenResponse, error) {
	if refreshToken != "" {
		resp, err := j.ExchangeRefreshToken(refreshToken)
		return &resp, err
	}

	return nil, nil
}
