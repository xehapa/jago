package models

type User struct {
	UserId    int    `json:"userId"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Position  string `json:"position"`
	JobTitle  string `json:"jobTitle"`
	Email     string `json:"email"`
	Mobile    string `json:"mobile"`
	Links     struct {
		Self string `json:"self"`
	} `json:"links"`
	Deleted  bool `json:"deleted"`
	Inactive bool `json:"inactive"`
}
