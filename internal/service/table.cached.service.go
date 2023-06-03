package service

import (
	"restaurant-management/internal/forms"
	"restaurant-management/internal/models"
	repo "restaurant-management/internal/repository"
	"restaurant-management/internal/se"
)

type cachedTableService struct {
	tableSrv TableService
	cache    repo.Cache
}

// Add implements TableService
func (ct *cachedTableService) Add(req *forms.Table) (*models.Table, *se.ServiceError) {
	panic("unimplemented")
}

// Delete implements TableService
func (ct *cachedTableService) Delete(tableId string) *se.ServiceError {
	panic("unimplemented")
}

// Edit implements TableService
func (ct *cachedTableService) Edit(tableId string, req *forms.EditTable) (*models.Table, *se.ServiceError) {
	panic("unimplemented")
}

// Get implements TableService
func (ct *cachedTableService) Get(tableId string) (*models.Table, *se.ServiceError) {
	panic("unimplemented")
}

// GetAll implements TableService
func (ct *cachedTableService) GetAll(role string) ([]*models.Table, *se.ServiceError) {
	panic("unimplemented")
}

func NewCachedTableService(tableSrv TableService, cache repo.Cache) TableService {
	return &cachedTableService{tableSrv: tableSrv, cache: cache}
}
