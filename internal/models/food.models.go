package models

import "time"

type Food struct {
	Id          string    `json:"id"`
	Name        string    `json:"name"`
	Price       float64   `json:"price"`
	Image       string    `json:"image"`
	MenuId      string    `json:"menu_id"`
	DateCreated time.Time `json:"date_created"`
	DateUpdated time.Time `json:"date_updated"`
}

type MenuFood struct {
	Menu   *Menu   `json:"menu"`
	Foods  []*Food `json:"foods"`
	MenuId string  `json:"-"`
}
