package models

type PlacementResponse struct {
	Items      []Placement `json:"items"`
	TotalCount int         `json:"totalCount"`
	Links      struct {
		First string `json:"first"`
		Prev  string `json:"prev"`
		Next  string `json:"next"`
		Last  string `json:"last"`
	} `json:"links"`
}

type Placement struct {
	PlacementID int    `json:"placementId"`
	JobTitle    string `json:"jobTitle"`
	Job         Job
	Candidate   Candidate
	Approved    bool
	ApprovedAt  string `json:"approvedAt"`
	Type        string
	Status      Status
	StartDate   string `json:"startDate"`
	EndDate     string `json:"endDate"`
	CreatedBy   User
	CreatedAt   string `json:"createdAt"`
	UpdatedBy   User
	UpdatedAt   string `json:"updatedAt"`
	Links       struct {
		Self string
	}
}

type EnhancedPlacement struct {
	PlacementId   int
	JobId         int
	JobOwnerName  string
	JobOwnerEmail string
	Link          string
	CreatedDate   string
}

type PlacementDetailResponse struct {
	PlacementID int `json:"placementId"`
	Job         Job
	Owner       User
	Recruiters  []User
}

type EnhancedPlacementDetail struct {
	PlacementID int    `json:"placementId"`
	JobID       int    `json:"jobId"`
	JobTitle    string `json:"jobTitle"`
	JobOwner    struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	} `json:"jobOwner"`
	Recruiters []struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}
}
