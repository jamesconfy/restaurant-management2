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

func (ta *tableHandler) Add(c *gin.Context) {
	var req forms.Table

	if err := c.ShouldBind(&req); err != nil {
		response.Error(c, *se.Validating(err))
		return
	}

	table, err := ta.tableSrv.Add(&req)
	if err != nil {
		response.Error(c, *err)
	}

	response.Success(c, "table created successfully", table)
}

func (ta *tableHandler) Get(c *gin.Context) {
	table, err := ta.tableSrv.Get(c.Param("tableId"))
	if err != nil {
		response.Error(c, *err)
		return
	}

	response.Success(c, "table fetched successfully", table)
}

func (ta *tableHandler) GetAll(c *gin.Context) {
	tables, err := ta.tableSrv.GetAll(c.GetString("role"))
	if err != nil {
		response.Error(c, *err)
		return
	}

	response.Success(c, "tables fetched successfully", tables, len(tables))
}

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
