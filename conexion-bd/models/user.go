package models

type User struct {
	UserName  string `json:"user_name"`
	Email     string `json:"email"`
	FirstName string `json:"firstnames"`
	LastName  string `json:"lastnames"`
	Phone     string `json:"mobilenumber"`
}
