package repo_test

import (
	"restaurant-management/internal/models"

	"github.com/bxcodec/faker/v4"
)

func generateMenu() *models.Menu {
	return &models.Menu{
		Name:     faker.Name(),
		Category: faker.Timezone(),
	}
}

func createAndAddMenu(menu *models.Menu) *models.Menu {
	if menu == nil {
		menu = generateMenu()
	}

	men, err := menuRepo.Add(menu)
	if err != nil {
		panic(err)
	}

	return men
}
