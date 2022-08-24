package main

import (
	"azura-lab-intern/study-case-1/controller"
	"azura-lab-intern/study-case-1/helpers"
	"azura-lab-intern/study-case-1/repository"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		panic(err.Error())
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
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

	helpers.InitDB(db)
	productRepo := repository.CreateProductRepository(db)
	categoryRepo := repository.CreateCategoryRepository(db)

	mux := controller.NewMux(categoryRepo, productRepo)

	log.Println("Listening in port 8000")
	http.ListenAndServe(":8000", mux)
}
