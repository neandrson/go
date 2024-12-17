package main

import "fmt"

type Book struct {
	Title  string
	Author string
	Year   int
	Genre  string
}

func NewBook(title, author string, year int, genre string) *Book {
	return &Book{
		Title:  title,
		Author: author,
		Year:   year,
		Genre:  genre,
	}
}

func main() {
	book := NewBook("Harry Potter", "J.K. Rowling", 1997, "Fantasy")
	fmt.Println("Title:", book.Title)
	fmt.Println("Author:", book.Author)
	fmt.Println("Year:", book.Year)
	fmt.Println("Genre:", book.Genre)
}
