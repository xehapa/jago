package api

import (
	"encoding/json"
	"fmt"
	"net/url"
	"time"

	"github.com/xehapa/jago/models"
	"github.com/xehapa/jago/utils"
)

type JobAdderClient struct {
	ApiUrl      string
	AccessToken string
	HTTPClient  utils.HTTPClient
}

func NewJobAdderClient() *JobAdderClient {
	return &JobAdderClient{
		HTTPClient: utils.NewHTTPClient(),
	}
}

func (c *JobAdderClient) SetHTTPClient(httpClient utils.HTTPClient) {
	c.HTTPClient = httpClient
}

func (j *JobAdderClient) GetAccessToken(refreshToken string) (*models.RefreshTokenResponse, error) {
	if refreshToken != "" {
		resp, err := j.ExchangeRefreshToken(refreshToken)
		return &resp, err
	}

	return nil, nil
}

func (j *JobAdderClient) GetPlacements() ([]models.Placement, error) {
	baseURL, err := url.Parse(j.ApiUrl + "placements")
	if err != nil {
		return nil, fmt.Errorf("failed to parse base URL: %w", err)
	}

	queryParams := url.Values{}
	createdAt := time.Now().UTC().AddDate(-1, 0, 0).Format(time.RFC3339)
	queryParams.Set("limit", "1000")
	queryParams.Set("createdAt>", createdAt)
	baseURL.RawQuery = queryParams.Encode()

	headers := make(map[string]string)
	headers["Authorization"] = "Bearer " + j.AccessToken
	headers["Content-Type"] = "application/json"

	responseBody, err := utils.NewHTTPClient().SendRequest("GET", baseURL.String(), nil, headers)

	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}

	var response struct {
		Items []models.Placement `json:"items"`
	}
	err = json.Unmarshal(responseBody, &response)

	if err != nil {
		return nil, fmt.Errorf("failed to parse response body: %w", err)
	}

	return response.Items, nil
}
