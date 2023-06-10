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
// Add Order godoc
// @Summary	Add order route
// @Description	Provide details to add an order
// @Tags	Order
// @Accept	json
// @Produce	json
// @Param	request	body	forms.Order	true	"Add order form"
// @Success	200  {object}  response.SuccessMessage{data=models.Order}
// @Failure	400  {object}  se.ServiceError
// @Failure	404  {object}  se.ServiceError
// @Failure	500  {object}  se.ServiceError
// @Router	/orders [post]
// @Security ApiKeyAuth
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
// Delete Order godoc
// @Summary	Delete order route
// @Description	Provide details to delete an order
// @Tags	Order
// @Accept	json
// @Produce	json
// @Param	orderId	path	string	true	"Order Id" Format(uuid)
// @Success	200  {string}	string	"order deleted successfully"
// @Failure	400  {object}  se.ServiceError
// @Failure	404  {object}  se.ServiceError
// @Failure	500  {object}  se.ServiceError
// @Router	/orders/{orderId} [delete]
// @Security ApiKeyAuth
func (o *orderHandler) Delete(c *gin.Context) {
	err := o.orderSrv.Delete(c.Param("orderId"))
	if err != nil {
		response.Error(c, *err)
		return
	}

	response.Success202(c, "order deleted successfully")
}

// Edit implements OrderHandler.
// Edit Order godoc
// @Summary	Edit order route
// @Description	Provide details to edit an order
// @Tags	Order
// @Accept	json
// @Produce	json
// @Param	request	body	forms.EditOrder	true	"Edit order form"
// @Param	orderId	path	string	true	"Order Id" Format(uuid)
// @Success	200  {object}	response.SuccessMessage{data=models.Order}
// @Failure	400  {object}  se.ServiceError
// @Failure	404  {object}  se.ServiceError
// @Failure	500  {object}  se.ServiceError
// @Router	/orders/{orderId} [patch]
// @Security ApiKeyAuth
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
// Get Order godoc
// @Summary	Get order route
// @Description	Provide order id to get a particular order
// @Tags	Order
// @Accept	json
// @Produce	json
// @Param	orderId	path	string	true	"Order Id" Format(uuid)
// @Success	200  {object}	response.SuccessMessage{data=models.Order}
// @Failure	400  {object}  se.ServiceError
// @Failure	404  {object}  se.ServiceError
// @Failure	500  {object}  se.ServiceError
// @Router	/orders/{orderId} [get]
// @Security ApiKeyAuth
func (o *orderHandler) Get(c *gin.Context) {
	order, err := o.orderSrv.Get(c.Param("orderId"))
	if err != nil {
		response.Error(c, *err)
		return
	}

	response.Success(c, "order gotten successfully", order)
}

// GetAll implements OrderHandler.
// Get All Order godoc
// @Summary	Get all order route
// @Description	Get all orders in the system
// @Tags	Order
// @Accept	json
// @Produce	json
// @Success	200  {object}	response.SuccessMessage{data=[]models.Order}
// @Failure	400  {object}  se.ServiceError
// @Failure	404  {object}  se.ServiceError
// @Failure	500  {object}  se.ServiceError
// @Router	/orders [get]
// @Security ApiKeyAuth
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