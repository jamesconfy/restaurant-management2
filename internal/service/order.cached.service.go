package service

import (
	"fmt"
	"restaurant-management/internal/forms"
	"restaurant-management/internal/models"
	repo "restaurant-management/internal/repository"
	"restaurant-management/internal/se"
	"restaurant-management/utils"
)

type cachedOrderService struct {
	orderSrv OrderService
	cache    repo.Cache
}

// Add implements OrderService.
func (co *cachedOrderService) Add(req *forms.Order) (order *models.Order, err *se.ServiceError) {
	order, err = co.orderSrv.Add(req)
	if err == nil {
		co.cache.DeleteByTag(utils.OrdersTag)
	}

	return
}

// Delete implements OrderService.
func (co *cachedOrderService) Delete(orderId string) (err *se.ServiceError) {
	err = co.orderSrv.Delete(orderId)
	if err == nil {
		co.cache.DeleteByTag(utils.OrdersTag, orderId)
	}

	return
}

// Edit implements OrderService.
func (co *cachedOrderService) Edit(orderId string, req *forms.EditOrder) (order *models.Order, err *se.ServiceError) {
	order, err = co.orderSrv.Edit(orderId, req)
	if err == nil {
		co.cache.DeleteByTag(utils.OrdersTag, orderId)
	}

	return
}

// Get implements OrderService.
func (co *cachedOrderService) Get(orderId string) (order *models.Order, err *se.ServiceError) {
	key := fmt.Sprintf("orders:%v", orderId)
	er := co.cache.Get(key, &order)
	if er == nil {
		return
	}

	order, err = co.orderSrv.Get(orderId)
	if err != nil {
		return
	}

	co.cache.AddByTag(key, order, orderId)
	return
}

// GetAll implements OrderService.
func (co *cachedOrderService) GetAll() (orders []*models.Order, err *se.ServiceError) {
	er := co.cache.Get(utils.OrdersTag, &orders)
	if er == nil {
		return
	}

	orders, err = co.orderSrv.GetAll()
	if err != nil {
		return
	}

	co.cache.AddByTag(utils.OrdersTag, orders, utils.OrdersTag)
	return

}

func NewCachedOrderService(orderSrv OrderService, cache repo.Cache) OrderService {
	return &cachedOrderService{orderSrv: orderSrv, cache: cache}
}
