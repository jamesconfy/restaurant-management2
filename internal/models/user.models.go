package models

import "time"

type User struct {
	Id          string    `json:"user_id"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	PhoneNumber string    `json:"phone_number"`
	Email       string    `json:"email"`
	Role        string    `json:"role"`
	Password    string    `json:"password"`
	DateCreated time.Time `json:"date_created"`
	DateUpdated time.Time `json:"date_updated"`
}

// type Role string

// func (r *Role) Access() bool {
// 	return *r != "USER"
// }
