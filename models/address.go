package models

type Address struct {
	Street      *string `json:"street"`
	City        string
	State       string
	PostalCode  string `json:"postalCode"`
	Country     string
	CountryCode string `json:"countryCode"`
}
