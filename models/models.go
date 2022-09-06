package models

import (
	"gopkg.in/guregu/null.v3"
)

type Product struct {
	ID          int         `db:"id" json:"id"`
	StoreID     int         `db:"store_id" json:"store_id"`
	CategoryID  int         `db:"category_id" json:"category_id,omitempty"`
	Name        string      `db:"name" json:"name"`
	Stock       int         `db:"stock" json:"stock"`
	Price       int         `db:"price" json:"price"`
	Description null.String `db:"description" json:"description,omitempty"`
	ImageLink   null.String `db:"image_link" json:"image_link"`
}

type Category struct {
	ID   int    `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
}

type User struct {
	ID          int      `db:"id" json:"id"`
	Name        string   `db:"full_name" json:"full_name"`
	NoTelp      string   `db:"phone_num" json:"phone_num"`
	Email       string   `db:"email" json:"email"`
	ReferalCode null.Int `db:"referal_code" json:"referal_code"`
}

type Order struct {
	ID        int `db:"id" json:"id"`
	CartID    int `db:"cart_id" json:"cart_id,omitempty"`
	ProductID int `db:"product_id" json:"product_id,omitempty"`
	Amount    int `db:"amount" json:"amount"`
}

type Cart struct {
	ID               int         `db:"id" json:"id"`
	UserID           int         `db:"user_id" json:"user_id"`
	DeliveryMethodID int         `db:"delivery_method_id" json:"delivery_method_id"`
	Note             null.String `db:"note" json:"note"`
}

type Store struct {
	ID   int `db:"id" json:"id"`
	Name int `db:"name" json:"name"`
}

type DeliveryMethod struct {
	ID   int `db:"id" json:"id"`
	Name int `db:"name" json:"name"`
}
