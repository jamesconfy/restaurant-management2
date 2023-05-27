package repo_test

import (
	"math/rand"
	"restaurant-management/internal/models"

	"github.com/bxcodec/faker/v4"
)

func generateFood(menu *models.Menu) *models.Food {
	if menu == nil {
		menu = createAndAddMenu(nil)
	}

	return &models.Food{
		Name:   faker.Name(),
		Price:  rand.Float64() * 1000,
		Image:  faker.IPv4(),
		MenuId: menu.Id,
	}
}
