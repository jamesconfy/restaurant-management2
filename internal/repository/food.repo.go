package repo

import (
	"database/sql"
	"restaurant-management/internal/models"
)

type FoodRepo interface {
	FoodExists(foodId string) (bool, error)

	Add(food *models.Food) (foo *models.Food, err error)
	Get(foodId string) (foo *models.Food, err error)
	GetAll() (foods []*models.Food, err error)
	Edit(foodId string, food *models.Food) (foo *models.Food, err error)
	Delete(foodId string) error
}

type foodRepo struct {
	conn *sql.DB
}

func (f *foodRepo) FoodExists(foodId string) (bool, error) {
	var name string

	query := `SELECT name FROM food WHERE id = $1`

	err := f.conn.QueryRow(query, foodId).Scan(&name)
	if err != nil {
		if err == sql.ErrNoRows {
			// Food does not exist
			return false, nil
		}
		// An error occurred while executing the query
		return true, err
	}

	// Food already exists
	return true, nil
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

func (f *foodRepo) Edit(foodId string, food *models.Food) (foo *models.Food, err error) {
	foo = new(models.Food)

	query := `UPDATE food SET name = $1, price = $2, image = $3, menu_id = $4, date_updated = CURRENT_TIMESTAMP WHERE id = $5 RETURNING id, name, price, image, menu_id, date_created, date_updated`

	err = f.conn.QueryRow(query, food.Name, food.Price, food.Image, food.MenuId, foodId).Scan(&foo.Id, &foo.Name, &foo.Price, &foo.Image, &foo.MenuId, &foo.DateCreated, &foo.DateUpdated)
	if err != nil {
		return
	}

	return
}

// Delete implements FoodRepo
func (f *foodRepo) Delete(foodId string) error {
	query := `DELETE FROM food WHERE id = $1`

	_, err := f.conn.Exec(query, foodId)
	if err != nil {
		return err
	}

	return nil
}

func NewFoodRepo(conn *sql.DB) FoodRepo {
	return &foodRepo{conn: conn}
}
