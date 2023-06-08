package models

import "time"

type Order struct {
	Id             string    `json:"id"`
	TableId        string    `json:"table_id"`
	PaymentMethod  string    `json:"payment_method"`
	DeliveryStatus string    `json:"delivery_status"`
	DeliveryId     int       `json:"-"`
	PaymentId      int       `json:"-"`
	DateCreated    time.Time `json:"date_created"`
	DateUpdated    time.Time `json:"date_updated"`
}
