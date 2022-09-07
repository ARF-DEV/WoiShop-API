package repository

import (
	"azura-lab-intern/study-case-1/models"
	"database/sql"
)

type UserRepository struct {
	db *sql.DB
}

func CreateUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (u *UserRepository) InsertUser(user models.User) (*models.User, error) {
	sqlStatement := `
	INSERT INTO user_data (full_name, email, phone_num,  referal_code) 
	VALUES ($1, $2, $3, $4) 
	RETURNING id, full_name, email, phone_num, referal_code;`

	var newUser models.User
	err := u.db.QueryRow(sqlStatement, user.Name, user.Email, user.NoTelp, user.ReferalCode).Scan(&newUser.ID, &newUser.Name, &newUser.Email, &newUser.NoTelp, &newUser.ReferalCode)

	if err != nil {
		return nil, err
	}

	return &newUser, nil
}

func (u *UserRepository) GetUserByNoTelp(noTelp string) (*models.User, error) {
	sqlStatement := `
	SELECT id, full_name, phone_num, email, referal_code FROM user_data WHERE phone_num = $1
	`

	row := u.db.QueryRow(sqlStatement, noTelp)

	var user models.User
	err := row.Scan(&user.ID, &user.Name, &user.NoTelp, &user.Email, &user.ReferalCode)

	if err != nil {
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
		return nil, err
	}

	return &user, nil
}
