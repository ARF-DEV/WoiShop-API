package helpers

import "database/sql"

func InitDB(DB *sql.DB) {
	sqlStatement := `
		drop table IF EXISTS product;
		drop table IF EXISTS category;
		drop table IF EXISTS user_data;

		create table IF NOT EXISTS user_data (
			id serial PRIMARY KEY,
			full_name VARCHAR(150) NOT NULL,
			phone_num VARCHAR(13) NOT NULL,
			email VARCHAR(100) NOT NULL,
			referal_code VARCHAR(100)
		);
		create table IF NOT EXISTS category (
			id serial PRIMARY KEY,
			name VARCHAR(100) NOT NULL
		);
		CREATE TABLE IF NOT EXISTS product (
			id serial PRIMARY KEY,
			category_id INT NOT NULL,
			name VARCHAR(150) NOT NULL,
			price INT NOT NULL,
			description VARCHAR(500),
			image_link VARCHAR(200),
			FOREIGN KEY (category_id) REFERENCES category (id) on delete set null
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
		insert into user_data (full_name, phone_num, referal_code, email) values ('Arief', '6285351411242', 4949, 'ariefuddinsatriadharma@gmail.com');
			
		insert into category (name)	values
		('pakaian'),
		('makanan'),
		('minuman');
		
		insert into product (category_id, name, price, description) values (3, 'Sugar - Fine', 9674, null);
		insert into product (category_id, name, price, description) values (3, 'Potatoes - Pei 10 Oz', 1159, 'Pressure ulcer of unspecified part of back, stage 1');
		insert into product (category_id, name, price, description) values (3, 'Bread - Roll, Soft White Round', 3034, 'Puncture wound with foreign body of right back wall of thorax without penetration into thoracic cavity, subsequent encounter');
		insert into product (category_id, name, price, description) values (2, 'Cake - Dulce De Leche', 4233, 'Acute lymphadenitis of lower limb');
		insert into product (category_id, name, price, description) values (2, 'Bar Energy Chocchip', 4392, 'Puncture wound without foreign body, left ankle, sequela');
		insert into product (category_id, name, price, description) values (1, 'Sauce - Vodka Blush', 1694, null);
		insert into product (category_id, name, price, description) values (1, 'Asparagus - Green, Fresh', 1815, 'Nondisplaced fracture of proximal phalanx of left lesser toe(s), subsequent encounter for fracture with routine healing');
		insert into product (category_id, name, price, description) values (2, 'Mcgillicuddy Vanilla Schnap', 3163, null);
		insert into product (category_id, name, price, description) values (2, 'Beef - Rouladin, Sliced', 5691, null);
		insert into product (category_id, name, price, description) values (2, 'Sobe - Orange Carrot', 1521, 'Corrosions involving 30-39% of body surface');
		insert into product (category_id, name, price, description) values (1, 'Lamb - Shoulder', 8640, null);
		insert into product (category_id, name, price, description) values (1, 'Island Oasis - Peach Daiquiri', 9614, 'Flat foot [pes planus] (acquired), right foot');
		insert into product (category_id, name, price, description) values (2, 'Tea - Decaf Lipton', 2443, null);
		insert into product (category_id, name, price, description) values (1, 'Wine - Rosso Toscano Igt', 6651, 'Dislocation of T6/T7 thoracic vertebra, sequela');
		insert into product (category_id, name, price, description) values (1, 'Pernod', 6027, null);
		insert into product (category_id, name, price, description) values (1, 'Rice Paper', 8584, null);
		insert into product (category_id, name, price, description) values (3, 'Soup - Campbells Beef Stew', 7530, 'Legal intervention involving unspecified gas, suspect injured, initial encounter');
		insert into product (category_id, name, price, description) values (3, 'Wine - Red, Marechal Foch', 3991, 'Nondisplaced fracture of left radial styloid process, subsequent encounter for closed fracture with routine healing');
		insert into product (category_id, name, price, description) values (3, 'Foil Wrap', 1753, 'Greenstick fracture of shaft of right ulna, subsequent encounter for fracture with routine healing');
		insert into product (category_id, name, price, description) values (2, 'Wine - White, Pinot Grigio', 1857, 'Other specified injury of muscle and tendon of long extensor muscle of toe at ankle and foot level, left foot, subsequent encounter');
		insert into product (category_id, name, price, description) values (1, 'Lettuce - Frisee', 5128, 'Fall on board other unpowered watercraft, sequela');
		insert into product (category_id, name, price, description) values (3, 'Muffin Batt - Ban Dream Zero', 5796, null);
		insert into product (category_id, name, price, description) values (3, 'Juice - Grapefruit, 341 Ml', 5543, null);
		insert into product (category_id, name, price, description) values (1, 'Pail - 15l White, With Handle', 5980, null);
		insert into product (category_id, name, price, description) values (1, 'Bread - Italian Sesame Poly', 2231, 'Nondisplaced fracture of medial malleolus of unspecified tibia, subsequent encounter for closed fracture with delayed healing');
		insert into product (category_id, name, price, description) values (2, 'Vinegar - Balsamic', 4364, null);
		insert into product (category_id, name, price, description) values (1, 'Coconut Milk - Unsweetened', 2851, null);
		insert into product (category_id, name, price, description) values (1, 'Plaintain', 7948, 'Chronic inflammation of postmastoidectomy cavity');
		insert into product (category_id, name, price, description) values (2, 'Cheese - Cheddarsliced', 9522, 'Displaced fracture of anterior column [iliopubic] of unspecified acetabulum');
		insert into product (category_id, name, price, description) values (3, 'Orange - Tangerine', 3917, 'Pathological fracture, unspecified finger(s)');
		insert into product (category_id, name, price, description) values (3, 'Wine - Magnotta - Belpaese', 2858, null);
		insert into product (category_id, name, price, description) values (2, 'Knife Plastic - White', 9444, 'Other fracture of right lower leg');
		insert into product (category_id, name, price, description) values (1, 'Wine - Casablanca Valley', 6713, 'Burn of unspecified degree of left scapular region, sequela');
		insert into product (category_id, name, price, description) values (1, 'Myers Planters Punch', 8376, 'Displaced fracture of olecranon process with intraarticular extension of right ulna, subsequent encounter for open fracture type IIIA, IIIB, or IIIC with malunion');
		insert into product (category_id, name, price, description) values (1, 'Lamb Tenderloin Nz Fr', 2298, 'Corrosion of third degree of multiple left fingers (nail), including thumb, initial encounter');
		insert into product (category_id, name, price, description) values (2, 'Tequila - Sauza Silver', 4475, 'Cauliflower ear, right ear');
		insert into product (category_id, name, price, description) values (3, 'Wine - Cotes Du Rhone', 7184, 'Exanthema subitum [sixth disease] due to human herpesvirus 7');
		insert into product (category_id, name, price, description) values (1, 'Island Oasis - Pina Colada', 4640, null);
		insert into product (category_id, name, price, description) values (2, 'Cheese Cloth', 9571, null);
		insert into product (category_id, name, price, description) values (3, 'Chinese Lemon Pork', 9218, 'Unspecified traumatic spondylolisthesis of fourth cervical vertebra');
		insert into product (category_id, name, price, description) values (3, 'Salt - Kosher', 7390, 'Nondisplaced bicondylar fracture of right tibia, subsequent encounter for open fracture type I or II with delayed healing');
		insert into product (category_id, name, price, description) values (3, 'Cocoa Powder - Natural', 9028, 'Solitary bone cyst, tibia and fibula');
		insert into product (category_id, name, price, description) values (2, 'Beans - Yellow', 7178, 'Corrosion of third degree of right thigh, sequela');
		insert into product (category_id, name, price, description) values (3, 'Longos - Burritos', 6817, 'Gastrostomy hemorrhage');
		insert into product (category_id, name, price, description) values (1, 'Chilli Paste, Hot Sambal Oelek', 6413, 'Displaced fracture of head of right radius');
		insert into product (category_id, name, price, description) values (1, 'Rolled Oats', 2617, 'Fracture of orbital floor, right side, sequela');
		insert into product (category_id, name, price, description) values (2, 'Lettuce - Lolla Rosa', 7357, 'Monoplegia of lower limb affecting right nondominant side');
		insert into product (category_id, name, price, description) values (3, 'Wine - Taylors Reserve', 1204, 'Unspecified open wound of abdominal wall, unspecified quadrant without penetration into peritoneal cavity');
		insert into product (category_id, name, price, description) values (1, 'Tea - Earl Grey', 2725, 'Pressure ulcer of right elbow, unstageable');
		insert into product (category_id, name, price, description) values (1, 'Muffin Hinge - 211n', 5430, 'Other displaced fracture of base of first metacarpal bone, unspecified hand, initial encounter for closed fracture');	
	`

	_, err := DB.Exec(sqlStatement)

	if err != nil {
		panic(err)
	}

}
