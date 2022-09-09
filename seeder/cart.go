package seeder

import (
	"database/sql"
	"log"
)

func MigrateCart(db *sql.DB) {
	sqlStatement := `
	
		drop table IF EXISTS cart cascade;
		create table IF NOT EXISTS cart (
			id serial PRIMARY KEY,
			user_id INT NOT NULL,
			delivery_method_id INT NOT NULL,
			note VARCHAR(200),
			FOREIGN KEY (user_id) REFERENCES user_data (id) on delete set null,
			FOREIGN KEY (delivery_method_id) REFERENCES delivery_method (id) on delete set null
		);
		insert into cart (user_id, delivery_method_id, note) values (2, 1, 'Occup of bus injured in collision w hv veh in traf, sequela');
		insert into cart (user_id, delivery_method_id, note) values (7, 5, 'Calcific tendinitis of unspecified shoulder');
		insert into cart (user_id, delivery_method_id, note) values (4, 4, 'Displ commnt fx shaft of rad, l arm, 7thN');
		insert into cart (user_id, delivery_method_id, note) values (3, 5, 'Nondisp fx of head of unsp rad, 7thQ');
		insert into cart (user_id, delivery_method_id, note) values (11, 3, 'Toxic effect of smoke, undetermined, subsequent encounter');
	`
	_, err := db.Exec(sqlStatement)

	if err != nil {
		log.Println("Error while Migrating Cart")
		panic(err)
	}
}
