package runner

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/xehapa/jago/api"
	"github.com/xehapa/jago/config"
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
	config := config.NewConfig(nil)
	client := api.NewJobAdderClient(config.ClientId, config.ClientSecret)

	// Get access token
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
}

func readString() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	str, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(str), nil
}
