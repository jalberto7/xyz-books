package initializers

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

// Environment Variables
func LoadEnvironmentVariables() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
		log.Fatal("Error loading .env file")
	}
}
