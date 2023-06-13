package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"time"

	"github.com/xehapa/jago/models"
	"github.com/xehapa/jago/utils"
)

func GetPlacements(apiUrl, accessToken string) ([]models.Placement, error) {
	baseURL, err := url.Parse(apiUrl + "placements")
	if err != nil {
		return nil, fmt.Errorf("failed to parse base URL: %w", err)
	}

	queryParams := url.Values{}
	createdAt := time.Now().UTC().AddDate(-1, 0, 0).Format(time.RFC3339)
	queryParams.Set("limit", "1000")
	queryParams.Set("createdAt", fmt.Sprintf(">%s", createdAt))
	baseURL.RawQuery = queryParams.Encode()

	headers := make(map[string]string)
	headers["Authorization"] = "Bearer " + accessToken
	headers["Content-Type"] = "application/json"

	allPlacements := make([]models.Placement, 0)
	page := 1

	if err != nil {
		return nil, fmt.Errorf("failed to create temporary file: %w", err)
	}

	for {
		fmt.Println("API URL:", baseURL.String())
		responseBody, err := utils.NewHTTPClient().SendRequest("GET", baseURL.String(), nil, headers)

		if err != nil {
			return nil, fmt.Errorf("failed to send request: %w", err)
		}

		var response models.PlacementResponse
		err = json.Unmarshal(responseBody, &response)

		if err != nil {
			return nil, fmt.Errorf("failed to parse response body: %w", err)
		}

		allPlacements = append(allPlacements, response.Items...)

		if response.TotalCount <= 1000 || response.Links.Next == "" {
			break
		}

		fmt.Printf("Moving to Page %d\n", page)

		nextURL, err := url.Parse(response.Links.Next)

		if err != nil {
			return nil, fmt.Errorf("failed to parse next page URL: %w", err)
		}
		baseURL = nextURL

		fmt.Println("Sleeping for 5 seconds...")
		time.Sleep(5 * time.Second)

		page++
	}

	return allPlacements, nil
}

func FetchDetailedPlacement(link, accessToken string) models.PlacementDetailResponse {
	headers := make(map[string]string)
	headers["Authorization"] = "Bearer " + accessToken
	headers["Content-Type"] = "application/json"

	response, err := utils.NewHTTPClient().SendRequest("GET", link, nil, headers)

	if err != nil {
		log.Fatal("Failed to get placement detail: ", err)
	}

	var placementDetail models.PlacementDetailResponse
	err = json.Unmarshal(response, &placementDetail)

	if err != nil {
		log.Fatal("Failed to parse response: ", err)
	}

	return placementDetail
}
