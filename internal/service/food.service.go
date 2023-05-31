package service

import (
	"restaurant-management/internal/forms"
	"restaurant-management/internal/models"
	repo "restaurant-management/internal/repository"
	"restaurant-management/internal/se"

	"github.com/docker/distribution/uuid"
)

type FoodService interface {
	Add(menuId string, req *forms.Food) (*models.Food, *se.ServiceError)
	Get(foodId string) (*models.Food, *se.ServiceError)
	GetAll() ([]*models.Food, *se.ServiceError)
	GetFoodsByMenu(menuId string) (*models.MenuFood, *se.ServiceError)
	Edit(foodId string, req *forms.EditFood) (*models.Food, *se.ServiceError)
	Delete(foodId string) *se.ServiceError
}

type foodSrv struct {
	repo     repo.FoodRepo
	menuRepo repo.MenuRepo
}

// Add implements FoodService
func (f *foodSrv) Add(menuId string, req *forms.Food) (*models.Food, *se.ServiceError) {
	if err := Validator.validate(req); err != nil {
		return nil, se.Validating(err)
	}

	if _, err := uuid.Parse(menuId); err != nil {
		return nil, se.Validating(err)
	}

	ok, err := f.menuRepo.CheckMenuExists(menuId)
	if err != nil || !ok {
		return nil, se.NotFoundOrInternal(err, "menu not found")
	}

	var food models.Food

	food.Name = req.Name
	food.Price = req.Price
	food.Image = req.Image
	food.MenuId = menuId

	foo, err := f.repo.Add(&food)
	if err != nil {
		return nil, se.Internal(err, "could not create food")
	}

	return foo, nil
}

// Get implements FoodService
func (f *foodSrv) Get(foodId string) (*models.Food, *se.ServiceError) {
	if _, err := uuid.Parse(foodId); err != nil {
		return nil, se.Validating(err)
	}

	food, err := f.repo.Get(foodId)
	if err != nil {
		return nil, se.NotFoundOrInternal(err, "food not found")
	}

	return food, nil
}

// GetAll implements FoodService
func (f *foodSrv) GetAll() ([]*models.Food, *se.ServiceError) {
	foods, err := f.repo.GetAll()
	if err != nil {
		return nil, se.Internal(err, "could not fetch foods")
	}

	return foods, nil
}

// GetFoodByMenu implements FoodService
func (f *foodSrv) GetFoodsByMenu(menuId string) (*models.MenuFood, *se.ServiceError) {
	if _, err := uuid.Parse(menuId); err != nil {
		return nil, se.Validating(err)
	}

	ok, err := f.menuRepo.CheckMenuExists(menuId)
	if err != nil || !ok {
		return nil, se.NotFoundOrInternal(err, "menu not found")
	}

	food, err := f.repo.GetFoodByMenu(menuId)
	if err != nil {
		return nil, se.NotFoundOrInternal(err, "food not found")
	}

	return food, nil
}

// Edit implements FoodService
func (f *foodSrv) Edit(foodId string, req *forms.EditFood) (*models.Food, *se.ServiceError) {
	if _, err := uuid.Parse(foodId); err != nil {
		return nil, se.Validating(err)
	}

	if err := Validator.validate(req); err != nil {
		return nil, se.Validating(err)
	}

	food, err := f.repo.Get(foodId)
	if err != nil {
		return nil, se.NotFoundOrInternal(err, "food not found")
	}

	food = f.getEdit(req, food)

	foo, err := f.repo.Edit(foodId, food)
	if err != nil {
		return nil, se.Internal(err)
	}

	return foo, nil
}

// Delete implements FoodService
func (f *foodSrv) Delete(foodId string) *se.ServiceError {
	if _, err := uuid.Parse(foodId); err != nil {
		return se.Validating(err)
	}

	err := f.repo.Delete(foodId)
	if err != nil {
		return se.Internal(err)
	}

	return nil
}

func NewFoodService(repo repo.FoodRepo, menuRepo repo.MenuRepo) FoodService {
	return &foodSrv{repo: repo, menuRepo: menuRepo}
}

// Auxillary Function
func (f *foodSrv) getEdit(req *forms.EditFood, food *models.Food) *models.Food {
	if req.Name != "" && req.Name != food.Name {
		food.Name = req.Name
	}

	if req.Image != "" && req.Image != food.Image {
		food.Image = req.Image
	}

	if req.Price != 0.0 && req.Price != food.Price {
		food.Price = req.Price
	}

	return food
}
