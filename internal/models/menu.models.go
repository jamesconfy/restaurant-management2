package models

import "time"

type Menu struct {
	Id          string    `json:"id"`
	Name        string    `json:"name"`
	Category    string    `json:"category"`
	DateCreated time.Time `json:"date_created"`
	DateUpdated time.Time `json:"date_updated"`
}
