package repository

import (
	"azura-lab-intern/study-case-1/helpers"
	"azura-lab-intern/study-case-1/models"
	"database/sql"
	"log"
)

type ProductRepository struct {
	db *sql.DB
}

func CreateProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{
		db: db,
	}
}

func (pr *ProductRepository) GetProductByCategory(categoryName string) ([]helpers.ProductCategory, error) {
	sqlStatement := `
	SELECT 
		p.id, p.name, p.price, c.id, c.name
		FROM 
			product as p
		INNER JOIN category as c
			ON c.id = p.category_id 
		WHERE c.name = $1;`

	rows, err := pr.db.Query(sqlStatement, categoryName)

	if err != nil {

		log.Println("error on get all product : ", err.Error())
		return nil, err
	}

	var products []helpers.ProductCategory
	for rows.Next() {
		var p helpers.ProductCategory

		err := rows.Scan(&p.ProductID, &p.ProductName, &p.ProductPrice, &p.CategoryID, &p.CategoryName)

		if err != nil {

			log.Println("error on get all product : ", err.Error())
			return nil, err
		}

		products = append(products, p)
	}

	return products, nil

}

func (pr *ProductRepository) GetProductByID(id int) (*models.Product, error) {

	sqlStatement := `SELECT id, name, price, description FROM product WHERE id = $1`

	row := pr.db.QueryRow(sqlStatement, id)

	var product models.Product
	err := row.Scan(&product.ID, &product.Name, &product.Price, &product.Description)
	if err != nil {

		log.Println("error on get product : ", err.Error())
		return nil, err
	}
	return &product, nil
}

func (pr *ProductRepository) GetAllProduct() ([]models.Product, error) {
	sqlStatement := `SELECT id, name, price FROM product`

	rows, err := pr.db.Query(sqlStatement)

	if err != nil {

		log.Println("error on get product : ", err.Error())
		return nil, err
	}

	var products []models.Product
	for rows.Next() {
		var p models.Product

		err := rows.Scan(&p.ID, &p.Name, &p.Price)

		if err != nil {

			log.Println("error on get product : ", err.Error())
			return nil, err
		}

		products = append(products, p)
	}

	return products, nil
}
