package seeder

import (
	"database/sql"
	"log"
)

func MigrateDeliveryMethod(db *sql.DB) {
	sqlStatement := `
		drop table IF EXISTS delivery_method cascade;
		create table IF NOT EXISTS delivery_method (
			id serial PRIMARY KEY,
			name VARCHAR(100)
		);


		insert into delivery_method (name) values 
		('Pickup'),
		('Express'),
		('09.00 - 11.00'),
		('12.00 - 16.00'),
		('18.00 - 20.00');
	`

	_, err := db.Exec(sqlStatement)

	if err != nil {
		log.Println("Error while Migrating Delivery Method")
		panic(err)
	}
}
