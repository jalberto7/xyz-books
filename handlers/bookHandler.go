package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/jalberto7/xyz-books/db"
	"github.com/jalberto7/xyz-books/models"
)

// Get all Book
func GetAllBooks(c *gin.Context) {
	var books []models.Book
	db.DB.Raw("SELECT * FROM book").Scan(&books)
	c.JSON(200, gin.H{
		"books": books,
	})
}

// Get book by id
func GetBookById(c *gin.Context) {
	var book models.Book
	db.DB.Raw("SELECT * FROM book WHERE book_id =?", c.Param("id")).Scan(&book)
	c.JSON(200, gin.H{
		"book": book,
	})
}

// Create a new Book
func CreateBook(c *gin.Context) {
	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	// GRANT ALL PRIVILEGES ON SEQUENCE book_book_id_seq to joemar;
	db.DB.Raw("INSERT INTO book(title, isbn13, isbn10, publication_year, publisher_id, edition, list_price) VALUES($1, $2, $3, $4, $5, $6, $7)",
		book.Title, book.ISBN13, book.ISBN10, book.PublicationYear, book.PublisherID, book.Edition, book.ListPrice).Scan(&book)

	c.JSON(200, gin.H{
		"book": book,
	})
}

// Update book
func UpdateBook(c *gin.Context) {
	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	db.DB.Raw("UPDATE book SET title = $1, isbn13 = $2, isbn10 = $3, publication_year = $4, publisher_id = $5, edition = $6, list_price = $7 WHERE book_id = $8",
		book.Title, book.ISBN13, book.ISBN10, book.PublicationYear, book.PublisherID, book.Edition, book.ListPrice, c.Param("id")).Scan(&book)

	c.JSON(200, gin.H{
		"book": book,
	})
}

// Delete book
func DeleteBook(c *gin.Context) {
	db.DB.Exec("DELETE FROM book WHERE book_id = $1", c.Param("id"))
	c.JSON(200, gin.H{
		"message": "Book deleted",
	})
}
