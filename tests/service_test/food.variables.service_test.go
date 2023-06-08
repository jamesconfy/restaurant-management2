package service_test

import (
	"math/rand"
	"restaurant-management/internal/forms"
	"restaurant-management/internal/models"

	"github.com/bxcodec/faker/v4"
)

func generateFoodForm(menu *models.Menu) *forms.Food {
	if menu == nil {
		menu = createAndAddMenu(nil)
	}

	return &forms.Food{
		Name:   faker.Name(),
		Price:  rand.Float64() * 100,
		Image:  faker.IPv4(),
		MenuId: menu.Id,
	}
}

func generateEditFoodForm(menu *models.Menu) *forms.EditFood {
	if menu == nil {
		menu = createAndAddMenu(nil)
	}

	return &forms.EditFood{
		Name:   faker.Name(),
		Price:  rand.Float64() * 100,
		Image:  faker.IPv4(),
		MenuId: menu.Id,
	}
}

func createAndAddFood(menu *models.Menu, food *forms.Food) *models.Food {
	if food == nil {
		food = generateFoodForm(menu)
	}

	foo, err := foodSrv.Add(food)
	if err != nil {
		panic(err)
	}

	return foo
}
