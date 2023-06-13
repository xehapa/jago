package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/xehapa/jago/models"
)

func EnhancePlacement(rawPlacements []models.Placement) []models.EnhancedPlacement {
	enhancedPlacements := make([]models.EnhancedPlacement, 0)

	for _, placement := range rawPlacements {
		enhancedPlacement := models.EnhancedPlacement{
			PlacementId:   placement.PlacementID,
			JobId:         placement.Job.JobID,
			JobOwnerName:  fmt.Sprintf("%s %s", placement.Job.Owner.FirstName, placement.Job.Owner.LastName),
			JobOwnerEmail: placement.Job.Owner.Email,
			Link:          placement.Links.Self,
			CreatedDate:   placement.CreatedAt,
		}

		if !placement.Job.Owner.Deleted {
			enhancedPlacements = append(enhancedPlacements, enhancedPlacement)
		}
	}

	return enhancedPlacements
}

func SavePlacementDetailToFile(data []models.EnhancedPlacementDetail, refreshToken string) {
	jsonData, err := json.MarshalIndent(data, "", "    ")

	if err != nil {
		log.Fatal("Failed to marshal placement detail: ", err)
	}

	if _, err := os.Stat("data"); os.IsNotExist(err) {
		err := os.Mkdir("data", 0755)

		if err != nil {
			log.Fatal("Failed to create folder: ", err)
		}
	}

	fileName := fmt.Sprintf("data/placement_%s.json", refreshToken)

	// Check if file exists and delete it then recreate it
	if _, err := os.Stat(fileName); err == nil {
		err := os.Remove(fileName)

		if err != nil {
			log.Fatal("Failed to remove file: ", err)
		}

		fmt.Printf("Removed existing file: %s\n", fileName)

		_, err = os.Create(fileName)

		if err != nil {
			log.Fatal("Failed to create file: ", err)
		}

		fmt.Printf("Created new file: %s\n", fileName)
	}

	// write data to file and close it after done
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0755)

	if err != nil {
		log.Fatal("Failed to open file: ", err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal("Failed to close file: ", err)
		}
	}(file)

	_, err = file.Write(jsonData)

	if err != nil {
		log.Fatal("Failed to write data to file: ", err)
	}

	fmt.Printf("Saved data to file: %s\n", fileName)
}
