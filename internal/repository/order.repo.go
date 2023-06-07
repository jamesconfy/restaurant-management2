package repo

import (
	"database/sql"
	"restaurant-management/internal/models"
)

type OrderRepo interface {
	Add(order *models.Order) (orde *models.Order, err error)
	Get(orderId string) (order *models.Order, err error)
	GetAll() (orders []*models.Order, err error)
}

type orderSql struct {
	conn *sql.DB
}

// Get implements OrderRepo.
func (o *orderSql) Get(orderId string) (order *models.Order, err error) {
	order = new(models.Order)

	query := `SELECT id, table_id, payment_id, delivery_id, date_created, date_updated FROM orders WHERE id = $1`

	err = o.conn.QueryRow(query, orderId).Scan(&order.Id, &order.TableId, &order.PaymentId, &order.DeliveryId, &order.DateCreated, &order.DateUpdated)
	if err != nil {
		return
	}

	return
}

// GetAll implements OrderRepo.
func (o *orderSql) GetAll() (orders []*models.Order, err error) {
	query := `SELECT id, table_id, payment_id, delivery_id, date_created, date_updated FROM orders`

	rows, err := o.conn.Query(query)
	if err != nil {
		return
	}

	for rows.Next() {
		var order models.Order

		err = rows.Scan(&order.Id, &order.TableId, &order.PaymentId, &order.DeliveryId, &order.DateCreated, &order.DateUpdated)
		if err != nil {
			return
		}

		orders = append(orders, &order)
	}

	return
}

func (o *orderSql) Add(order *models.Order) (orde *models.Order, err error) {
	orde = new(models.Order)

	query := `INSERT INTO orders(table_id) VALUES($1) RETURNING id, table_id, payment_id, delivery_id, date_created, date_updated`

	err = o.conn.QueryRow(query, order.TableId).Scan(&orde.Id, &orde.TableId, &orde.PaymentId, &orde.DeliveryId, &orde.DateCreated, &orde.DateUpdated)
	if err != nil {
		return
	}

	return
}

func NewOrderRepo(conn *sql.DB) OrderRepo {
	return &orderSql{conn: conn}
}
