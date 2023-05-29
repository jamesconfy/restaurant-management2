package service_test

import (
	"restaurant-management/internal/forms"
	"restaurant-management/internal/models"

	"github.com/bxcodec/faker/v4"
	"golang.org/x/crypto/bcrypt"
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

func generateEditUserForm() *forms.EditUser {
	return &forms.EditUser{
		FirstName:   faker.FirstName(),
		LastName:    faker.LastName(),
		PhoneNumber: faker.Phonenumber(),
		Email:       faker.Email(),
		Address:     faker.MacAddress(),
		Avatar:      faker.IPv4(),
	}
}

func generateUserModel(user *forms.User) *models.User {
	return &models.User{
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		PhoneNumber: user.PhoneNumber,
		Password:    user.Password,
		Email:       user.Email,
		Address:     user.Address,
		Avatar:      user.Avatar,
	}
}

func createAndRegisterUser(user *forms.User) *models.User {
	if user == nil {
		user = generateUserForm()
	}

	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		panic(err)
	}

	user.Password = string(password)
	resultUser, err := userRepo.Add(generateUserModel(user))
	if err != nil {
		panic(err)
	}

	return resultUser
}

func createUserAuth(user *models.User) *models.Auth {
	var auth models.Auth
	var err error

	if user == nil {
		user = createAndRegisterUser(nil)
	}

	auth.AccessToken, auth.RefreshToken, err = authSrv.Create(user)
	if err != nil {
		panic(err)
	}

	auth.UserId = user.Id
	authRe, err := authRepo.Add(&auth)
	if err != nil {
		panic(err)
	}

	return authRe
}
