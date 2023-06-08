package forms

type Order struct {
	TableId string `json:"table_id" validate:"required,min=1"`
}

type EditOrder struct {
	DeliveryId int `json:"delivery_id"`
	PaymentId  int `json:"payment_id"`
}
