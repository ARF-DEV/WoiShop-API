package models

import "database/sql"

type Product struct {
	ID          int            `db:"id" json:"id"`
	CategoryID  int            `db:"category_id" json:"category_id,omitempty"`
	Name        string         `db:"name" json:"name"`
	Price       int            `db:"price" json:"price"`
	Description sql.NullString `db:"description" json:"description,omitempty"`
}

type Category struct {
	ID   int    `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
}
