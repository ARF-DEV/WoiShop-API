package repository

import (
	"azura-lab-intern/study-case-1/models"
	"database/sql"
)

type CartRepository struct {
	db *sql.DB
}

func NewCartRepository(db *sql.DB) *CartRepository {
	return &CartRepository{
		db: db,
	}
}

func (c *CartRepository) GetCartByID(id int) (*models.Cart, error) {
	sqlStatement := `SELECT id, user_id, delivery_method_id, note FROM cart WHERE id = $1`

	var v models.Cart
	err := c.db.QueryRow(sqlStatement, id).Scan(&v.ID, &v.UserID, &v.DeliveryMethodID, &v.Note)

	if err != nil {
		return nil, err
	}

	return &v, nil
}

func (c *CartRepository) GetAllCart() ([]models.Cart, error) {

	sqlStatement := `SELECT id, user_id, delivery_method_id, note FROM cart`

	var res []models.Cart
	rows, err := c.db.Query(sqlStatement)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var cart models.Cart

		err = rows.Scan(&cart.ID, &cart.UserID, &cart.DeliveryMethodID, &cart.Note)
		if err != nil {
			return nil, err
		}

		res = append(res, cart)
	}

	return res, nil
}
