package repository

import (
	"azura-lab-intern/study-case-1/models"
	"database/sql"
	"log"
)

type UserRepository struct {
	db *sql.DB
}

func CreateUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (u *UserRepository) InsertUser(user models.User) error {
	sqlStatement := `INSERT INTO user_data (full_name, email, phone_num,  referal_code) VALUES ($1, $2, $3, $4);`

	_, err := u.db.Exec(sqlStatement, user.Name, user.Email, user.NoTelp, user.ReferalCode)

	if err != nil {
		log.Println("Error on running sqlStatement: ", err.Error())
		return err
	}

	return nil
}

func (u *UserRepository) GetUserByNoTelp(noTelp string) (*models.User, error) {
	sqlStatement := `
	SELECT id, full_name, phone_num, email, referal_code FROM user_data WHERE phone_num = $1
	`

	row := u.db.QueryRow(sqlStatement, noTelp)

	var user models.User
	err := row.Scan(&user.ID, &user.Name, &user.NoTelp, &user.Email, &user.ReferalCode)

	if err != nil {
		log.Println("Error on Scan User By Telp: ", err.Error())
		return nil, err
	}

	return &user, nil
}

func (u *UserRepository) GetUserByEmail(email string) (*models.User, error) {

	sqlStatement := `
	SELECT id, full_name, phone_num, email, referal_code FROM user_data WHERE email = $1
	`

	row := u.db.QueryRow(sqlStatement, email)

	var user models.User
	err := row.Scan(&user.ID, &user.Name, &user.NoTelp, &user.Email, &user.ReferalCode)

	if err != nil {
		log.Println("Error on Scan User By Telp: ", err.Error())
		return nil, err
	}

	return &user, nil
}
