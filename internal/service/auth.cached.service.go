package service

import (
	"fmt"
	"restaurant-management/config"
	"restaurant-management/internal/models"
	repo "restaurant-management/internal/repository"
	"time"

	"github.com/go-redis/redis/v8"
)

type cachedAuthService struct {
	authSrv AuthService
	cache   repo.Cache
}

// Create implements AuthService
func (ca *cachedAuthService) Create(user *models.User) (string, string, error) {
	return ca.authSrv.Create(user)
}

// Validate implements AuthService
func (ca *cachedAuthService) Validate(url string) (*Token, error) {
	var tok *Token

	key := fmt.Sprintf("validate:%v", url)
	err := ca.cache.Get(key, &tok)
	if err == nil && err != redis.Nil {
		return tok, err
	}

	toke, er := ca.authSrv.Validate(url)
	if er != nil {
		return nil, er
	}

	ca.cache.AddByTag(key, toke, toke.Id, time.Now().Add(time.Hour*time.Duration(config.Environment.EXPIRES_AT)))
	return toke, nil
}

func NewCachedAuthService(authSrv AuthService, cache repo.Cache) AuthService {
	return &cachedAuthService{authSrv: authSrv, cache: cache}
}
