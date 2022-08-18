package repository

import (
	"azura-lab-intern/study-case-1/models"
	"database/sql"
	"log"

	"github.com/lib/pq"
)

type ProductRepository struct {
	db *sql.DB
}

func CreateProductRepostitory(db *sql.DB) *ProductRepository {
	return &ProductRepository{
		db: db,
	}
}

func (pr *ProductRepository) GetAllProduct() ([]models.Product, *pq.Error) {
	sqlStatement := `SELECT id, name, price FROM product`

	rows, err := pr.db.Query(sqlStatement)

	if err != nil {
		err := err.(*pq.Error)
		log.Println("error on get all product : ", err.Message)
		return nil, err
	}

	var products []models.Product
	for rows.Next() {
		var p models.Product

		err := rows.Scan(&p.ID, &p.Name, &p.Price)

		if err != nil {
			err := err.(*pq.Error)
			log.Println("error on get all product : ", err.Message)
			return nil, err
		}

	}

	return products, nil
}
