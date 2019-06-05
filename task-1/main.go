package main

import (
	"context"
	"database/sql"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
)

// Declaring global varibles
var db *sql.DB
var ctx context.Context

func init() {
	db = createDbPool()
	ctx = context.Background()
}

func main() {
	// Create New Server Instance
	e := echo.New()

	// Declaring Available routes
	e.GET("/", welcomeRouteHandler)
	// e.GET("/books/:id", showBookHandler)
	e.POST("/books", addBookHandler)
	// e.PUT("/books/:id", updateBook)
	// e.DELETE("/books/:id", removeBook)
	// Start Server
	e.Logger.Fatal(e.Start(":1323"))
}

// // Book object
// type Book struct {
// 	ID     uuid.UUID `json:"id" form:"id" query:"id"`
// 	Title  string    `json:"title" form:"title" query:"title"`
// 	Author string    `json:"author" form:"author" query:"author"`
// }

// Welcome home page
func welcomeRouteHandler(c echo.Context) (err error) {
	return c.String(http.StatusOK, "Hello World!\n")
}

// func getBook(c echo.Context) (err error) {

// 	//fetch ID
// 	// convert string to uuid
// 	id := c.Param("id")

// 	uuid, err := uuid.FromString(id)
// 	if err != nil {
// 		output := map[string]string{
// 			"Error": "Invalid uuid Sent!!",
// 		}
// 		return c.JSONPretty(http.StatusBadRequest, output, "	")
// 	}

// 	// // Open Connection Handler to database
// 	// db, err := sql.Open("mysql", "root:db-key@/go_learning")
// 	// if err != nil {
// 	// 	panic(err)
// 	// }
// 	// defer db.Close()

// 	// Query Database for book with "id"
// 	fetchQuery, err := db.Prepare("SELECT title, author FROM books WHERE id = ?")
// 	if err != nil {
// 		// If error is caused by no data found
// 		if err == sql.ErrNoRows {
// 			// return nil to json output
// 			output := map[string]interface{}{
// 				"books": nil,
// 			}
// 			return c.JSONPretty(http.StatusOK, output, "	")
// 		}
// 		// else if error is not as expected above
// 		panic(err.Error()) // proper error handling
// 	}
// 	defer fetchQuery.Close()

// 	book := Book{
// 		ID: uuid,
// 	}

// 	fetchQuery.QueryRowContext(ctx, id).Scan(&book.Title, &book.Author)

// 	output := map[string]interface{}{
// 		"books": book,
// 	}
// 	return c.JSONPretty(http.StatusOK, output, "	")
// }
