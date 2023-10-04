package models

type Author struct {
	AuthorID   int    `json:"author_id"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	MiddleName string `json:"middle_name"`
}

type Book struct {
	BookID          int     `json:"book_id"`
	Title           string  `json:"title"`
	ISBN13          string  `json:"isbn13"`
	ISBN10          string  `json:"isbn10"`
	ListPrice       float64 `json:"list_price"`
	PublicationYear int     `json:"publication_year"`
	PublisherID     int     `json:"publisher_id"`
	ImageURL        string  `json:"image_url"`
	Edition         string  `json:"edition"`
}

type Publisher struct {
	PublisherID   int    `json:"publisher_id"`
	PublisherName string `json:"publisher_name"`
}

type BookInformation struct {
	Title           string   `json:"title"`
	Author          []string `json:"author"`
	ISBN13          string   `json:"isbn13"`
	ISBN10          string   `json:"isbn10"`
	PublicationYear int      `json:"publication_year"`
	Publisher       string   `json:"publisher"`
	Edition         string   `json:"edition"`
	ListPrice       float64  `json:"list_price"`
}

type AuthorBook struct {
	AuthorID int    `json:"author_id"`
	BookID   int    `json:"book_id"`
	BookName string `json:"book_name"`
}
