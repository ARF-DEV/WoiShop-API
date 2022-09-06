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
