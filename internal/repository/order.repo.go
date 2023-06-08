package repo

import (
	"database/sql"
	"restaurant-management/internal/models"
)

type OrderRepo interface {
	Add(order *models.Order) (orde *models.Order, err error)
	Get(orderId string) (order *models.Order, err error)
	GetAll() (orders []*models.Order, err error)
	Edit(orderId string, order *models.Order) (*models.Order, error)
	Delete(orderId string) (err error)
}

type orderSql struct {
	conn *sql.DB
}

// Delete implements OrderRepo.
func (o *orderSql) Delete(orderId string) (err error) {
	query := `DELETE FROM orders WHERE id = $1`

	_, err = o.conn.Exec(query, orderId)
	if err != nil {
		return
	}

	return
}

// Edit implements OrderRepo.
func (o *orderSql) Edit(orderId string, order *models.Order) (orde *models.Order, err error) {
	orde = new(models.Order)

	query := `UPDATE orders SET table_id = $1, delivery_id = $2, payment_id = $3, date_updated = CURRENT_TIMESTAMP WHERE id = $4 RETURNING id, table_id, delivery_id, payment_id, date_created, date_updated`

	err = o.conn.QueryRow(query, order.TableId, order.DeliveryId, order.PaymentId, orderId).Scan(&orde.Id, &orde.TableId, &orde.DeliveryId, &orde.PaymentId, &orde.DateCreated, &orde.DateUpdated)
	if err != nil {
		return
	}

	return o.Get(orde.Id)
}

// Get implements OrderRepo.
func (o *orderSql) Get(orderId string) (order *models.Order, err error) {
	order = new(models.Order)

	query := `SELECT o.id, o.table_id, pm.id, d.id, pm.payment_type, d.status, o.date_created, o.date_updated FROM orders o JOIN delivery d ON o.delivery_id = d.id JOIN payment_method pm ON o.payment_id = pm.id WHERE o.id = $1`

	err = o.conn.QueryRow(query, orderId).Scan(&order.Id, &order.TableId, &order.PaymentId, &order.DeliveryId, &order.PaymentMethod, &order.DeliveryStatus, &order.DateCreated, &order.DateUpdated)
	if err != nil {
		return
	}

	return
}

// GetAll implements OrderRepo.
func (o *orderSql) GetAll() (orders []*models.Order, err error) {
	query := `SELECT o.id, o.table_id, pm.payment_type, d.status, o.date_created, o.date_updated FROM orders o JOIN delivery d ON o.delivery_id = d.id JOIN payment_method pm ON o.payment_id = pm.id ORDER BY o.date_created`

	rows, err := o.conn.Query(query)
	if err != nil {
		return
	}

	for rows.Next() {
		var order models.Order

		err = rows.Scan(&order.Id, &order.TableId, &order.PaymentMethod, &order.DeliveryStatus, &order.DateCreated, &order.DateUpdated)
		if err != nil {
			return
		}

		orders = append(orders, &order)
	}

	return
}

func (o *orderSql) Add(order *models.Order) (orde *models.Order, err error) {
	var id string

	query := `INSERT INTO orders(table_id) VALUES($1) RETURNING id`

	err = o.conn.QueryRow(query, order.TableId).Scan(&id)
	if err != nil {
		return
	}

	return o.Get(id)
}

func NewOrderRepo(conn *sql.DB) OrderRepo {
	return &orderSql{conn: conn}
}
