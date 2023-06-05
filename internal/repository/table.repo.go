package repo

import (
	"database/sql"
	"restaurant-management/internal/models"
)

type TableRepo interface {
	TableExists(tableId string) (bool, error)

	Add(table *models.Table) (*models.Table, error)
	Get(id string) (tabl *models.Table, err error)
	GetAll(role string) (tables []*models.Table, err error)
	Edit(tableId string, table *models.Table) (tabl *models.Table, err error)
	Delete(id string) (err error)
}

type tableSql struct {
	conn *sql.DB
}

func (t *tableSql) TableExists(tableId string) (bool, error) {
	var seats string

	query := `SELECT seats FROM tables WHERE id = $1`

	err := t.conn.QueryRow(query, tableId).Scan(&seats)
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

func (t *tableSql) Add(table *models.Table) (tabl *models.Table, err error) {
	tabl = new(models.Table)

	query := `INSERT INTO tables(seats) VALUES($1) RETURNING id, seats, number, booked, date_created, date_updated`

	err = t.conn.QueryRow(query, table.Seats).Scan(&tabl.Id, &tabl.Seats, &tabl.Number, &tabl.Booked, &tabl.DateCreated, &tabl.DateUpdated)
	if err != nil {
		return
	}

	return
}

func (t *tableSql) Get(id string) (tabl *models.Table, err error) {
	tabl = new(models.Table)

	query := `SELECT id, seats, number, booked, date_created, date_updated FROM tables WHERE id = $1`

	err = t.conn.QueryRow(query, id).Scan(&tabl.Id, &tabl.Seats, &tabl.Number, &tabl.Booked, &tabl.DateCreated, &tabl.DateUpdated)
	if err != nil {
		return
	}

	return
}

func (t *tableSql) GetAll(role string) (tables []*models.Table, err error) {
	query := `SELECT id, seats, number, booked, date_created, date_updated FROM tables`

	if role == "USER" {
		query = `SELECT id, seats, number, booked, date_created, date_updated FROM tables WHERE booked = false`
	}

	rows, err := t.conn.Query(query)
	if err != nil {
		return
	}

	for rows.Next() {
		var table models.Table

		err = rows.Scan(&table.Id, &table.Seats, &table.Number, &table.Booked, &table.DateCreated, &table.DateUpdated)
		if err != nil {
			return nil, err
		}

		tables = append(tables, &table)
	}

	return
}

func (t *tableSql) Edit(tableId string, table *models.Table) (tabl *models.Table, err error) {
	tabl = new(models.Table)

	query := `UPDATE tables SET seats = $1, booked = $2, date_updated = CURRENT_TIMESTAMP WHERE id = $3 RETURNING id, seats, number, booked, date_created, date_updated`

	err = t.conn.QueryRow(query, table.Seats, table.Booked, tableId).Scan(&tabl.Id, &tabl.Seats, &tabl.Number, &tabl.Booked, &tabl.DateCreated, &tabl.DateUpdated)
	if err != nil {
		return
	}

	return
}

func (t *tableSql) Delete(id string) (err error) {
	query := `DELETE FROM tables WHERE id = $1`

	_, err = t.conn.Exec(query, id)
	if err != nil {
		return
	}

	return
}

func NewTableRepo(conn *sql.DB) TableRepo {
	return &tableSql{conn: conn}
}
