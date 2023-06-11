package models

type Job struct {
	JobID    int `json:"jobId"`
	JobTitle string
	Company  Company
	Contact  Contact
	Status   Status
	Source   string
	Owner    User
}
