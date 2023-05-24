package forms

type Table struct {
	Seats int `json:"seats" validate:"required,min=1"`
}
