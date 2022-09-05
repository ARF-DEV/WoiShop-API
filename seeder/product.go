package seeder

import (
	"database/sql"
	"log"
)

func MigrateProduct(db *sql.DB) {
	sqlStatement := `
	
		drop table IF EXISTS product cascade;
		CREATE TABLE IF NOT EXISTS product (
			id serial PRIMARY KEY,
			category_id INT NOT NULL,
			store_id INT NOT NULL,
			name VARCHAR(150) NOT NULL,
			price INT NOT NULL,
			stock INT NOT NULL,
			description VARCHAR(500),
			image_link VARCHAR(200),
			FOREIGN KEY (category_id) REFERENCES category (id) on delete set null,
			FOREIGN KEY (store_id) REFERENCES store (id) on delete set null
		);
		insert into product (store_id, category_id, name, price, description, stock) values (5, 3, 'Apple - Granny Smith', 64002, 'Bypass Left Fallopian Tube to Uterus with Synthetic Substitute, Percutaneous Endoscopic Approach', 98);
		insert into product (store_id, category_id, name, price, description, stock) values (9, 3, 'Beef Dry Aged Tenderloin Aaa', 49768, 'Supplement Right Radius with Synthetic Substitute, Percutaneous Approach', 1);
		insert into product (store_id, category_id, name, price, description, stock) values (3, 1, 'Frangelico', 10038, 'Supplement Right Basilic Vein with Nonautologous Tissue Substitute, Open Approach', 19);
		insert into product (store_id, category_id, name, price, description, stock) values (2, 2, 'Mix - Cappucino Cocktail', 29615, 'Bypass Left Renal Vein to Lower Vein with Autologous Tissue Substitute, Percutaneous Endoscopic Approach', 34);
		insert into product (store_id, category_id, name, price, description, stock) values (4, 1, 'Gloves - Goldtouch Disposable', 58869, 'Supplement Right Shoulder Muscle with Autologous Tissue Substitute, Percutaneous Endoscopic Approach', 57);
		insert into product (store_id, category_id, name, price, description, stock) values (10, 3, 'Lid Coffee Cup 8oz Blk', 25486, 'Repair Right Foot, Open Approach', 97);
		insert into product (store_id, category_id, name, price, description, stock) values (10, 2, 'Towel Multifold', 30759, 'Drainage of Right Middle Ear, Open Approach', 7);
		insert into product (store_id, category_id, name, price, description, stock) values (5, 3, 'Cookie Dough - Double', 28511, 'Repair Coccyx, External Approach', 28);
		insert into product (store_id, category_id, name, price, description, stock) values (9, 1, 'Cheese - Pied De Vents', 24540, 'Supplement Left Abdomen Tendon with Synthetic Substitute, Percutaneous Endoscopic Approach', 38);
		insert into product (store_id, category_id, name, price, description, stock) values ( 7, 3, 'Beef - Tenderloin - Aa', 65601, 'Insertion of Other Device into Upper Jaw, Percutaneous Endoscopic Approach', 5);
		insert into product (store_id, category_id, name, price, description, stock) values ( 10, 2, 'Trueblue - Blueberry 12x473ml', 30438, 'Excision of Right Pelvic Bone, Open Approach, Diagnostic', 22);
		insert into product (store_id, category_id, name, price, description, stock) values ( 1, 2, 'Ocean Spray - Ruby Red', 49984, 'Dilation of Left Axillary Artery, Bifurcation, with Three Drug-eluting Intraluminal Devices, Percutaneous Approach', 10);
		insert into product (store_id, category_id, name, price, description, stock) values ( 9, 1, 'Beans - Navy, Dry', 46017, 'Excision of Right External Iliac Vein, Percutaneous Endoscopic Approach, Diagnostic', 51);
		insert into product (store_id, category_id, name, price, description, stock) values ( 5, 2, 'Wine - Pinot Grigio Collavini', 47442, 'Extirpation of Matter from Cervical Vertebral Joint, Percutaneous Endoscopic Approach', 55);
		insert into product (store_id, category_id, name, price, description, stock) values ( 3, 1, 'Brandy - Bar', 66445, 'Beam Radiation of Brain Stem using Neutron Capture', 37);
		insert into product (store_id, category_id, name, price, description, stock) values ( 7, 3, 'Wine - Chateau Timberlay', 18691, 'Tinnitus Masker Assessment using Hearing Aid Selection / Fitting / Test Equipment', 51);
		insert into product (store_id, category_id, name, price, description, stock) values ( 3, 1, 'Wine - Red, Gallo, Merlot', 57573, 'Occlusion of Ileocecal Valve with Extraluminal Device, Percutaneous Approach', 33);
		insert into product (store_id, category_id, name, price, description, stock) values ( 2, 2, 'Raspberries - Fresh', 42213, 'Extirpation of Matter from Penis, Open Approach', 61);
		insert into product (store_id, category_id, name, price, description, stock) values ( 10, 1, 'Sea Bass - Whole', 51994, 'Replacement of Right Brachial Vein with Autologous Tissue Substitute, Open Approach', 34);
		insert into product (store_id, category_id, name, price, description, stock) values ( 6, 3, 'Chicken - White Meat With Tender', 8490, 'Replacement of Left Nipple with Autologous Tissue Substitute, Percutaneous Approach', 84);
		insert into product (store_id, category_id, name, price, description, stock) values ( 6, 1, 'Wine - Chardonnay Mondavi', 66988, 'Revision of Nonautologous Tissue Substitute in Lymphatic, Open Approach', 2);
		insert into product (store_id, category_id, name, price, description, stock) values ( 2, 2, 'Wakami Seaweed', 41420, 'Repair Left Knee Region, External Approach', 47);
		insert into product (store_id, category_id, name, price, description, stock) values ( 5, 1, 'Baking Powder', 17543, 'Removal of Nonautologous Tissue Substitute from Right Lower Extremity, Percutaneous Endoscopic Approach', 26);
		insert into product (store_id, category_id, name, price, description, stock) values ( 3, 1, 'Crush - Orange, 355ml', 16458, 'Removal of Traction Apparatus on Left Toe', 82);
		insert into product (store_id, category_id, name, price, description, stock) values ( 9, 2, 'Bread - 10 Grain Parisian', 31798, 'Resection of Right Maxillary Sinus, Open Approach', 87);
		insert into product (store_id, category_id, name, price, description, stock) values ( 1, 1, 'Cream - 35%', 51617, 'Extirpation of Matter from Right Large Intestine, Percutaneous Endoscopic Approach', 49);
		insert into product (store_id, category_id, name, price, description, stock) values ( 2, 1, 'Coffee - Espresso', 41108, 'Computerized Tomography (CT Scan) of Right Tibia/Fibula using Low Osmolar Contrast', 59);
		insert into product (store_id, category_id, name, price, description, stock) values ( 7, 3, 'Sardines', 18095, 'Supplement Right Abdomen Bursa and Ligament with Nonautologous Tissue Substitute, Percutaneous Endoscopic Approach', 67);
		insert into product (store_id, category_id, name, price, description, stock) values ( 8, 2, 'Truffle Cups Green', 9825, 'Bypass Right Common Iliac Artery to Abdominal Aorta with Autologous Arterial Tissue, Percutaneous Endoscopic Approach', 60);
		insert into product (store_id, category_id, name, price, description, stock) values ( 6, 2, 'Swiss Chard - Red', 57859, 'Removal of Diaphragmatic Pacemaker Lead from Diaphragm, Percutaneous Endoscopic Approach', 10);
		insert into product (store_id, category_id, name, price, description, stock) values ( 2, 3, 'Flower - Leather Leaf Fern', 50238, 'Supplement Left Index Finger with Synthetic Substitute, Open Approach', 21);
		insert into product (store_id, category_id, name, price, description, stock) values ( 1, 3, 'Hog / Sausage Casing - Pork', 23997, 'Compression of Right Inguinal Region using Intermittent Pressure Device', 45);
		insert into product (store_id, category_id, name, price, description, stock) values ( 1, 3, 'Alize Gold Passion', 9235, 'Fusion of Right Shoulder Joint, Percutaneous Endoscopic Approach', 48);
		insert into product (store_id, category_id, name, price, description, stock) values ( 3, 3, 'Coffee - Frthy Coffee Crisp', 57321, 'Removal of Infusion Device from Cervicothoracic Vertebral Disc, External Approach', 49);
		insert into product (store_id, category_id, name, price, description, stock) values ( 4, 1, 'Pasta - Fusili Tri - Coloured', 41173, 'Release Upper Esophagus, Percutaneous Endoscopic Approach', 54);
		insert into product (store_id, category_id, name, price, description, stock) values ( 2, 2, 'Bar Special K', 18318, 'Destruction of Soft Palate, Open Approach', 34);
		insert into product (store_id, category_id, name, price, description, stock) values ( 2, 3, 'Soup - Verve - Chipotle Chicken', 64156, 'Dilation of Right Cephalic Vein, Open Approach', 41);
		insert into product (store_id, category_id, name, price, description, stock) values ( 9, 1, 'Jolt Cola - Electric Blue', 46365, 'Supplement Esophagus with Nonautologous Tissue Substitute, Via Natural or Artificial Opening', 73);
		insert into product (store_id, category_id, name, price, description, stock) values ( 9, 3, 'Lemonade - Kiwi, 591 Ml', 50503, 'Occlusion of Left Subclavian Vein with Extraluminal Device, Percutaneous Approach', 63);
		insert into product (store_id, category_id, name, price, description, stock) values ( 8, 2, 'Water - Evian 355 Ml', 32349, 'Excision of Right External Iliac Artery, Open Approach', 97);
		insert into product (store_id, category_id, name, price, description, stock) values ( 3, 1, 'Soho Lychee Liqueur', 28029, 'Removal of Synthetic Substitute from Left Fibula, Open Approach', 65);
		insert into product (store_id, category_id, name, price, description, stock) values ( 2, 1, 'Pasta - Fett Alfredo, Single Serve', 20587, 'Insertion of Limb Lengthening External Fixation Device into Left Upper Femur, Open Approach', 6);
		insert into product (store_id, category_id, name, price, description, stock) values ( 4, 1, 'Potatoes - Fingerling 4 Oz', 47198, 'Traction of Right Hand using Traction Apparatus', 58);
		insert into product (store_id, category_id, name, price, description, stock) values ( 9, 1, 'Bar Mix - Pina Colada, 355 Ml', 4598, 'Destruction of Right Lower Arm Subcutaneous Tissue and Fascia, Open Approach', 95);
		insert into product (store_id, category_id, name, price, description, stock) values ( 8, 2, 'Nut - Hazelnut, Whole', 64204, 'Dilation of Right External Carotid Artery with Drug-eluting Intraluminal Device, Percutaneous Approach', 5);
		insert into product (store_id, category_id, name, price, description, stock) values ( 5, 2, 'Cardamon Seed / Pod', 68229, 'Excision of Cerebral Hemisphere, Percutaneous Endoscopic Approach, Diagnostic', 40);
		insert into product (store_id, category_id, name, price, description, stock) values ( 6, 2, 'Soup - Knorr, Classic Can. Chili', 47307, 'Transfer Left Foot Tendon, Percutaneous Endoscopic Approach', 26);
		insert into product (store_id, category_id, name, price, description, stock) values ( 9, 2, 'Veal - Chops, Split, Frenched', 41912, 'Reposition Left Maxilla, External Approach', 52);
		insert into product (store_id, category_id, name, price, description, stock) values ( 6, 3, 'Cheese - Swiss Sliced', 2878, 'Dilation of Left Peroneal Artery, Percutaneous Endoscopic Approach', 40);
		insert into product (store_id, category_id, name, price, description, stock) values ( 10, 1, 'Bread - Corn Muffaletta', 52517, 'Supplement Left Foot Tendon with Synthetic Substitute, Percutaneous Endoscopic Approach', 52);
	`
	_, err := db.Exec(sqlStatement)

	if err != nil {
		log.Println("Error while Migrating Product")
		panic(err)
	}
}
