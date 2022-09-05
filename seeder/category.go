package seeder

import (
	"database/sql"
	"log"
)

func MigrateCategory(db *sql.DB) {
	sqlStatement := `
	
		drop table IF EXISTS category cascade;

		create table IF NOT EXISTS category (
			id serial PRIMARY KEY,
			name VARCHAR(100) NOT NULL
		);
		insert into category (name)	values
		('pakaian'),
		('makanan'),
		('minuman');
	`
	_, err := db.Exec(sqlStatement)

	if err != nil {
		log.Println("Error while Migrating Category")
		panic(err)
	}
}
