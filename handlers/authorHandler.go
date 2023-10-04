package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jalberto7/xyz-books/db"
	"github.com/jalberto7/xyz-books/models"
)

// Get all authors
func GetAllAuthors(c *gin.Context) {
	var authors []models.Author
	db.DB.Raw("SELECT * FROM author").Scan(&authors)
	c.JSON(200, gin.H{
		"authors": authors,
	})
}

// Create author
func CreateAuthor(c *gin.Context) {
	var author models.Author
	if err := c.ShouldBindJSON(&author); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	db.DB.Raw("INSERT INTO author(first_name, last_name, middle_name) VALUES($1, $2, $3) RETURNING*", author.FirstName, author.LastName, author.MiddleName).Scan(&author)
	c.JSON(200, gin.H{
		"author": author,
	})
}

// Get author by id
func GetAuthorById(c *gin.Context) {
	var author models.Author
	result := db.DB.Raw("SELECT * FROM author WHERE author_id = $1", c.Param("id")).Scan(&author)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Id not exist",
		})
		return
	}

	c.JSON(200, gin.H{
		"author": author,
	})
}

func UpdateAuthor(c *gin.Context) {
	var author models.Author
	if err := c.ShouldBindJSON(&author); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	db.DB.Raw("UPDATE author SET first_name = $1, last_name = $2, middle_name = $3 WHERE author_id = $4", author.FirstName, author.LastName, author.MiddleName, c.Param("id")).Scan(&author)
	c.JSON(200, gin.H{
		"author": author,
	})
}

// Delete author
func DeleteAuthor(c *gin.Context) {
	db.DB.Exec("DELETE FROM author WHERE author_id = $1", c.Param("id"))
	c.JSON(200, gin.H{
		"message": "Author deleted",
	})
}
