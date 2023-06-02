package service_test

import (
	"restaurant-management/internal/forms"
	"restaurant-management/internal/models"

	"github.com/bxcodec/faker/v4"
)

func generateMenuForm() *forms.Menu {
	return &forms.Menu{
		Name:     faker.Name(),
		Category: faker.Username(),
	}
}

func generateEditMenuForm() *forms.EditMenu {
	return &forms.EditMenu{
		Name:     faker.Name(),
		Category: faker.Username(),
	}
}

func createAndAddMenu(menu *forms.Menu) *models.Menu {
	if menu == nil {
		menu = generateMenuForm()
	}

	men, err := menuSrv.Add(menu)
	if err != nil {
		panic(err)
	}

	return men
}
