package handlers

import (
	"restaurant-management/internal/forms"
	"restaurant-management/internal/response"
	"restaurant-management/internal/se"
	"restaurant-management/internal/service"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

type OrderHandler interface {
	Add(c *gin.Context)
	Get(c *gin.Context)
	GetAll(c *gin.Context)
	Edit(c *gin.Context)
	Delete(c *gin.Context)
}

type orderHandler struct {
	orderSrv service.OrderService
	casbin   *casbin.Enforcer
}

// Add implements OrderHandler.
func (o *orderHandler) Add(c *gin.Context) {
	var req forms.Order

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, *se.Validating(err))
		return
	}

	order, err := o.orderSrv.Add(&req)
	if err != nil {
		response.Error(c, *err)
		return
	}

	response.Success(c, "order added successfully", order)
}

// Delete implements OrderHandler.
func (o *orderHandler) Delete(c *gin.Context) {
	err := o.orderSrv.Delete(c.Param("orderId"))
	if err != nil {
		response.Error(c, *err)
		return
	}

	response.Success202(c, "order deleted successfully")
}

// Edit implements OrderHandler.
func (o *orderHandler) Edit(c *gin.Context) {
	var req forms.EditOrder

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, *se.Validating(err))
		return
	}

	order, err := o.orderSrv.Edit(c.Param("orderId"), &req)
	if err != nil {
		response.Error(c, *err)
		return
	}

	response.Success(c, "order updated successfully", order)
}

// Get implements OrderHandler.
func (o *orderHandler) Get(c *gin.Context) {
	order, err := o.orderSrv.Get(c.Param("orderId"))
	if err != nil {
		response.Error(c, *err)
		return
	}

	response.Success(c, "order gotten successfully", order)
}

// GetAll implements OrderHandler.
func (o *orderHandler) GetAll(c *gin.Context) {
	orders, err := o.orderSrv.GetAll()
	if err != nil {
		response.Error(c, *err)
		return
	}

	response.Success(c, "orders gotten successfully", orders, len(orders))
}

func NewOrderHandler(orderSrv service.OrderService, casbin *casbin.Enforcer) OrderHandler {
	return &orderHandler{orderSrv: orderSrv, casbin: casbin}
}
