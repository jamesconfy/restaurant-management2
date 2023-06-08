package service

import (
	"restaurant-management/internal/forms"
	"restaurant-management/internal/models"
	repo "restaurant-management/internal/repository"
	"restaurant-management/internal/se"

	"github.com/docker/distribution/uuid"
)

type OrderService interface {
	Add(req *forms.Order) (*models.Order, *se.ServiceError)
	Get(orderId string) (*models.Order, *se.ServiceError)
	GetAll() ([]*models.Order, *se.ServiceError)
	Edit(orderId string, req *forms.EditOrder) (*models.Order, *se.ServiceError)
	Delete(orderId string) *se.ServiceError
}

type orderSrv struct {
	repo      repo.OrderRepo
	tableRepo repo.TableRepo
}

// Delete implements OrderService.
func (o *orderSrv) Delete(orderId string) *se.ServiceError {
	if _, err := uuid.Parse(orderId); err != nil {
		return se.Internal(err, "invalid order id")
	}

	err := o.repo.Delete(orderId)
	if err != nil {
		return se.Internal(err)
	}

	return nil
}

// Edit implements OrderService.
func (o *orderSrv) Edit(orderId string, req *forms.EditOrder) (*models.Order, *se.ServiceError) {
	if _, err := uuid.Parse(orderId); err != nil {
		return nil, se.Internal(err, "invalid order id")
	}

	order, err := o.repo.Get(orderId)
	if err != nil {
		return nil, se.NotFoundOrInternal(err, "order not found")
	}

	order = o.getEdit(order, req)

	order, err = o.repo.Edit(orderId, order)
	if err != nil {
		return nil, se.Internal(err)
	}

	return order, nil
}

// Get implements OrderService.
func (o *orderSrv) Get(orderId string) (*models.Order, *se.ServiceError) {
	if _, err := uuid.Parse(orderId); err != nil {
		return nil, se.Internal(err, "invalid order id")
	}

	order, err := o.repo.Get(orderId)
	if err != nil {
		return nil, se.NotFoundOrInternal(err, "order not found")
	}

	return order, nil
}

// GetAll implements OrderService.
func (o *orderSrv) GetAll() ([]*models.Order, *se.ServiceError) {
	orders, err := o.repo.GetAll()
	if err != nil {
		return nil, se.Internal(err)
	}

	return orders, nil
}

// Add implements OrderService.
func (o *orderSrv) Add(req *forms.Order) (*models.Order, *se.ServiceError) {
	if err := Validator.validate(req); err != nil {
		return nil, se.Validating(err)
	}

	if _, err := uuid.Parse(req.TableId); err != nil {
		return nil, se.Internal(err, "invalid table id")
	}

	ok, er := o.tableRepo.TableExists(req.TableId)
	if !ok || er != nil {
		return nil, se.NotFoundOrInternal(er, "table not found")
	}

	var order models.Order

	order.TableId = req.TableId

	orde, err := o.repo.Add(&order)
	if err != nil {
		return nil, se.Internal(err)
	}

	return orde, nil
}

func NewOrderService(repo repo.OrderRepo, tableRepo repo.TableRepo) OrderService {
	return &orderSrv{repo: repo, tableRepo: tableRepo}
}

// Auxillary Functions
func (o *orderSrv) getEdit(order *models.Order, req *forms.EditOrder) *models.Order {
	if req.DeliveryId != 0 && req.DeliveryId != order.DeliveryId {
		order.DeliveryId = req.DeliveryId
	}

	if req.PaymentId != 0 && req.PaymentId != order.PaymentId {
		order.PaymentId = req.PaymentId
	}

	return order
}
