package runner

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/xehapa/jago/api"
)

type AccessTokenProvider interface {
	GetAccessToken(refreshToken string) (*RefreshTokenResponse, error)
}

type RefreshTokenResponse struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	ApiUrl       string `json:"api"`
	RefreshToken string `json:"refresh_token"`
}

func RunMain() {
	// Prompt user for refresh token
	var refreshToken string
	fmt.Print("Enter Refresh Token: ")
	refreshToken, _ = readString()
	client := api.NewJobAdderClient()

	refreshTokenResponse, err := client.GetAccessToken(refreshToken)

	if err != nil {
		log.Fatalf("failed to get access token: %v", err)
	}

	jsonResponse, err := json.MarshalIndent(refreshTokenResponse, "", "  ")
	if err != nil {
		log.Fatalf("failed to marshal response to JSON: %v", err)
	}

	fmt.Println("Response:")
	fmt.Println(string(jsonResponse))

	client.ApiUrl = refreshTokenResponse.ApiUrl
	client.AccessToken = refreshTokenResponse.AccessToken

	placements, err := client.GetPlacements()

	if err != nil {
		log.Fatalf("failed to get placements: %v", err)
	}

	itemsJSON, err := json.MarshalIndent(placements, "", "  ")

	if err != nil {
		log.Fatal("failed to marshal Items as JSON: %w", err)
	}

	fmt.Println(string(itemsJSON))
}

func readString() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	str, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(str), nil
}
