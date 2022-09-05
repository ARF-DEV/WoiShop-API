package seeder

import "database/sql"

func MigrateOrder(db *sql.DB) {

	sqlStatement := `
		drop table IF EXISTS order_data cascade;
		
		create table IF NOT EXISTS order_data (
			id serial PRIMARY KEY,
			cart_id INT NOT NULL,
			product_id INT NOT NULL,
			amount INT NOT NULL,
			FOREIGN KEY (cart_id) REFERENCES cart (id) on delete set null,
			FOREIGN KEY (product_id) REFERENCES product (id) on delete set null
		);

		insert into order_data (cart_id, product_id, amount) values (3, 4, 3);
		insert into order_data (cart_id, product_id, amount) values (5, 1, 8);
		insert into order_data (cart_id, product_id, amount) values (2, 3, 7);
		insert into order_data (cart_id, product_id, amount) values (2, 1, 1);
		insert into order_data (cart_id, product_id, amount) values (5, 3, 4);
		insert into order_data (cart_id, product_id, amount) values (3, 4, 5);
		insert into order_data (cart_id, product_id, amount) values (4, 1, 5);
		insert into order_data (cart_id, product_id, amount) values (5, 5, 8);
		insert into order_data (cart_id, product_id, amount) values (3, 2, 9);
		insert into order_data (cart_id, product_id, amount) values (1, 4, 10);	
		`

	_, err := db.Exec(sqlStatement)

	if err != nil {
		panic(err)
	}

}
