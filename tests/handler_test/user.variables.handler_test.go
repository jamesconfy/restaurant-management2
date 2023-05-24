package handler_test

import (
	"fmt"
	"restaurant-management/internal/forms"
	"restaurant-management/internal/models"

	"github.com/bxcodec/faker/v4"
)

func generateUserForm() *forms.User {
	return &forms.User{
		FirstName:   faker.FirstName(),
		LastName:    faker.LastName(),
		PhoneNumber: faker.Phonenumber(),
		Email:       faker.Email(),
		Password:    faker.Password(),
		Address:     faker.MacAddress(),
		Avatar:      faker.IPv4(),
	}
}

func generateAdminForm() *forms.User {
	return &forms.User{
		FirstName:   faker.FirstName(),
		LastName:    faker.LastName(),
		PhoneNumber: faker.Phonenumber(),
		Email:       faker.Email(),
		Password:    faker.Password(),
		Address:     faker.MacAddress(),
		Avatar:      faker.IPv4(),
		Role:        "ADMIN",
	}
}

func createAndRegisterUser(user *forms.User) *models.User {
	if user == nil {
		user = generateUserForm()
	}

	resultUser, err := userSrv.Create(user)
	if err != nil {
		panic(err)
	}

	return resultUser
}

func generateLoginForm(user *forms.User) *forms.Login {
	if user == nil {
		user = generateUserForm()

		_ = createAndRegisterUser(user)
	}

	return &forms.Login{
		Email:    user.Email,
		Password: user.Password,
	}
}

func loginUserAndGenerateAuth(login *forms.Login) string {
	if login == nil {
		login = generateLoginForm(nil)
	}

	auth, err := userSrv.Login(login)
	if err != nil {
		panic(err)
	}

	return fmt.Sprintf("Bearer %v", auth.AccessToken)
}
