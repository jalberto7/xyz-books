package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/jalberto7/xyz-books/db"
	"github.com/jalberto7/xyz-books/models"
)

// Get all publishers
func GetAllPublishers(c *gin.Context) {
	var publishers []models.Publisher
	db.DB.Raw("SELECT * FROM publisher").Scan(&publishers)
	c.JSON(200, gin.H{
		"publishers": publishers,
	})
}

// Get publisher by id
func GetPublisherById(c *gin.Context) {
	var publisher models.Publisher
	db.DB.Raw("SELECT * FROM publisher WHERE publisher_id =?", c.Param("id")).Scan(&publisher)
	c.JSON(200, gin.H{
		"publisher": publisher,
	})
}

// Create a new publisher
func CreatePublisher(c *gin.Context) {
	var publisher models.Publisher
	if err := c.ShouldBindJSON(&publisher); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	db.DB.Raw("INSERT INTO publisher(publisher_name) VALUES($1)", publisher.PublisherName).Scan(&publisher)

	c.JSON(200, gin.H{
		"publisher": publisher,
	})
}

// Update publisher
func UpdatePublisher(c *gin.Context) {
	var publisher models.Publisher
	if err := c.ShouldBindJSON(&publisher); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	db.DB.Raw("UPDATE publisher SET publisher_name = $1 WHERE publisher_id = $2", publisher.PublisherName, c.Param("id")).Scan(&publisher)

	c.JSON(200, gin.H{
		"publisher": publisher,
	})
}

// Delete publisher
func DeletePublisher(c *gin.Context) {
	db.DB.Exec("DELETE FROM publisher WHERE publisher_id = $1", c.Param("id"))
	c.JSON(200, gin.H{
		"message": "Publisher deleted",
	})
}
