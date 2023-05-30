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
}

type menuSrv struct {
	repo repo.MenuRepo
}

func (m *menuSrv) Add(req *forms.Menu) (*models.Menu, *se.ServiceError) {
	if err := Validator.Validate(req); err != nil {
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

func (m *menuSrv) Get(menuId string) (*models.Menu, *se.ServiceError) {
	if _, err := uuid.Parse(menuId); err != nil {
		return nil, se.Internal(err, "invalid menu id")
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

	menu, er := m.getEdit(req, menu)
	if er != nil {
		return nil, er
	}

	men, err := m.repo.Edit(menu)
	if err != nil {
		return nil, se.Internal(err)
	}

	return men, nil
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

func (m *menuSrv) getEdit(req *forms.EditMenu, menu *models.Menu) (*models.Menu, *se.ServiceError) {
	if req.Name != "" && req.Name != menu.Name {
		menu.Name = req.Name
	}

	if req.Category != "" && req.Category != menu.Category {
		menu.Category = req.Category
	}

	if ok, err := m.repo.Check(req.Name, req.Category); ok {
		return nil, se.ConflictOrInternal(err, "name and category already exists")
	}

	return menu, nil
}
