package repo

import (
	"database/sql"
	"restaurant-management/internal/models"
)

type FoodRepo interface {
	Add(food *models.Food) (foo *models.Food, err error)
	Get(foodId string) (foo *models.Food, err error)
	GetAll() (foods []*models.Food, err error)
	Edit(food *models.Food) (foo *models.Food, err error)
}

type foodRepo struct {
	conn *sql.DB
}

func (f *foodRepo) Add(food *models.Food) (foo *models.Food, err error) {
	foo = new(models.Food)

	query := `INSERT INTO food(name, price, image, menu_id) VALUES ($1, $2, $3, $4) RETURNING id, name, price, image, menu_id, date_created, date_updated`

	err = f.conn.QueryRow(query, food.Name, food.Price, food.Image, food.MenuId).Scan(&foo.Id, &foo.Name, &foo.Price, &foo.Image, &foo.MenuId, &foo.DateCreated, &foo.DateUpdated)
	if err != nil {
		return
	}

	return
}

func (f *foodRepo) Get(foodId string) (foo *models.Food, err error) {
	foo = new(models.Food)

	query := `SELECT id, name, price, image, menu_id, date_created, date_updated FROM food WHERE id = $1`

	err = f.conn.QueryRow(query, foodId).Scan(&foo.Id, &foo.Name, &foo.Price, &foo.Image, &foo.MenuId, &foo.DateCreated, &foo.DateUpdated)
	if err != nil {
		return
	}

	return
}

func (f *foodRepo) GetAll() (foods []*models.Food, err error) {
	query := `SELECT id, name, price, image, menu_id, date_created, date_updated FROM food`

	rows, err := f.conn.Query(query)
	if err != nil {
		return
	}

	for rows.Next() {
		var food models.Food

		err = rows.Scan(&food.Id, &food.Name, &food.Price, &food.Image, &food.MenuId, &food.DateCreated, &food.DateUpdated)
		if err != nil {
			return
		}

		foods = append(foods, &food)
	}

	return
}

func (f *foodRepo) Edit(food *models.Food) (foo *models.Food, err error) {
	foo = new(models.Food)

	query := `UPDATE food SET name = $1, price = $2, image = $3, menu_id = $4 RETURNING id, name, price, image, menu_id, date_created, date_updated`

	err = f.conn.QueryRow(query, food.Name, food.Price, food.Image, food.MenuId).Scan(&foo.Id, &foo.Name, &foo.Price, &foo.Image, &foo.MenuId, &foo.DateCreated, &foo.DateUpdated)
	if err != nil {
		return
	}

	return
}

func NewFoodRepo(conn *sql.DB) FoodRepo {
	return &foodRepo{conn: conn}
}
