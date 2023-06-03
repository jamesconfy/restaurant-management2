package service

import (
	"restaurant-management/internal/forms"
	"restaurant-management/internal/models"
	repo "restaurant-management/internal/repository"
	"restaurant-management/internal/se"
)

type cachedFoodService struct {
	foodSrv FoodService
	cache   repo.Cache
}

// Add implements FoodService
func (cf *cachedFoodService) Add(req *forms.Food) (*models.Food, *se.ServiceError) {
	panic("unimplemented")
}

// Delete implements FoodService
func (cf *cachedFoodService) Delete(foodId string) *se.ServiceError {
	panic("unimplemented")
}

// Edit implements FoodService
func (cf *cachedFoodService) Edit(foodId string, req *forms.EditFood) (*models.Food, *se.ServiceError) {
	panic("unimplemented")
}

// Get implements FoodService
func (cf *cachedFoodService) Get(foodId string) (*models.Food, *se.ServiceError) {
	panic("unimplemented")
}

// GetAll implements FoodService
func (cf *cachedFoodService) GetAll() ([]*models.Food, *se.ServiceError) {
	panic("unimplemented")
}

func NewCachedFoodService(foodSrv FoodService, cache repo.Cache) FoodService {
	return &cachedFoodService{foodSrv: foodSrv, cache: cache}
}
