package runner

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

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
	timeStart := time.Now()

	refreshTokenResponse, err := client.GetAccessToken(refreshToken)

	if err != nil {
		log.Fatalf("failed to get access token: %v", err)
	}

	client.ApiUrl = refreshTokenResponse.ApiUrl
	client.AccessToken = refreshTokenResponse.AccessToken

	_, err = client.GetPlacements()

	if err != nil {
		log.Fatalf("failed to get placements: %v", err)
	}

	elapsedTime := time.Since(timeStart).Round(time.Second)

	fmt.Printf("Elapsed time: %s\n", elapsedTime)
}

func readString() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	str, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(str), nil
}
