package main

import (
	"fmt"
)

type book struct {
	title  string
	author person
	year   int
}

type person struct{} // implement this

var bookList = []book{
	{"Computing Machinery and Intelligence", "Alan Turing", 1950},
	{"Domain-Driven Design: Tackling Complexity in the Heart of Software", "Eric Evans", 2003},
	{"Hackers, Heroes of the Computer Revolution", "Steven Levy", 1984},
}

func main() {
	selectedAuthor := "Eric Evans"
	var selectedBooks []book
	for _, b := range bookList {
		if isBookAuthor(b, selectedAuthor) {
			selectedBooks = append(selectedBooks, b)
		}
	}
	fmt.Printf("%+v\n", selectedBooks)
}

func isBookAuthor(b book, author string) bool {
	return b.author == author
}
