package handlers

type OrderItemsHandler interface{}

type orderItemsHandler struct {
}

func NewOrderItemHandler() OrderItemsHandler {
	return &orderItemsHandler{}
}
