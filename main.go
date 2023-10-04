package main

import (
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jalberto7/xyz-books/db"
	"github.com/jalberto7/xyz-books/handlers"
	"github.com/jalberto7/xyz-books/initializers"
)

func init() {
	initializers.LoadEnvironmentVariables()
	db.ConnectDB()
}

func main() {
	fmt.Println("Process Running")

	r := gin.Default()

	// enable cors
	r.Use(cors.Default())
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/authors", handlers.GetAllAuthors)
	r.POST("/authors/create", handlers.CreateAuthor)
	r.GET("/authors/:id", handlers.GetAuthorById)
	r.PUT("/authors/update/:id", handlers.UpdateAuthor)
	r.DELETE("/authors/delete/:id", handlers.DeleteAuthor)

	r.GET("/books", handlers.GetAllBooks)
	r.POST("/books/create", handlers.CreateBook)
	r.GET("/books/:id", handlers.GetBookById)
	r.PUT("/books/update/:id", handlers.UpdateBook)
	r.DELETE("/books/delete/:id", handlers.DeleteBook)

	r.GET("/publishers", handlers.GetAllPublishers)
	r.POST("/publishers/create", handlers.CreatePublisher)
	r.GET("/publishers/:id", handlers.GetPublisherById)
	r.PUT("/publishers/update/:id", handlers.UpdatePublisher)
	r.DELETE("/publishers/delete/:id", handlers.DeletePublisher)

	r.GET("/books/information/:title", handlers.GetBookInformation)
	r.GET("/books/information", handlers.GetAllBooksInformation)
	r.POST("/books/information/create", handlers.CreateAuthorBook)
	r.PUT("/books/information/:author_id/:book_id", handlers.UpdateAuthorBook)
	r.DELETE("/books/information/delete/:author_id/:book_id", handlers.DeleteAuthorBook)

	// ISBN CONVERSION
	r.GET("/convert-isbn", handlers.ConvertISBN)

	r.Run() // listen and serve on 0.0.0.0:8000
}
