package helpers

import "database/sql"

func InitDB(DB *sql.DB) {
	sqlStatement := `
		drop table IF EXISTS product;
		drop table IF EXISTS category;

		create table IF NOT EXISTS category (
			id serial PRIMARY KEY,
			name VARCHAR(100) NOT NULL
		);
		CREATE TABLE IF NOT EXISTS product (
			id serial PRIMARY KEY,
			category_id INT NOT NULL,
			name VARCHAR(150) NOT NULL,
			price INT NOT NULL,
			FOREIGN KEY (category_id) REFERENCES category (id) on delete set null
		);
	`

	_, err := DB.Exec(sqlStatement)

	if err != nil {
		panic(err)
	}

}
