package main

import (
	"belajar-gin/routers"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "admin123"
	dbname   = "db-book-sql"
)

func main() {
	var PORT = ":8080"

	routers.StartServer().Run(PORT)

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := openDB(psqlInfo)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	fmt.Println("Successfully connected to database")

	// psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// db, err := openDB(psqlInfo)
	// if err != nil {
	// 	panic(err)
	// }

	// defer db.Close()

	// fmt.Println("Successfully connected to database")
	// postBooks(db)
}

// func postBooks(db *sql.DB) {
// 	sqlStatement := `
// 	insert into "books"("title", "author", "description")
// 	values($1, $2, $3)
// 	`

// 	// _, err := db.Exec(sqlStatement, newBook.Title, newBook.Author, newBook.Desc)
// 	_, err := db.Exec(sqlStatement, "One Piece", "Eichiro Oda", "Mugiwara")

// 	if err != nil {
// 		panic(err)
// 	}

// }

func openDB(psqlInfo string) (*sql.DB, error) {
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
