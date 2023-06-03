package service

import (
	"fmt"
	"restaurant-management/internal/forms"
	"restaurant-management/internal/models"
	repo "restaurant-management/internal/repository"
	"restaurant-management/internal/se"
)

type cachedUserService struct {
	userService UserService
	cache       repo.Cache
}

// Add implements UserService
func (cu *cachedUserService) Add(req *forms.User) (*models.User, *se.ServiceError) {
	cu.cache.DeleteByTag("users:all")
	return cu.userService.Add(req)
}

// Delete implements UserService
func (cu *cachedUserService) Delete(userId string) *se.ServiceError {
	cu.cache.DeleteByTag("users:all", userId)
	return cu.userService.Delete(userId)
}

// DeleteToken implements UserService
func (cu *cachedUserService) DeleteAuth(userId, accessToken string) *se.ServiceError {
	return cu.userService.DeleteAuth(userId, accessToken)
}

// DeleteToken implements UserService
func (cu *cachedUserService) ClearAuth(userId, accessToken string) *se.ServiceError {
	cu.cache.DeleteByTag(userId)
	return cu.userService.ClearAuth(userId, accessToken)
}

// Edit implements UserService
func (cu *cachedUserService) Edit(userId string, req *forms.EditUser) (*models.User, *se.ServiceError) {
	user, err := cu.userService.Edit(userId, req)
	if err == nil {
		key := fmt.Sprintf("users:%s", userId)

		cu.cache.DeleteByTag("users:all", key)
	}

	return user, err
}

// GetAll implements UserService
func (cu *cachedUserService) GetAll() ([]*models.User, *se.ServiceError) {
	var users []*models.User

	err := cu.cache.Get("users:all", &users)
	if err == nil {
		return users, nil
	}

	users, er := cu.userService.GetAll()
	if er != nil {
		return nil, er
	}

	cu.cache.AddByTag("users:all", users, "users:all")
	return users, nil
}

// Get implements UserService
func (cu *cachedUserService) Get(userId string) (*models.User, *se.ServiceError) {
	var user *models.User

	key := fmt.Sprintf("users:%s", userId)
	err := cu.cache.Get(key, &user)
	if err == nil {
		return user, nil
	}

	user, er := cu.userService.Get(userId)
	if er != nil {
		return nil, er
	}

	cu.cache.AddByTag(key, user, userId)
	return user, er
}

// Login implements UserService
func (cu *cachedUserService) Login(req *forms.Login) (*models.Auth, *se.ServiceError) {
	return cu.userService.Login(req)
}

func NewCachedUserService(userService UserService, cache repo.Cache) UserService {
	return &cachedUserService{userService: userService, cache: cache}
}
