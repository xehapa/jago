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
	Links         string
	CreatedDate   string
}
