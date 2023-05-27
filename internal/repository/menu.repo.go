package repo

import (
	"database/sql"
	"restaurant-management/internal/models"
)

type MenuRepo interface {
	Add(menu *models.Menu) (men *models.Menu, err error)
}

type menuRepo struct {
	conn *sql.DB
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

func NewMenuRepo(conn *sql.DB) MenuRepo {
	return &menuRepo{conn: conn}
}
