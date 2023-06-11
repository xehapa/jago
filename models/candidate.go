package models

type Candidate struct {
	CandidateID      int    `json:"candidateId"`
	FirstName        string `json:"firstName"`
	LastName         string `json:"lastName"`
	Email            string
	Phone            string
	Mobile           string
	ContactMethod    string `json:"contactMethod"`
	Salutation       string
	Unsubscribed     bool
	CandidateAddress Address
	Status           Status
	Rating           string
	Source           string
	Seeking          string
	DateOfBirth      string `json:"dateOfBirth"`
}
