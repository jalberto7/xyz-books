package main

import (
	"github.com/jalberto7/xyz-books/db"
	"github.com/jalberto7/xyz-books/initializers"
	"github.com/jalberto7/xyz-books/models"
)

func init() {
	db.ConnectDB()
	initializers.LoadEnvironmentVariables()

}

func main() {
	db.DB.AutoMigrate(&models.Author{})
}
