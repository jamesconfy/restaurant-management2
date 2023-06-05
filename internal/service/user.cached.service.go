package service

import (
	"fmt"
	"restaurant-management/internal/forms"
	"restaurant-management/internal/models"
	repo "restaurant-management/internal/repository"
	"restaurant-management/internal/se"
	"restaurant-management/utils"
)

type cachedUserService struct {
	userService UserService
	cache       repo.Cache
}

// Add implements UserService
func (cu *cachedUserService) Add(req *forms.User) (user *models.User, err *se.ServiceError) {
	user, err = cu.userService.Add(req)
	if err == nil {
		cu.cache.DeleteByTag(utils.UsersTag)
	}

	return
}

// Delete implements UserService
func (cu *cachedUserService) Delete(userId string) (err *se.ServiceError) {
	err = cu.userService.Delete(userId)
	if err == nil {
		cu.cache.DeleteByTag(utils.UsersTag, userId)
	}

	return
}

// DeleteToken implements UserService
func (cu *cachedUserService) DeleteAuth(userId, accessToken string) (err *se.ServiceError) {
	err = cu.userService.DeleteAuth(userId, accessToken)
	if err == nil {
		cu.cache.DeleteByTag(userId)
	}

	return
}

// DeleteToken implements UserService
func (cu *cachedUserService) ClearAuth(userId, accessToken string) (err *se.ServiceError) {
	err = cu.userService.ClearAuth(userId, accessToken)
	if err == nil {
		cu.cache.DeleteByTag(userId)
	}

	return
}

// Edit implements UserService
func (cu *cachedUserService) Edit(userId string, req *forms.EditUser) (*models.User, *se.ServiceError) {
	user, err := cu.userService.Edit(userId, req)
	if err == nil {
		cu.cache.DeleteByTag(utils.UsersTag, fmt.Sprintf("users:%s", userId))
	}

	return user, err
}

// GetAll implements UserService
func (cu *cachedUserService) GetAll() ([]*models.User, *se.ServiceError) {
	var users []*models.User

	err := cu.cache.Get(utils.UsersTag, &users)
	if err == nil {
		return users, nil
	}

	users, er := cu.userService.GetAll()
	if er != nil {
		return nil, er
	}

	cu.cache.AddByTag(utils.UsersTag, users, utils.UsersTag)
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
