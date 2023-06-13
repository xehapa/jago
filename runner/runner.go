package runner

import (
	"bufio"
	"fmt"
	"github.com/xehapa/jago/utils"
	"log"
	"os"
	"strings"
	"time"

	"github.com/xehapa/jago/api"
	"github.com/xehapa/jago/models"
)

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

	var placements []models.Placement
	placements = client.GetPlacements()

	var enhancedPlacements []models.EnhancedPlacement
	enhancedPlacements = utils.EnhancePlacement(placements)

	client.GetPlacementDetail(enhancedPlacements, refreshToken)

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
