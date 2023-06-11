package models

type User struct {
	UserID    int    `json:"userId"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Position  string
	Email     string
	Phone     string
	Mobile    string
	Inactive  bool
	Deleted   bool
}
