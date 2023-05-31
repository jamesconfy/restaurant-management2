package forms

type Food struct {
	Name  string  `json:"name" validate:"required,min=1"`
	Price float64 `json:"price" validate:"required,min=0"`
	Image string  `json:"image"`
}

type EditFood struct {
	Name  string  `json:"name" validate:"required,min=1"`
	Price float64 `json:"price" validate:"required,min=0"`
	Image string  `json:"image"`
}
