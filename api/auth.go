package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/xehapa/jago/config"
	"github.com/xehapa/jago/models"
	"github.com/xehapa/jago/utils"
)

func (j *JobAdderClient) ExchangeRefreshToken(refreshToken string) (models.RefreshTokenResponse, error) {
	config := config.NewConfig(nil)

	headers := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	// Create the request body
	body := fmt.Sprintf("client_id=%s&client_secret=%s&grant_type=refresh_token&refresh_token=%s",
		j.ClientID, j.ClientSecret, refreshToken)

	// Send the request using the HTTP client
	resp, err := utils.NewHTTPClient().SendRequest(http.MethodPost, config.AuthUrl, []byte(body), headers)

	if err != nil {
		fmt.Println(resp)
		log.Fatal(err)
	}

	// Parse the response body
	var tokenResp models.RefreshTokenResponse
	err = json.Unmarshal(resp, &tokenResp)

	if err != nil {
		log.Fatal(fmt.Errorf("failed to parse token response: %v", err))
	}

	return tokenResp, nil
}
