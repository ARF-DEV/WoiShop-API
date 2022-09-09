package repository

import (
	"azura-lab-intern/study-case-1/models"
	"database/sql"
)

type OrderRepository struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{
		db: db,
	}
}

func (o *OrderRepository) GetOrderByCartID(CartID int) ([]models.Order, error) {
	sqlStatement := `SELECT id, product_id, amount FROM order_data WHERE cart_id = $1`

	rows, err := o.db.Query(sqlStatement, CartID)

	if err != nil {
		return nil, err
	}

	var orders []models.Order

	for rows.Next() {
		var o models.Order

		err = rows.Scan(&o.ID, &o.ProductID, &o.Amount)
		if err != nil {
			return nil, err
		}

		orders = append(orders, o)
	}

	return orders, nil
}

func (c *OrderRepository) ChangeOrderAmount(orderID int, newAmount int) (*models.Order, error) {
	sqlStatement := "UPDATE order_data SET amount = $1 WHERE id = $2 RETURNING id, cart_id, product_id, amount;"

	var updatedOrder models.Order
	err := c.db.QueryRow(sqlStatement, newAmount, orderID).
		Scan(&updatedOrder.ID, &updatedOrder.CartID, &updatedOrder.ProductID, &updatedOrder.Amount)

	if err != nil {
		return nil, err
	}
	return &updatedOrder, nil
}

func (c *OrderRepository) CreateOrder(order models.Order) (*models.Order, error) {
	sqlStatement := "INSERT INTO order_data (cart_id, product_id, amount) VALUES ($1, $2, $3) RETURNING id, cart_id, product_id, amount;"

	var o models.Order

	err := c.db.QueryRow(sqlStatement, order.CartID, order.ProductID, order.Amount).
		Scan(&o.ID, &o.CartID, &o.ProductID, &o.Amount)

	if err != nil {
		return nil, err
	}

	return &o, nil
}

func (c *OrderRepository) DeleteOrderByID(id int) (*models.Order, error) {
	sqlStatement := `DELETE FROM order_data WHERE id = $1 RETURNING id, cart_id, product_id, amount`

	var deletedRow models.Order
	err := c.db.QueryRow(sqlStatement, id).Scan(&deletedRow.ID, &deletedRow.CartID, deletedRow.ProductID, deletedRow.Amount)

	if err != nil {
		return nil, err
	}

	return &deletedRow, nil
}
