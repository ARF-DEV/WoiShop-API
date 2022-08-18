package main

import (
	"azura-lab-intern/study-case-1/controller"
	"azura-lab-intern/study-case-1/repository"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "123456"
	dbname   = "study_case_1"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Database Connected!")

	productRepo := repository.CreateProductRepository(db)
	categoryRepo := repository.CreateCategoryRepository(db)

	mux := controller.NewMux(categoryRepo, productRepo)

	log.Println("Listening in port 8000")
	http.ListenAndServe(":8000", mux)
}
