package models

type Status struct {
	StatusID int `json:"statusId"`
	Name     string
	Active   bool
	Default  bool
}
