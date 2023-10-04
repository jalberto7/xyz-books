package handlers

import (
	"database/sql"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jalberto7/xyz-books/models"
	"github.com/lib/pq"
)

// GetAllBooksInformation
func GetAllBooksInformation(c *gin.Context) {
	// Database connection information
	dbURL := os.Getenv("DB_URL")
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Prepare the SQL statement with JOIN
	query := `
	SELECT b.title, array_agg(concat(a.first_name, ' ', a.last_name)) AS author,
		b.isbn13, b.isbn10, b.publication_year, p.publisher_name,
		b.edition, b.list_price
	FROM book b
	JOIN author_book ab ON b.book_id = ab.book_id
	JOIN author a ON ab.author_id = a.author_id
	JOIN publisher p ON b.publisher_id = p.publisher_id
	GROUP BY b.title, b.isbn13, b.isbn10, b.publication_year, p.publisher_name, b.edition, b.list_price`

	// Query the database
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var books []models.BookInformation

	// Fetch the book information from the query results
	for rows.Next() {
		var book models.BookInformation
		if err := rows.Scan(&book.Title, pq.Array(&book.Author), &book.ISBN13, &book.ISBN10, &book.PublicationYear, &book.Publisher, &book.Edition, &book.ListPrice); err != nil {
			log.Fatal(err)
		}
		books = append(books, book)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	// Return the book information as JSON response
	c.JSON(http.StatusOK, gin.H{
		"books": books,
	})
}

// GetBookInformation by title
func GetBookInformation(c *gin.Context) {
	// Retrieve the book title from the URL parameter
	titleParam := c.Param("title")

	// URL decode the title
	title, err := url.QueryUnescape(titleParam)
	if err != nil {
		log.Fatal(err)
	}

	// Database connection information
	dbURL := os.Getenv("DB_URL")
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Prepare the SQL statement with JOIN
	query := `
	SELECT b.title, array_agg(concat(a.first_name, ' ', a.last_name)) AS author,
		b.isbn13, b.isbn10, b.publication_year, p.publisher_name,
		b.edition, b.list_price
	FROM book b
	JOIN author_book ab ON b.book_id = ab.book_id
	JOIN author a ON ab.author_id = a.author_id
	JOIN publisher p ON b.publisher_id = p.publisher_id
	WHERE b.title = $1
	GROUP BY b.title, b.isbn13, b.isbn10, b.publication_year, p.publisher_name, b.edition, b.list_price`

	// Query the database
	rows, err := db.Query(query, title)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var books []models.BookInformation

	// Fetch the book information from the query results
	for rows.Next() {
		var book models.BookInformation
		if err := rows.Scan(&book.Title, pq.Array(&book.Author), &book.ISBN13, &book.ISBN10, &book.PublicationYear, &book.Publisher, &book.Edition, &book.ListPrice); err != nil {
			log.Fatal(err)
		}
		books = append(books, book)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	// Return the book information as JSON response
	c.JSON(http.StatusOK, books)
}

func CreateAuthorBook(c *gin.Context) {
	// Retrieve the author ID and book title from the request body
	var req struct {
		AuthorID  int    `json:"author_id"`
		BookTitle string `json:"book_title"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Database connection information
	dbURL := os.Getenv("DB_URL")
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Get the book ID based on the provided book title
	var bookID int
	err = db.QueryRow("SELECT book_id FROM book WHERE title = $1", req.BookTitle).Scan(&bookID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get book ID"})
		return
	}

	// Insert the new entry in the author_book table
	_, err = db.Exec("INSERT INTO author_book (author_id, book_id) VALUES ($1, $2)", req.AuthorID, bookID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return a success response
	c.JSON(http.StatusOK, gin.H{"message": "New author-book entry created"})
}

// UpdateAuthorBookEntry updates
/*\
For updating an author-book record, select the HTTP method as `PUT`.
In the "Body" section of the request, choose the "raw" option and set the data format to JSON.
Provide the updated author-book details in the request body. For example:
json
{
	"author_id": 3,
	"book_id": 4
}
*/
func UpdateAuthorBook(c *gin.Context) {
	// Retrieve the author ID and book ID from the path parameters
	authorID := c.Param("author_id")
	bookID := c.Param("book_id")

	// Retrieve the updated author-book details from the request body
	var req struct {
		AuthorID int `json:"author_id"`
		BookID   int `json:"book_id"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Database connection information
	dbURL := os.Getenv("DB_URL")
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Update the author-book details in the database
	_, err = db.Exec("UPDATE author_book SET author_id = $1, book_id = $2 WHERE author_id = $3 AND book_id = $4", req.AuthorID, req.BookID, authorID, bookID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return a success response
	c.JSON(http.StatusOK, gin.H{"message": "Author-book record updated successfully"})
}

// delete a record
/*
	Launch Postman and make sure you have the correct endpoint URLs configured.
	In both cases, the endpoints should be similar to `http://localhost:8000/author-book/1/2`
	where `1` is the author ID and `2` is the book ID to update or delete.
*/
func DeleteAuthorBook(c *gin.Context) {
	// Retrieve the author ID and book ID from the path parameters
	authorID := c.Param("author_id")
	bookID := c.Param("book_id")

	// Database connection information
	dbURL := os.Getenv("DB_URL")
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Delete the author-book record from the database
	_, err = db.Exec("DELETE FROM author_book WHERE author_id = $1 AND book_id = $2", authorID, bookID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return a success response
	c.JSON(http.StatusOK, gin.H{"message": "Author-book record deleted successfully"})
}
