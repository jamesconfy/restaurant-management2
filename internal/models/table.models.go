package models

import "time"

type Table struct {
	Id          string    `json:"id"`
	Seats       int       `json:"seats"`
	Number      int       `json:"number"`
	Booked      bool      `json:"booked"`
	DateCreated time.Time `json:"date_created"`
	DateUpdated time.Time `json:"date_updated"`
}
