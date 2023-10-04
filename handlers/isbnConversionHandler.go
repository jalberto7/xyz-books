package handlers

import (
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

/*
To test the APIs using curl or Postman:

1. Start the server by running `go run main.go` in the terminal.
2. Use curl or Postman to make a GET request to `http://localhost:8000/convert-isbn?isbn=<ISBN>`.
   - Replace `<ISBN>` with the ISBN you want to convert.
     For example, to convert an ISBN-10 to ISBN-13, use:

     curl -X GET "http://localhost:8000/convert-isbn?isbn=0321982381"

     To convert an ISBN-13 to ISBN-10, use:

     curl -X GET "http://localhost:8000/convert-isbn?isbn=9780321982384"

3. The response will contain the converted ISBN.
*/

func ConvertISBN(c *gin.Context) {
	isbn := c.Query("isbn")

	if len(isbn) == 10 {
		isbn13, err := ConvertISBN10ToISBN13(isbn)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusOK, gin.H{"isbn13": isbn13})
		}
	} else if len(isbn) == 13 {
		isbn10, err := ConvertISBN13ToISBN10(isbn)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusOK, gin.H{"isbn10": isbn10})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ISBN"})
	}
}

func ConvertISBN10ToISBN13(isbn10 string) (string, error) {
	isbn10 = strings.ReplaceAll(isbn10, "-", "")
	isbn10 = strings.ReplaceAll(isbn10, " ", "")

	if len(isbn10) != 10 {
		return "", errors.New("Invalid ISBN-10")
	}

	// Calculate the check digit for ISBN-13
	isbn12 := "978" + isbn10[:9]
	checkDigit := 0
	for i, c := range isbn12 {
		digit, err := strconv.Atoi(string(c))
		if err != nil {
			return "", err
		}
		if i%2 == 0 {
			checkDigit += digit
		} else {
			checkDigit += digit * 3
		}
	}
	checkDigit = 10 - (checkDigit % 10)
	isbn13 := isbn12 + strconv.Itoa(checkDigit)

	return isbn13, nil
}

func ConvertISBN13ToISBN10(isbn13 string) (string, error) {
	isbn13 = strings.ReplaceAll(isbn13, "-", "")
	isbn13 = strings.ReplaceAll(isbn13, " ", "")

	if len(isbn13) != 13 {
		return "", errors.New("Invalid ISBN-13")
	}

	// Check if the ISBN-13 starts with "978" or "979"
	if !strings.HasPrefix(isbn13, "978") && !strings.HasPrefix(isbn13, "979") {
		return "", errors.New("Invalid ISBN-13 prefix")
	}

	// Remove the prefix and check digit
	isbn12 := isbn13[3:12]

	// Calculate the check digit for ISBN-10
	checkDigit := 0
	for i, c := range isbn12 {
		digit, err := strconv.Atoi(string(c))
		if err != nil {
			return "", err
		}
		checkDigit += (i + 1) * digit
	}
	checkDigit = checkDigit % 11
	isbn10 := isbn12 + strconv.Itoa(checkDigit)
	if checkDigit == 10 {
		isbn10 = isbn12 + "X"
	}

	return isbn10, nil
}
