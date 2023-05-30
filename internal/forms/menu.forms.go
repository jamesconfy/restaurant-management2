package forms

type Menu struct {
	Name     string `json:"name" validate:"required,min=1"`
	Category string `json:"category" validate:"required,min=1"`
}

type EditMenu struct {
	Name     string `json:"name"`
	Category string `json:"category"`
}
