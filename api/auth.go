// auth.go

package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/xehapa/jago/models"
	"github.com/xehapa/jago/utils"
)

func Auth(clientID, clientSecret string) {
	authURL := models.AuthURL + "/authorize"

	// Set the authorization request parameters
	params := url.Values{}
	params.Set("response_type", "code")
	params.Set("client_id", clientID)
	params.Set("scope", "read")
	params.Set("redirect_uri", "http://localhost:8000/jobadder/connect")

	// Construct the authorization request URL
	authURL += "?" + params.Encode()

	fmt.Println("Please open the following URL in your browser and authorize the application:")
	fmt.Println(authURL)

	fmt.Print("Enter the authorization code: ")
	var authCode string
	fmt.Scanln(&authCode)
}

func ExchangeCodeForToken(code, clientID, clientSecret, redirectURI string) (string, error) {
	tokenURL := models.AuthURL + "/token"

	// Set the token request parameters
	params := url.Values{}
	params.Set("grant_type", "authorization_code")
	params.Set("code", code)
	params.Set("client_id", clientID)
	params.Set("client_secret", clientSecret)
	params.Set("redirect_uri", redirectURI)

	// Construct the token request URL
	reqURL := fmt.Sprintf("%s?%s", tokenURL, params.Encode())

	// Create an HTTP client
	client := &http.Client{}

	// Send the token request
	resp, err := client.Get(reqURL)
	if err != nil {
		return "", fmt.Errorf("failed to send token request: %v", err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read token response: %v", err)
	}

	// Check if the response is an error
	if resp.StatusCode != http.StatusOK {
		var errorResp struct {
			Error            string `json:"error"`
			ErrorDescription string `json:"error_description"`
		}
		if err := json.Unmarshal(body, &errorResp); err != nil {
			return "", fmt.Errorf("failed to parse error response: %v", err)
		}
		return "", fmt.Errorf("token request failed: %s", errorResp.ErrorDescription)
	}

	// Parse the token response
	var tokenResp struct {
		AccessToken string `json:"access_token"`
		TokenType   string `json:"token_type"`
		ExpiresIn   int    `json:"expires_in"`
	}
	if err := json.Unmarshal(body, &tokenResp); err != nil {
		return "", fmt.Errorf("failed to parse token response: %v", err)
	}

	return tokenResp.AccessToken, nil
}

func (j *JobAdderClient) ExchangeRefreshToken(refreshToken string) (models.RefreshTokenResponse, error) {
	tokenURL := models.AuthURL + "/token"

	headers := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	// Create the request body
	body := fmt.Sprintf("client_id=%s&client_secret=%s&grant_type=refresh_token&refresh_token=%s",
		j.ClientID, j.ClientSecret, refreshToken)

	// Send the request using the HTTP client
	resp, err := utils.NewHTTPClient().SendRequest(http.MethodPost, tokenURL, []byte(body), headers)

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
