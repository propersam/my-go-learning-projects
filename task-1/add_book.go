package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/labstack/echo"
	"github.com/myesui/uuid"
)

// Book object
type Book struct {
	ID     string `json:"id" form:"id" query:"id"`
	Title  string `json:"title" form:"title" query:"title"`
	Author string `json:"author" form:"author" query:"author"`
	ISBN   string `json:"isbn" form:"isbn" query:"isbn"`
}

// addBookHandler is the handler used to store book details to DB
func addBookHandler(c echo.Context) error {
	// // Get context headers
	// contxt := c.Request().Context()

	//Create book object
	book := new(Book)
	if err := c.Bind(book); err != nil {
		return c.JSON(http.StatusBadRequest, ErrorFmt(err))
	}
	// ID:
	// Generate UUID version 4 in String format
	id := uuid.NewV4().String()
	book.ID = id

	// Input Validation
	book.Author = strings.TrimSpace(book.Author)
	book.Title = strings.TrimSpace(book.Title)
	book.ISBN = strings.TrimSpace(book.ISBN)

	// Author:
	if book.Author == "" {
		return c.JSON(http.StatusBadRequest, ErrorFmt("book author field must not be empty"))
	}

	// Title:
	if book.Title == "" {
		return c.JSON(http.StatusBadRequest, ErrorFmt("book title field must not be empty"))
	}

	// ISBN:
	if book.ISBN == "" {
		return c.JSON(http.StatusBadRequest, ErrorFmt("book ISBN field must not be empty"))
	}
	// Attempt to save book in database

	// obtain exclusive connection for this write action
	conn, err := db.Conn(ctx)
	defer conn.Close() //return connection back to pool

	// Perform write operation transaction
	txn, err := conn.BeginTx(ctx, nil)
	stmt := "INSERT INTO books(ID, Title, Author, isbn) VALUES( ?, ?, ?, ?);"
	_, err = txn.Exec(stmt, book.ID, book.Title, book.Author, book.ISBN)
	if err != nil {
		log.Println(err)
		txn.Rollback()
		return c.JSON(http.StatusInternalServerError, ErrorFmt("something went wrong. Try again"))
	}
	txn.Commit()

	return c.NoContent(http.StatusOK)
}
