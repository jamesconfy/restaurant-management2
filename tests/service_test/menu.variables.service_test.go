package service_test

import (
	"restaurant-management/internal/forms"

	"github.com/bxcodec/faker/v4"
)

func generateMenuForm() *forms.Menu {
	return &forms.Menu{
		Name:     faker.Name(),
		Category: faker.Username(),
	}
}
