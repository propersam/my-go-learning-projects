package main

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	uuid "github.com/satori/go.uuid"
)

func main() {
	// Create New Server Instance
	e := echo.New()

	// Declaring Available routes
	e.GET("/", welcomeRoute)
	e.GET("/books/:id", getBook)
	e.POST("/books", addBook)
	// e.PUT("/books/:id", updateBook)
	// e.DELETE("/books/:id", removeBook)

	// Start Server
	e.Logger.Fatal(e.Start(":1323"))
}

// Book object
type Book struct {
	ID     uuid.UUID `json:"id" form:"id" query:"id"`
	Title  string    `json:"title" form:"title" query:"title"`
	Author string    `json:"author" form:"author" query:"author"`
}

// Welcome home page
func welcomeRoute(c echo.Context) (err error) {
	return c.String(http.StatusOK, "Hello World!\n")
}

func addBook(c echo.Context) (err error) {

	// Generate UUID version 4 or Panic on error
	id := uuid.Must(uuid.NewV4())

	//Create book object
	book := new(Book)
	if err := c.Bind(book); err != nil {
		output := map[string]string{
			"Error": "An Invalid Parameter field was passed",
		}
		return c.JSONPretty(http.StatusBadRequest, output, "	")
	}
	book.ID = id

	fmt.Println(*book)
	// connecting to database
	db, err := sql.Open("mysql", "root:db-key@/go_learning")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//prepare statement for inserting Data
	stmtIns, err := db.Prepare("INSERT INTO books(ID, Title, Author) VALUES( ?, ?, ?);")
	if err != nil {
		panic(err.Error()) // proper error handling
	}
	defer stmtIns.Close() // Close the statement when we exit this function

	// execute insert query
	_, err = stmtIns.Exec(book.ID, book.Title, book.Author)
	if err != nil {
		panic(err.Error())
	}

	// Return book uuid as a sign of success
	output := map[string]interface{}{
		"book": book,
	}

	return c.JSONPretty(http.StatusOK, output, "	")
}

func getBook(c echo.Context) (err error) {

	//fetch ID
	// convert string to uuid
	id := c.Param("id")

	uuid, err := uuid.FromString(id)
	if err != nil {
		output := map[string]string{
			"Error": "Invalid uuid Sent!!",
		}
		return c.JSONPretty(http.StatusBadRequest, output, "	")
	}

	// Open Connection Handler to database
	db, err := sql.Open("mysql", "root:db-key@/go_learning")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Query Database for book with "id"
	fetchQuery, err := db.Prepare("SELECT title, author FROM books WHERE id = ?")
	if err != nil {
		// If error is caused by no data found
		if err == sql.ErrNoRows {
			// return nil to json output
			output := map[string]interface{}{
				"books": nil,
			}
			return c.JSONPretty(http.StatusOK, output, "	")
		}
		// else if error is not as expected above
		panic(err.Error()) // proper error handling
	}
	defer fetchQuery.Close()

	book := Book{
		ID: uuid,
	}

	fetchQuery.QueryRowContext(context.Background(), id).Scan(&book.Title, &book.Author)

	output := map[string]interface{}{
		"books": book,
	}
	return c.JSONPretty(http.StatusOK, output, "	")
}
