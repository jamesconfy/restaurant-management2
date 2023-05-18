package service

import "restaurant-management/internal/se"

type HomeService interface {
	CreateHome() (string, *se.ServiceError)
}

type homeSrv struct{}

func (h *homeSrv) CreateHome() (string, *se.ServiceError) {
	return "You have gotten to the home route of restaurant-management", nil
}

func NewHomeService() HomeService {
	return &homeSrv{}
}
