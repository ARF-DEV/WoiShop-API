package models

type Product struct {
	ID         int    `db:"id"`
	CategoryID int    `db:"category_id"`
	Name       string `db:"name"`
	Price      int    `db:"price"`
}

type Category struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}
