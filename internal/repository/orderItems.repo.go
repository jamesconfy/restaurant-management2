package repo

import (
	"database/sql"
	"restaurant-management/internal/models"
)

type OrderItemRepo interface {
	Add(item *models.OrderItem) (*models.OrderItem, error)
	Get(orderItemId string) (*models.OrderItem, error)
	GetAll() ([]*models.OrderItem, error)
	Delete(orderItemId string) error
}

type orderItemRepo struct {
	conn *sql.DB
}

// Add implements OrderItemRepo.
func (oi *orderItemRepo) Add(item *models.OrderItem) (ite *models.OrderItem, err error) {
	ite = new(models.OrderItem)

	query := `INSERT INTO orderitems(quantity, order_id, food_id) VALUES ($1, $2, $3) ON CONFLICT (food_id, order_id)
	DO UPDATE SET quantity = $1, date_updated = CURRENT_TIMESTAMP RETURNING id, quantity, order_id, food_id, date_created, date_updated`

	err = oi.conn.QueryRow(query, item.Quantity, item.OrderId, item.FoodId).Scan(&ite.Id, &ite.Quantity, &ite.OrderId, &ite.FoodId, &ite.DateCreated, &ite.DateUpdated)
	if err != nil {
		return
	}

	return
}

// GetAll implements OrderItemRepo.
func (oi *orderItemRepo) GetAll() (items []*models.OrderItem, err error) {

	query := `SELECT id, quantity, order_id, food_id, date_created, date_updated FROM orderitems`

	rows, err := oi.conn.Query(query)
	if err != nil {
		return
	}

	for rows.Next() {
		var item models.OrderItem

		err = rows.Scan(&item.Id, &item.Quantity, &item.OrderId, &item.FoodId, &item.DateCreated, &item.DateUpdated)
		if err != nil {
			return
		}

		items = append(items, &item)
	}

	return
}

// Get implements OrderItemRepo.
func (oi *orderItemRepo) Get(orderItemId string) (item *models.OrderItem, err error) {
	item = new(models.OrderItem)

	query := `SELECT id, quantity, order_id, food_id, date_created, date_updated FROM orderitems WHERE id = $1`

	err = oi.conn.QueryRow(query, orderItemId).Scan(&item.Id, &item.Quantity, &item.OrderId, &item.FoodId, &item.DateCreated, &item.DateUpdated)
	if err != nil {
		return
	}

	return
}

// Delete implements OrderItemRepo.
func (oi *orderItemRepo) Delete(orderItemId string) error {
	query := `DELETE FROM orderitems WHERE id = $1`

	_, err := oi.conn.Exec(query, orderItemId)
	if err != nil {
		return err
	}

	return nil
}

func NewOrderItemRepo(conn *sql.DB) OrderItemRepo {
	return &orderItemRepo{conn: conn}
}

