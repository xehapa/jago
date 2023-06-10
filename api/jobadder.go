package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/xehap/jago/models"
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

func (c *JobAdderClient) GetJobs() ([]models.Job, error) {
	url := "https://api.jobadder.com/v2/jobs"
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(c.apiKey, c.apiSecret)

	resp, err := c.httpClient.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	var jobs []models.Job
	err = json.Unmarshal(body, &jobs)

	if err != nil {
		return nil, err
	}

	return jobs, nil
}
