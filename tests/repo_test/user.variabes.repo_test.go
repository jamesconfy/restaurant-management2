package repo_test

import (
	"restaurant-management/internal/models"

	"github.com/bxcodec/faker/v4"
)

func generateUser() *models.User {
	return &models.User{
		FirstName:   faker.FirstName(),
		LastName:    faker.LastName(),
		Email:       faker.Email(),
		PhoneNumber: faker.Phonenumber(),
		Password:    faker.Password(),
	}
}

func createAndAddUser(user *models.User) *models.User {
	if user == nil {
		user = generateUser()
	}

	user, err := userRepo.Add(user)
	if err != nil {
		panic(err)
	}

	return user
}
