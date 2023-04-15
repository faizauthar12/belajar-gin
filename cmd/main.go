package main

import (
	"belajar-gin/models"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "admin123"
	dbname   = "db-book-sql"
)

const PORT = ":8080"

type application struct {
	books *models.BooksModel
}

func main() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Jakarta", host, user, password, dbname, port)
	db, err := openDB(dsn)

	if err != nil {
		panic(err.Error())
	}

	app := &application{
		books: &models.BooksModel{DB: db},
	}

	route := app.routes()
	route.Run(PORT)
}

func openDB(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}
	return db, nil
}
