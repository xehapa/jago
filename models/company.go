package models

type Company struct {
	CompanyID int `json:"companyId"`
	Name      string
	Status    Status
	Owner     User
}
