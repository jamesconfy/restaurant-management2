package service

import (
	"restaurant-management/internal/forms"
	"restaurant-management/internal/models"
	repo "restaurant-management/internal/repository"
	"restaurant-management/internal/se"

	"github.com/docker/distribution/uuid"
)

type MenuService interface {
	Add(req *forms.Menu) (*models.Menu, *se.ServiceError)
	Get(menuId string) (*models.MenuFood, *se.ServiceError)
	GetAll() ([]*models.Menu, *se.ServiceError)
	Edit(menuId string, req *forms.EditMenu) (*models.Menu, *se.ServiceError)
	Delete(menuId string) *se.ServiceError
}

type menuSrv struct {
	repo repo.MenuRepo
}

func (m *menuSrv) Add(req *forms.Menu) (*models.Menu, *se.ServiceError) {
	if err := Validator.validate(req); err != nil {
		return nil, se.Validating(err)
	}

	if ok, err := m.repo.Check(req.Name, req.Category); ok {
		return nil, se.ConflictOrInternal(err, "name and category already exists")
	}

	var menu models.Menu

	menu.Name = req.Name
	menu.Category = req.Category

	men, err := m.repo.Add(&menu)
	if err != nil {
		return nil, se.Internal(err)
	}

	return men, nil
}

func (m *menuSrv) Get(menuId string) (*models.MenuFood, *se.ServiceError) {
	if _, err := uuid.Parse(menuId); err != nil {
		return nil, se.Validating(err)
	}

	menu, err := m.repo.Get(menuId)
	if err != nil {
		return nil, se.NotFoundOrInternal(err, "menu not found")
	}

	return menu, nil
}

func (m *menuSrv) GetAll() ([]*models.Menu, *se.ServiceError) {
	menu, err := m.repo.GetAll()
	if err != nil {
		return nil, se.Internal(err)
	}

	return menu, nil
}

func (m *menuSrv) Edit(menuId string, req *forms.EditMenu) (*models.Menu, *se.ServiceError) {
	if _, err := uuid.Parse(menuId); err != nil {
		return nil, se.Internal(err, "invalid menu id")
	}

	menu, err := m.repo.Get(menuId)
	if err != nil {
		return nil, se.NotFoundOrInternal(err, "menu not found")
	}

	men, er := m.getEdit(req, menu)
	if er != nil {
		return nil, er
	}

	resultMenu, err := m.repo.Edit(menuId, men)
	if err != nil {
		return nil, se.Internal(err)
	}

	return resultMenu, nil
}

func (m *menuSrv) Delete(menuId string) *se.ServiceError {
	if _, err := uuid.Parse(menuId); err != nil {
		return se.Internal(err, "invalid menu id")
	}

	err := m.repo.Delete(menuId)
	if err != nil {
		return se.NotFoundOrInternal(err, "menu not found")
	}

	return nil
}

func NewMenuService(repo repo.MenuRepo) MenuService {
	return &menuSrv{repo: repo}
}

func (m *menuSrv) getEdit(req *forms.EditMenu, menu *models.MenuFood) (*models.Menu, *se.ServiceError) {
	if req.Name != "" && req.Name != menu.Menu.Name {
		menu.Menu.Name = req.Name
	}

	if req.Category != "" && req.Category != menu.Menu.Category {
		menu.Menu.Category = req.Category
	}

	if ok, err := m.repo.Check(req.Name, req.Category); ok {
		return nil, se.ConflictOrInternal(err, "name and category already exists")
	}

	return menu.Menu, nil
}
