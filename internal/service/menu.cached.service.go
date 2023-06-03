package service

import (
	"restaurant-management/internal/forms"
	"restaurant-management/internal/models"
	repo "restaurant-management/internal/repository"
	"restaurant-management/internal/se"
)

type cachedMenuService struct {
	menuSrv MenuService
	cache   repo.Cache
}

// Add implements MenuService
func (cm *cachedMenuService) Add(req *forms.Menu) (*models.Menu, *se.ServiceError) {
	panic("unimplemented")
}

// Delete implements MenuService
func (cm *cachedMenuService) Delete(menuId string) *se.ServiceError {
	panic("unimplemented")
}

// Edit implements MenuService
func (cm *cachedMenuService) Edit(menuId string, req *forms.EditMenu) (*models.Menu, *se.ServiceError) {
	panic("unimplemented")
}

// Get implements MenuService
func (cm *cachedMenuService) Get(menuId string) (*models.MenuFood, *se.ServiceError) {
	panic("unimplemented")
}

// GetAll implements MenuService
func (cm *cachedMenuService) GetAll() ([]*models.Menu, *se.ServiceError) {
	panic("unimplemented")
}

func NewCachedMenuService(menuSrv MenuService, cache repo.Cache) MenuService {
	return &cachedMenuService{menuSrv: menuSrv, cache: cache}
}
