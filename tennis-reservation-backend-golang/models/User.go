package models

type User struct {
	Id        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	IsActive  bool   `json:"isActive"`
	Role      string `json:"role"`
}
