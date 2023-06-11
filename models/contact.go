package models

type Contact struct {
	ContactID    int    `json:"contactId"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	Position     string
	Salutation   string
	Unsubscribed bool
	Email        string
	Phone        string
	Mobile       string
	Status       Status
	Owner        User
}
