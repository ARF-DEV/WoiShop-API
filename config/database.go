package config

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
)

func ConfigDatabase() *sql.DB {

	godotenv.Load()

	// host := os.Getenv("PGHOST")
	// port := os.Getenv("PGPORT")
	// user := os.Getenv("PGUSER")
	// password := os.Getenv("PGPASSWORD")
	// dbname := os.Getenv("PGDATABASE")
	dbURL := os.Getenv("DATABASE_URL")
	// fmt.Println(host, port, user, password, dbname)
	// psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=require",
	// host, port, user, password, dbname)

	db, err := sql.Open("pgx", dbURL)

	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Database Connected!")
	return db
}
