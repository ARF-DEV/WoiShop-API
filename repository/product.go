package repository

import (
	"azura-lab-intern/study-case-1/helpers"
	"azura-lab-intern/study-case-1/models"
	"database/sql"
	"log"

	"github.com/lib/pq"
)

type ProductRepository struct {
	db *sql.DB
}

func CreateProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{
		db: db,
	}
}

func (pr *ProductRepository) GetProductByCategory(categoryName string) ([]helpers.ProductCategory, *pq.Error) {
	sqlStatement := `
	SELECT 
		p.id, p.name, p.price, c.id, c.name
		FROM 
			product as p
		INNER JOIN category as c
			ON category.id = product.category_id 
		WHERE category.name = ?;`

	rows, err := pr.db.Query(sqlStatement, categoryName)

	if err != nil {
		err := err.(*pq.Error)
		log.Println("error on get all product : ", err.Message)
		return nil, err
	}

	var products []helpers.ProductCategory
	for rows.Next() {
		var p helpers.ProductCategory

		err := rows.Scan(&p.ProductID, &p.ProductName, &p.ProductPrice, &p.CategoryID, &p.ProductName)

		if err != nil {
			err := err.(*pq.Error)
			log.Println("error on get all product : ", err.Message)
			return nil, err
		}

		products = append(products, p)
	}

	return products, nil

}

func (pr *ProductRepository) GetProductByID(id int) (*models.Product, *pq.Error) {

	sqlStatement := `SELECT id, name, price FROM product WHERE id = ?`

	row := pr.db.QueryRow(sqlStatement)

	var product models.Product
	err := row.Scan(&product.ID, &product.Name, &product.Price)
	if err != nil {
		err := err.(*pq.Error)
		log.Println("error on get all product : ", err.Message)
		return nil, err
	}
	return &product, nil
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

		products = append(products, p)
	}

	return products, nil
}
