package handlers

import (
	"restaurant-management/internal/forms"
	"restaurant-management/internal/response"
	"restaurant-management/internal/se"
	"restaurant-management/internal/service"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

type MenuHandler interface {
	Add(c *gin.Context)
	Get(c *gin.Context)
	GetAll(c *gin.Context)
	Edit(c *gin.Context)
	Delete(c *gin.Context)
}

type menuHanlder struct {
	menuSrv service.MenuService
	cashbin *casbin.Enforcer
}

// Add Menu godoc
// @Summary	Add Menu route
// @Description	Provide details to add menu
// @Tags	Menu
// @Accept	json
// @Produce	json
// @Param	request	body	forms.Menu	true "Menu details"
// @Success	200  {object}  response.SuccessMessage{data=models.Menu}
// @Failure	400  {object}  se.ServiceError
// @Failure	404  {object}  se.ServiceError
// @Failure	500  {object}  se.ServiceError
// @Router	/menus [post]
// @Security ApiKeyAuth
func (m *menuHanlder) Add(c *gin.Context) {
	var req forms.Menu

	if err := c.ShouldBind(&req); err != nil {
		response.Error(c, *se.Validating(err))
		return
	}

	menu, err := m.menuSrv.Add(&req)
	if err != nil {
		response.Error(c, *err)
		return
	}

	response.Success(c, "menu created successfully", menu)
}

// Get menu godoc
// @Summary	Get menu route
// @Description	Provide details to get menu
// @Tags	Menu
// @Accept	json
// @Produce	json
// @Param	menuId	path	string	true "Menu Id" Format(uuid)
// @Success	200  {object}  response.SuccessMessage{data=models.MenuFood}
// @Failure	400  {object}  se.ServiceError
// @Failure	404  {object}  se.ServiceError
// @Failure	500  {object}  se.ServiceError
// @Router	/menus/{menuId} [get]
func (m *menuHanlder) Get(c *gin.Context) {
	menu, err := m.menuSrv.Get(c.Param("menuId"))
	if err != nil {
		response.Error(c, *err)
		return
	}

	response.Success(c, "menu gotten successfully", menu)
}

// Get all menu godoc
// @Summary	Get all menu route
// @Description	Provide details to get all menu
// @Tags	Menu
// @Accept	json
// @Produce	json
// @Success	200  {object}  response.SuccessMessage{data=[]models.Menu}
// @Failure	400  {object}  se.ServiceError
// @Failure	404  {object}  se.ServiceError
// @Failure	500  {object}  se.ServiceError
// @Router	/menus [get]
func (m *menuHanlder) GetAll(c *gin.Context) {
	menu, err := m.menuSrv.GetAll()
	if err != nil {
		response.Error(c, *err)
		return
	}

	response.Success(c, "menus gotten successfully", menu, len(menu))
}

// Edit Menu godoc
// @Summary	Edit menu route
// @Description	Provide details to edit menu
// @Tags	Menu
// @Accept	json
// @Produce	json
// @Param	menuId	path	string	true "Menu Id" Format(uuid)
// @Param	request	body	forms.EditMenu	true "Menu details"
// @Success	200  {object}  response.SuccessMessage{data=models.Menu}
// @Failure	400  {object}  se.ServiceError
// @Failure	404  {object}  se.ServiceError
// @Failure	500  {object}  se.ServiceError
// @Router	/menus/{menuId} [patch]
// @Security ApiKeyAuth
func (m *menuHanlder) Edit(c *gin.Context) {
	var req forms.EditMenu

	if err := c.ShouldBind(&req); err != nil {
		response.Error(c, *se.Validating(err))
		return
	}

	menu, err := m.menuSrv.Edit(c.Param("menuId"), &req)
	if err != nil {
		response.Error(c, *err)
		return
	}

	response.Success(c, "menu updated successfully", menu)
}

// Delete Menu godoc
// @Summary	Delete menu route
// @Description	Provide details to delete a menu
// @Tags	Menu
// @Accept	json
// @Produce	json
// @Param	menuId	path	string	true "Menu Id" Format(uuid)
// @Success	200  {string}  string	"menu deleted successfully"
// @Failure	400  {object}  se.ServiceError
// @Failure	404  {object}  se.ServiceError
// @Failure	500  {object}  se.ServiceError
// @Router	/menus/{menuId} [delete]
// @Security ApiKeyAuth
func (m *menuHanlder) Delete(c *gin.Context) {
	err := m.menuSrv.Delete(c.Param("menuId"))
	if err != nil {
		response.Error(c, *err)
		return
	}

	response.Success202(c, "menu deleted successfully")
}

func NewMenuHandler(menuSrv service.MenuService, cashbin *casbin.Enforcer) MenuHandler {
	return &menuHanlder{menuSrv: menuSrv, cashbin: cashbin}
}
