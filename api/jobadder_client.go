package api

import (
	"fmt"
	"log"

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

func (j *JobAdderClient) GetPlacements() []models.Placement {
	allPlacements, err := GetPlacements(j.ApiUrl, j.AccessToken)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Total Placements Fetched: %d\n", len(allPlacements))

	return allPlacements
}

func (j *JobAdderClient) GetPlacementDetail(placements []models.EnhancedPlacement, refreshToken string) models.EnhancedPlacementDetail {
	i := 0
	var data models.EnhancedPlacementDetail
	var lists []models.EnhancedPlacementDetail

	for _, placement := range placements {
		placementDetail := FetchDetailedPlacement(placement.Link, j.AccessToken)

		data = models.EnhancedPlacementDetail{
			PlacementID: placementDetail.PlacementID,
			JobID:       placementDetail.Job.JobID,
			JobTitle:    placementDetail.Job.JobTitle,
			JobOwner: struct {
				Name  string `json:"name"`
				Email string `json:"email"`
			}{
				Name:  placementDetail.Job.Owner.FirstName + " " + placementDetail.Job.Owner.LastName,
				Email: placementDetail.Job.Owner.Email,
			},
		}

		var recruiters = make([]struct {
			Name  string `json:"name"`
			Email string `json:"email"`
		}, 0)

		for _, recruiter := range placementDetail.Recruiters {
			recruiters = append(recruiters, struct {
				Name  string `json:"name"`
				Email string `json:"email"`
			}{
				Name:  recruiter.FirstName + " " + recruiter.LastName,
				Email: recruiter.Email,
			})
		}

		//data.Recruiters = recruiters

		i++

		lists = append(lists, data)

		fmt.Printf("%d. %d: %s\n", i, placementDetail.PlacementID, placementDetail.Job.JobTitle)
	}

	utils.SavePlacementDetailToFile(lists, refreshToken)

	return data
}
