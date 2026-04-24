package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"

	"github.com/Carlosaac23/go-bookstore/internal/service"
	"github.com/Carlosaac23/go-bookstore/internal/store"
	"github.com/Carlosaac23/go-bookstore/internal/transport"
)

func main() {
	// Connect to SQLite
	db, err := sql.Open("sqlite3", "./books.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create table (books) if no exists
	q := `
		CREATE TABLE IF NOT EXISTS books (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			title TEXT NOT NULL,
			author TEXT NOT NULL
		)
	`
	if _, err := db.Exec(q); err != nil {
		log.Fatal(err.Error())
	}

	// Add dependencies
	bookStore := store.New(db)
	bookService := service.New(bookStore)
	bookHandler := transport.New(bookService)

	// Config routes
	http.HandleFunc("/books", bookHandler.HandleBooks)
	http.HandleFunc("/books/", bookHandler.HandleBookByID)

	fmt.Println("Server running at http://localhost:8000")
	fmt.Println("API Endpoints:")
	fmt.Println("  GET     /books       - Get all books")
	fmt.Println("  POST    /books       - Create a new book")
	fmt.Println("  GET     /books/{id}  - Get a specific book")
	fmt.Println("  PUT     /books/{id}  - Update a book")
	fmt.Println("  DELETE  /books/{id}  - Delete a book")

	// Run and listen server
	log.Fatal(http.ListenAndServe(":8000", nil))
}
