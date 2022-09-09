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

func (c *CartRepository) GetCartByID(ID int) (*models.Cart, error) {
	sqlStatement := `SELECT id, user_id, delivery_method_id, note FROM cart WHERE id = $1`

	var v models.Cart
	err := c.db.QueryRow(sqlStatement, ID).Scan(&v.ID, &v.UserID, &v.DeliveryMethodID, &v.Note)

	if err != nil {
		return nil, err
	}

	return &v, nil
}
func (c *CartRepository) GetCartByUserID(userID int) (*models.Cart, error) {
	sqlStatement := `SELECT id, user_id, delivery_method_id, note FROM cart WHERE user_id = $1`

	var v models.Cart
	err := c.db.QueryRow(sqlStatement, userID).Scan(&v.ID, &v.UserID, &v.DeliveryMethodID, &v.Note)

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

func (c *CartRepository) DeleteCartByID(id int) error {
	sqlStatement := `DELETE FROM cart WHERE id = $1 RETURNING id`

	var deletedID int
	err := c.db.QueryRow(sqlStatement, id).Scan(&deletedID)

	if err != nil {
		return err
	}

	return nil
}

func (c *CartRepository) AddCart(cart models.Cart) (*models.Cart, error) {
	sqlStatement := "INSERT INTO cart (user_id, delivery_method_id, note) VALUES ($1, $2, $3) RETURNING id,user_id, delivery_method_id, note"

	var res models.Cart
	err := c.db.QueryRow(sqlStatement, cart.UserID, cart.DeliveryMethodID, cart.Note).
		Scan(&res.ID, &res.UserID, &res.DeliveryMethodID, &res.Note)

	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *CartRepository) UpdateCart(cart models.Cart) error {
	sqlStatement := `
	UPDATE cart
	SET delivery_method_id = &1,
	SET note = $2
	WHERE id = $3;
	`

	_, err := c.db.Exec(sqlStatement, cart.DeliveryMethodID, cart.Note, cart.ID)

	if err != nil {
		return err
	}

	return nil
}
