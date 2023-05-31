package service_test

import (
	"math/rand"
	"restaurant-management/internal/forms"
	"restaurant-management/internal/models"

	"github.com/bxcodec/faker/v4"
)

func generateFoodForm() *forms.Food {
	return &forms.Food{
		Name:  faker.Name(),
		Price: rand.Float64() * 100,
		Image: faker.IPv4(),
	}
}

func createAndAddFood(menu *models.Menu, food *forms.Food) *models.Food {
	if menu == nil {
		menu = createAndAddMenu(nil)
	}

	if food == nil {
		food = generateFoodForm()
	}

	foo, err := foodSrv.Add(menu.Id, food)
	if err != nil {
		panic(err)
	}

	return foo
}
