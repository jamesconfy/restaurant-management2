package service

import (
	"fmt"
	"restaurant-management/internal/forms"
	"restaurant-management/internal/models"
	repo "restaurant-management/internal/repository"
	"restaurant-management/internal/se"
	"restaurant-management/utils"
)

type cachedFoodService struct {
	foodSrv FoodService
	cache   repo.Cache
}

// Add implements FoodService
func (cf *cachedFoodService) Add(req *forms.Food) (food *models.Food, err *se.ServiceError) {
	food, err = cf.foodSrv.Add(req)
	if err == nil {
		cf.cache.DeleteByTag(utils.FoodsTag)
	}

	return
}

// Delete implements FoodService
func (cf *cachedFoodService) Delete(foodId string) (err *se.ServiceError) {
	err = cf.foodSrv.Delete(foodId)
	if err == nil {
		cf.cache.DeleteByTag(foodId)
	}
	return
}

// Edit implements FoodService
func (cf *cachedFoodService) Edit(foodId string, req *forms.EditFood) (food *models.Food, err *se.ServiceError) {
	food, err = cf.foodSrv.Edit(foodId, req)
	if err == nil {
		cf.cache.DeleteByTag(utils.FoodsTag, foodId)
	}

	return
}

// Get implements FoodService
func (cf *cachedFoodService) Get(foodId string) (food *models.Food, er *se.ServiceError) {
	key := fmt.Sprintf("foods:%v", foodId)
	err := cf.cache.Get(key, &food)
	if err == nil {
		return food, nil
	}

	food, er = cf.foodSrv.Get(foodId)
	if er != nil {
		return
	}

	cf.cache.AddByTag(key, food, foodId)
	return
}

// GetAll implements FoodService
func (cf *cachedFoodService) GetAll() (foods []*models.Food, er *se.ServiceError) {
	err := cf.cache.Get(utils.FoodsTag, &foods)
	if err == nil {
		return
	}

	foods, er = cf.foodSrv.GetAll()
	if er != nil {
		return
	}

	cf.cache.AddByTag(utils.FoodsTag, foods, utils.FoodsTag)
	return
}

func NewCachedFoodService(foodSrv FoodService, cache repo.Cache) FoodService {
	return &cachedFoodService{foodSrv: foodSrv, cache: cache}
}
