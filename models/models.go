package models

import (
	"gopkg.in/guregu/null.v3"
)

type Product struct {
	ID          int         `db:"id" json:"id"`
	CategoryID  int         `db:"category_id" json:"category_id,omitempty"`
	Name        string      `db:"name" json:"name"`
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
