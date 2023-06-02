package handlers

import (
	"restaurant-management/internal/forms"
	"restaurant-management/internal/response"
	"restaurant-management/internal/se"
	"restaurant-management/internal/service"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

type TableHanlder interface {
	Add(c *gin.Context)
	Get(c *gin.Context)
	GetAll(c *gin.Context)
	Edit(c *gin.Context)
	Delete(c *gin.Context)
}

type tableHandler struct {
	tableSrv service.TableService
	cashbin  *casbin.Enforcer
}

// Add Table godoc
// @Summary	Add table route
// @Description	Provide details to add table
// @Tags	Table
// @Accept	json
// @Produce	json
// @Param	request	body	forms.Table	true "Table details"
// @Success	200  {object}  response.SuccessMessage{data=models.Table}
// @Failure	400  {object}  se.ServiceError
// @Failure	404  {object}  se.ServiceError
// @Failure	500  {object}  se.ServiceError
// @Router	/tables [post]
// @Security ApiKeyAuth
func (ta *tableHandler) Add(c *gin.Context) {
	var req forms.Table

	if err := c.ShouldBind(&req); err != nil {
		response.Error(c, *se.Validating(err))
		return
	}

	table, err := ta.tableSrv.Add(&req)
	if err != nil {
		response.Error(c, *err)
		return
	}

	response.Success(c, "table created successfully", table)
}

// Get Table godoc
// @Summary	Get table route
// @Description	Provide details to get table
// @Tags	Table
// @Accept	json
// @Produce	json
// @Param	tableId	path	string	true "Table Id"  Format(uuid)
// @Success	200  {object}  response.SuccessMessage{data=models.Table}
// @Failure	400  {object}  se.ServiceError
// @Failure	404  {object}  se.ServiceError
// @Failure	500  {object}  se.ServiceError
// @Router	/tables/{tableId} [get]
// @Security ApiKeyAuth
func (ta *tableHandler) Get(c *gin.Context) {
	table, err := ta.tableSrv.Get(c.Param("tableId"))
	if err != nil {
		response.Error(c, *err)
		return
	}

	response.Success(c, "table fetched successfully", table)
}

// Get All Table godoc
// @Summary	Get all table route
// @Description	Provide details to get table
// @Tags	Table
// @Accept	json
// @Produce	json
// @Success	200  {object}  response.SuccessMessage{data=[]models.Table}
// @Failure	400  {object}  se.ServiceError
// @Failure	404  {object}  se.ServiceError
// @Failure	500  {object}  se.ServiceError
// @Router	/tables [get]
// @Security ApiKeyAuth
func (ta *tableHandler) GetAll(c *gin.Context) {
	tables, err := ta.tableSrv.GetAll(c.GetString("role"))
	if err != nil {
		response.Error(c, *err)
		return
	}

	response.Success(c, "tables fetched successfully", tables, len(tables))
}

// Edit Table godoc
// @Summary	Edit table route
// @Description	Provide details to edit table
// @Tags	Table
// @Accept	json
// @Produce	json
// @Param	tableId	path	string	true "Table Id"  Format(uuid)
// @Param	request	body	forms.EditTable	true "Table details"
// @Success	200  {object}  response.SuccessMessage{data=models.Table}
// @Failure	400  {object}  se.ServiceError
// @Failure	404  {object}  se.ServiceError
// @Failure	500  {object}  se.ServiceError
// @Router	/tables/{tableId} [patch]
// @Security ApiKeyAuth
func (ta *tableHandler) Edit(c *gin.Context) {
	var req forms.EditTable

	if err := c.ShouldBind(&req); err != nil {
		response.Error(c, *se.Validating(err))
		return
	}

	table, err := ta.tableSrv.Edit(c.Param("tableId"), &req)
	if err != nil {
		response.Error(c, *err)
		return
	}

	response.Success(c, "table updated successfully", table)
}

// Delete Table godoc
// @Summary	Delete table route
// @Description	Provide details to delete a table
// @Tags	Table
// @Accept	json
// @Produce	json
// @Param	tableId	path	string	true "Table Id"  Format(uuid)
// @Success	200  {string}  string	"table deleted successfully"
// @Failure	400  {object}  se.ServiceError
// @Failure	404  {object}  se.ServiceError
// @Failure	500  {object}  se.ServiceError
// @Router	/tables/{tableId} [delete]
// @Security ApiKeyAuth
func (ta *tableHandler) Delete(c *gin.Context) {
	err := ta.tableSrv.Delete(c.Param("tableId"))
	if err != nil {
		response.Error(c, *err)
		return
	}

	response.Success202(c, "table deleted successfully")
}

func NewTableHandler(tableSrv service.TableService, cashbin *casbin.Enforcer) TableHanlder {
	return &tableHandler{tableSrv: tableSrv, cashbin: cashbin}
}
