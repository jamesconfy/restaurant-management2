package models

import "time"

type Order struct {
	Id          string    `json:"id"`
	TableId     string    `json:"table_id"`
	DeliveryId  int       `json:"delivery_id"`
	PaymentId   int       `json:"payment_id"`
	DateCreated time.Time `json:"date_created"`
	DateUpdated time.Time `json:"date_updated"`
}
