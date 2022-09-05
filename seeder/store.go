package seeder

import (
	"database/sql"
	"log"
)

func MigrateStore(db *sql.DB) {

	sqlStatement := `
		drop table IF EXISTS store cascade;
	
		create table IF NOT EXISTS store (
			id serial PRIMARY KEY,
			name VARCHAR(150) NOT NULL
		);
		insert into store (name) values ('Livetube');
		insert into store (name) values ('Photobug');
		insert into store (name) values ('Fatz');
		insert into store (name) values ('Thoughtmix');
		insert into store (name) values ('Jabbertype');
		insert into store (name) values ('Viva');
		insert into store (name) values ('Pixope');
		insert into store (name) values ('Voomm');
		insert into store (name) values ('Babbleset');
		insert into store (name) values ( 'Skipstorm');
	`
	_, err := db.Exec(sqlStatement)

	if err != nil {
		log.Println("Error while Migrating Store:")
		panic(err)
	}
}
