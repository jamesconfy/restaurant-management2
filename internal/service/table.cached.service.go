package service

import (
	"fmt"
	"restaurant-management/internal/forms"
	"restaurant-management/internal/models"
	repo "restaurant-management/internal/repository"
	"restaurant-management/internal/se"
	"restaurant-management/utils"
)

type cachedTableService struct {
	tableSrv TableService
	cache    repo.Cache
}

// Add implements TableService
func (ct *cachedTableService) Add(req *forms.Table) (table *models.Table, err *se.ServiceError) {
	table, err = ct.tableSrv.Add(req)
	if err == nil {
		ct.cache.DeleteByTag(utils.TablesTag)
	}

	return
}

// Delete implements TableService
func (ct *cachedTableService) Delete(tableId string) (err *se.ServiceError) {
	err = ct.tableSrv.Delete(tableId)
	if err == nil {
		ct.cache.DeleteByTag(utils.TablesTag, tableId)
	}

	return
}

// Edit implements TableService
func (ct *cachedTableService) Edit(tableId string, req *forms.EditTable) (table *models.Table, err *se.ServiceError) {
	table, err = ct.tableSrv.Edit(tableId, req)
	if err == nil {
		ct.cache.DeleteByTag(tableId)
	}

	return
}

// Get implements TableService
func (ct *cachedTableService) Get(tableId string) (table *models.Table, err *se.ServiceError) {
	key := fmt.Sprintf("tables:%v", tableId)
	er := ct.cache.Get(key, &table)
	if er == nil {
		return
	}

	table, err = ct.tableSrv.Get(tableId)
	if err != nil {
		return
	}

	ct.cache.AddByTag(key, table, tableId)
	return
}

// GetAll implements TableService
func (ct *cachedTableService) GetAll(role string) (tables []*models.Table, err *se.ServiceError) {
	key := fmt.Sprintf("%v:%v", utils.TablesTag, role)
	er := ct.cache.Get(key, &tables)
	if er == nil {
		return
	}

	tables, err = ct.tableSrv.GetAll(role)
	if err != nil {
		return
	}

	ct.cache.AddByTag(key, tables, utils.TablesTag)
	return
}

func NewCachedTableService(tableSrv TableService, cache repo.Cache) TableService {
	return &cachedTableService{tableSrv: tableSrv, cache: cache}
}
