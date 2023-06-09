package repo

import (
	"database/sql"
	"restaurant-management/internal/models"
)

type OrderItemRepo interface {
	Add(item *models.OrderItem) (*models.OrderItem, error)
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

func NewOrderItemRepo(conn *sql.DB) OrderItemRepo {
	return &orderItemRepo{conn: conn}
}
