package forms

type Table struct {
	Seats int `json:"seats" validate:"required,min=1"`
}

type EditTable struct {
	Seats  int  `json:"seats"`
	Booked bool `json:"booked"`
}
