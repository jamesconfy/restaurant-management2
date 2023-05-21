package models

import "time"

type Order struct {
	Id          string    `json:"id"`
	TableId     string    `json:"table_id"`
	DateCreated time.Time `json:"date_created"`
	DateUpdated time.Time `json:"date_updated"`
}
