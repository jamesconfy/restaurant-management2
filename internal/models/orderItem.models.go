package models

import "time"

type OrderItem struct {
	Id          string    `json:"id"`
	Quantity    int       `json:"quantity"`
	OrderId     string    `json:"order_id"`
	FoodId      string    `json:"food_id"`
	DateCreated time.Time `json:"date_created"`
	DateUpdated time.Time `json:"date_updated"`
}
