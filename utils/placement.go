package utils

import (
	"fmt"

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
			Links:         placement.Links.Self,
			CreatedDate:   placement.CreatedAt,
		}

		if !placement.Job.Owner.Deleted {
			enhancedPlacements = append(enhancedPlacements, enhancedPlacement)
		}
	}

	return enhancedPlacements
}
