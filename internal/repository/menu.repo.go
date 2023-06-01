package repo

import (
	"database/sql"
	"restaurant-management/internal/models"
)

type MenuRepo interface {
	Check(name, category string) (bool, error)
	CheckMenuExists(menuId string) (bool, error)

	Add(menu *models.Menu) (men *models.Menu, err error)
	Get(menuId string) (men *models.MenuFood, err error)
	GetAll() (menu []*models.Menu, err error)
	Edit(menuId string, menu *models.Menu) (men *models.Menu, err error)
	Delete(menuId string) error
}

type menuRepo struct {
	conn *sql.DB
}

// CheckMenu implements MenuRepo
func (m *menuRepo) CheckMenuExists(menuId string) (bool, error) {
	var name string

	query := `SELECT name FROM menu WHERE id = $1`

	err := m.conn.QueryRow(query, menuId).Scan(&name)
	if err != nil {
		if err == sql.ErrNoRows {
			// Name and Category does not exist
			return false, nil
		}
		// An error occurred while executing the query
		return true, err
	}

	// Name and Category already exists
	return true, nil
}

func (m *menuRepo) Check(name, category string) (bool, error) {
	var id string

	query := `SELECT id FROM menu WHERE name = $1 AND category = $2`

	err := m.conn.QueryRow(query, name, category).Scan(&id)
	if err != nil {
		if err == sql.ErrNoRows {
			// Name and Category does not exist
			return false, nil
		}
		// An error occurred while executing the query
		return true, err
	}

	// Name and Category already exists
	return true, nil
}

func (m *menuRepo) Add(menu *models.Menu) (men *models.Menu, err error) {
	men = new(models.Menu)

	query := `INSERT INTO menu(name, category) VALUES($1, $2) RETURNING id, name, category, date_created, date_updated`

	err = m.conn.QueryRow(query, menu.Name, menu.Category).Scan(&men.Id, &men.Name, &men.Category, &men.DateCreated, &men.DateUpdated)
	if err != nil {
		return
	}

	return
}

func (m *menuRepo) get(menuId string) (menu *models.Menu, err error) {
	menu = new(models.Menu)

	query := `SELECT id, name, category, date_created, date_updated FROM menu WHERE id = $1`

	err = m.conn.QueryRow(query, menuId).Scan(&menu.Id, &menu.Name, &menu.Category, &menu.DateCreated, &menu.DateUpdated)
	if err != nil {
		return
	}

	return
}

func (m *menuRepo) Get(menuId string) (*models.MenuFood, error) {
	var menufood models.MenuFood

	menu, err := m.get(menuId)
	if err != nil {
		return nil, err
	}

	menufood.Menu = menu

	query := `SELECT f.id, f.name, f.price, f.image, f.date_created, f.date_updated FROM food f WHERE f.menu_id = $1;`

	rows, err := m.conn.Query(query, menuId)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var food models.Food

		err := rows.Scan(&food.Id, &food.Name, &food.Price, &food.Image, &food.DateCreated, &food.DateUpdated)
		if err != nil {
			return nil, err
		}

		menufood.Foods = append(menufood.Foods, &food)
	}

	return &menufood, nil
}

func (m *menuRepo) GetAll() (menu []*models.Menu, err error) {
	query := `SELECT id, name, category, date_created, date_updated FROM menu`

	rows, err := m.conn.Query(query)
	if err != nil {
		return
	}

	for rows.Next() {
		var men models.Menu

		err = rows.Scan(&men.Id, &men.Name, &men.Category, &men.DateCreated, &men.DateUpdated)
		if err != nil {
			return
		}

		menu = append(menu, &men)
	}

	return
}

func (m *menuRepo) Edit(menuId string, menu *models.Menu) (men *models.Menu, err error) {
	men = new(models.Menu)

	query := `UPDATE menu SET name = $1, category = $2, date_updated = CURRENT_TIMESTAMP WHERE id = $3 RETURNING id, name, category, date_created, date_updated`

	err = m.conn.QueryRow(query, menu.Name, menu.Category, menuId).Scan(&men.Id, &men.Name, &men.Category, &men.DateCreated, &men.DateUpdated)
	if err != nil {
		return
	}

	return
}

func (m *menuRepo) Delete(menuId string) error {
	query := `DELETE FROM menu WHERE id = $1`

	_, err := m.conn.Exec(query, menuId)
	if err != nil {
		return err
	}

	return nil
}

func NewMenuRepo(conn *sql.DB) MenuRepo {
	return &menuRepo{conn: conn}
}
