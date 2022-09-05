package seeder

import (
	"database/sql"
	"log"
)

func MigrateUser(db *sql.DB) {
	sqlStatement := `
	
		drop table IF EXISTS user_data cascade;
		create table IF NOT EXISTS user_data (
			id serial PRIMARY KEY,
			full_name VARCHAR(150) NOT NULL,
			phone_num VARCHAR(13) NOT NULL,
			email VARCHAR(100) NOT NULL,
			referal_code VARCHAR(100)
		);
		insert into user_data (full_name, phone_num, referal_code, email) values ('Sig Wisdom', '4071675964', 6924, 'swisdom0@usatoday.com');
		insert into user_data (full_name, phone_num, referal_code, email) values ('Ripley Yerlett', '5668163093', 1223, 'ryerlett1@com.com');
		insert into user_data (full_name, phone_num, referal_code, email) values ('Tanner Woolner', '6969464843', 5316, 'twoolner2@stumbleupon.com');
		insert into user_data (full_name, phone_num, referal_code, email) values ('Jim Gwyneth', '5874603472', 6060, 'jgwyneth3@ibm.com');
		insert into user_data (full_name, phone_num, referal_code, email) values ('Darnell Friedman', '6284241088', 6101, 'dfriedman4@topsy.com');
		insert into user_data (full_name, phone_num, referal_code, email) values ('Deni Flowers', '6695709987', 4641, 'dflowers5@adobe.com');
		insert into user_data (full_name, phone_num, referal_code, email) values ('Leonard Matiasek', '5932811188', 7260, 'lmatiasek6@ibm.com');
		insert into user_data (full_name, phone_num, referal_code, email) values ('Guinna Misson', '6846128141', 8938, 'gmisson7@sitemeter.com');
		insert into user_data (full_name, phone_num, referal_code, email) values ('Atlanta Asmus', '6378349674', 2019, 'aasmus8@yahoo.co.jp');
		insert into user_data (full_name, phone_num, referal_code, email) values ('Linus Hendrikse', '9016918269', 7174, 'lhendrikse9@google.ca');
		insert into user_data (full_name, phone_num, referal_code, email) values ('Arief', '9016918269', 7174, 'popo@google.ca');
	`
	_, err := db.Exec(sqlStatement)

	if err != nil {
		log.Println("Error while Migrating User")
		panic(err)
	}
}
