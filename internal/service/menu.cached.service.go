package service

import (
	"fmt"
	"restaurant-management/internal/forms"
	"restaurant-management/internal/models"
	repo "restaurant-management/internal/repository"
	"restaurant-management/internal/se"
	"restaurant-management/utils"
)

type cachedMenuService struct {
	menuSrv MenuService
	cache   repo.Cache
}

// Add implements MenuService
func (cm *cachedMenuService) Add(req *forms.Menu) (menu *models.Menu, err *se.ServiceError) {
	menu, err = cm.menuSrv.Add(req)
	if err == nil {
		cm.cache.DeleteByTag(utils.MenusTag)
	}

	return
}

// Delete implements MenuService
func (cm *cachedMenuService) Delete(menuId string) (err *se.ServiceError) {
	err = cm.menuSrv.Delete(menuId)
	if err == nil {
		cm.cache.DeleteByTag(utils.MenusTag, menuId)
	}

	return
}

// Edit implements MenuService
func (cm *cachedMenuService) Edit(menuId string, req *forms.EditMenu) (menu *models.Menu, err *se.ServiceError) {
	menu, err = cm.menuSrv.Edit(menuId, req)
	if err == nil {
		cm.cache.DeleteByTag(menuId)
	}

	return
}

// Get implements MenuService
func (cm *cachedMenuService) Get(menuId string) (menu *models.MenuFood, err *se.ServiceError) {
	key := fmt.Sprintf("menus:%v", menuId)
	er := cm.cache.Get(key, &menu)
	if er == nil {
		return
	}

	menu, err = cm.menuSrv.Get(menuId)
	if err != nil {
		return
	}

	cm.cache.AddByTag(key, menu, menuId)
	return
}

// GetAll implements MenuService
func (cm *cachedMenuService) GetAll() (menus []*models.Menu, err *se.ServiceError) {
	er := cm.cache.Get(utils.MenusTag, &menus)
	if er == nil {
		return
	}

	menus, err = cm.menuSrv.GetAll()
	if err != nil {
		return
	}

	cm.cache.AddByTag(utils.MenusTag, menus, utils.MenusTag)
	return
}

func NewCachedMenuService(menuSrv MenuService, cache repo.Cache) MenuService {
	return &cachedMenuService{menuSrv: menuSrv, cache: cache}
}
