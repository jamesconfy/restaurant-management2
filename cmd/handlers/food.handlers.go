package handlers

import (
	"restaurant-management/internal/forms"
	"restaurant-management/internal/response"
	"restaurant-management/internal/se"
	"restaurant-management/internal/service"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

type FoodHandler interface {
	Add(c *gin.Context)
	Get(c *gin.Context)
	GetAll(c *gin.Context)
	Edit(c *gin.Context)
	Delete(c *gin.Context)
}

type foodHandler struct {
	foodSrv service.FoodService
	cashbin *casbin.Enforcer
}

// Add Food godoc
// @Summary	Add Food route
// @Description	Provide details to add food
// @Tags	Food
// @Accept	json
// @Produce	json
// @Param	request	body	forms.Food	true "Food details"
// @Success	200  {object}  response.SuccessMessage{data=models.Food}
// @Failure	400  {object}  se.ServiceError
// @Failure	404  {object}  se.ServiceError
// @Failure	500  {object}  se.ServiceError
// @Router	/foods [post]
// @Security ApiKeyAuth
func (f *foodHandler) Add(c *gin.Context) {
	var req forms.Food

	if err := c.ShouldBind(&req); err != nil {
		response.Error(c, *se.Validating(err))
		return
	}

	food, err := f.foodSrv.Add(&req)
	if err != nil {
		response.Error(c, *err)
		return
	}

	response.Success(c, "food created successfully", food)
}

// Delete Food godoc
// @Summary	Delete Food route
// @Description	Provide details to delete food
// @Tags	Food
// @Accept	json
// @Produce	json
// @Param	foodId	path	string	true "Food Id"  Format(uuid)
// @Success	200  {string}  string	"food deleted successfully"
// @Failure	400  {object}  se.ServiceError
// @Failure	404  {object}  se.ServiceError
// @Failure	500  {object}  se.ServiceError
// @Router	/foods/{foodId} [delete]
// @Security ApiKeyAuth
func (f *foodHandler) Delete(c *gin.Context) {
	err := f.foodSrv.Delete(c.Param("foodId"))
	if err != nil {
		response.Error(c, *err)
		return
	}

	response.Success202(c, "food gotten successfully")
}

// Edit Food godoc
// @Summary	Edit Food route
// @Description	Provide details to edit food
// @Tags	Food
// @Accept	json
// @Produce	json
// @Param	foodId	path	string	true "Food Id"  Format(uuid)
// @Param	request	body	forms.EditFood	true "Food details"
// @Success	200  {object}  response.SuccessMessage{data=models.Food}
// @Failure	400  {object}  se.ServiceError
// @Failure	404  {object}  se.ServiceError
// @Failure	500  {object}  se.ServiceError
// @Router	/foods/{foodId} [patch]
// @Security ApiKeyAuth
func (f *foodHandler) Edit(c *gin.Context) {
	var req forms.EditFood

	if err := c.ShouldBind(&req); err != nil {
		response.Error(c, *se.Validating(err))
		return
	}

	food, err := f.foodSrv.Edit(c.Param("foodId"), &req)
	if err != nil {
		response.Error(c, *err)
		return
	}

	response.Success(c, "food updated successfully", food)
}

// Get Food godoc
// @Summary	Get Food route
// @Description	Provide details to get food
// @Tags	Food
// @Accept	json
// @Produce	json
// @Param	foodId	path	string	true "Food Id"  Format(uuid)
// @Success	200  {object}  response.SuccessMessage{data=models.Food}
// @Failure	400  {object}  se.ServiceError
// @Failure	404  {object}  se.ServiceError
// @Failure	500  {object}  se.ServiceError
// @Router	/foods/{foodId} [get]
func (f *foodHandler) Get(c *gin.Context) {
	food, err := f.foodSrv.Get(c.Param("foodId"))
	if err != nil {
		response.Error(c, *err)
		return
	}

	response.Success(c, "food gotten successfully", food)
}

// Get All Food godoc
// @Summary	Get all Food route
// @Description	Provide details to get all food
// @Tags	Food
// @Accept	json
// @Produce	json
// @Success	200  {object}  response.SuccessMessage{data=[]models.Food}
// @Failure	400  {object}  se.ServiceError
// @Failure	404  {object}  se.ServiceError
// @Failure	500  {object}  se.ServiceError
// @Router	/foods [get]
func (f *foodHandler) GetAll(c *gin.Context) {
	

	foods, err := f.foodSrv.GetAll()
	if err != nil {
		response.Error(c, *err)
		return
	}

	response.Success(c, "foods fetched successfully", foods, len(foods))
}

func NewFoodHanlder(foodSrv service.FoodService, cashbin *casbin.Enforcer) FoodHandler {
	return &foodHandler{foodSrv: foodSrv, cashbin: cashbin}
}
