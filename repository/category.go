package repository

import (
	"azura-lab-intern/study-case-1/models"
	"database/sql"
)

type CategoryRepository struct {
	db *sql.DB
}

func CreateCategoryRepository(db *sql.DB) *CategoryRepository {
	return &CategoryRepository{
		db: db,
	}
}

func (cr *CategoryRepository) GetAllCategory() ([]models.Category, error) {

	sqlStatement := `SELECT id, name FROM category`

	rows, err := cr.db.Query(sqlStatement)

	if err != nil {
		return nil, err
	}

	var results []models.Category
	for rows.Next() {
		var c models.Category

		err := rows.Scan(&c.ID, &c.Name)

		if err != nil {
			return nil, err
		}

		results = append(results, c)
	}

	return results, nil
}
