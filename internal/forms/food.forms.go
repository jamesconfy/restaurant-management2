package forms

type Food struct {
	Name   string  `json:"name" validate:"required,min=1"`
	Price  float64 `json:"price" validate:"required,min=0"`
	MenuId string  `json:"menu_id" validate:"required"`
	Image  string  `json:"image"`
}

type EditFood struct {
	Name   string  `json:"name"`
	Price  float64 `json:"price"`
	Image  string  `json:"image"`
	MenuId string  `json:"menu_id"`
}
