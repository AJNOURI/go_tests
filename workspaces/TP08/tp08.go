package main

import (
	"fmt"
)

type person struct { // implement this
	firstname string
	lastname  string
}

type book struct {
	title  string
	author person
	year   int
}

var bookList = []book{
	{"Computing Machinery and Intelligence", person{"Alan", "Turing"}, 1950},
	{"Domain-Driven Design: Tackling Complexity in the Heart of Software", person{"Eric", "Evans"}, 2003},
	{"Hackers, Heroes of the Computer Revolution", person{"Steven", "Levy"}, 1984},
}

func main() {
	selectedAuthor := person{"Eric", "Evans"}
	var selectedBooks []book
	for _, b := range bookList {
		if isBookAuthor(b, selectedAuthor) {
			selectedBooks = append(selectedBooks, b)
		}
	}
	fmt.Printf("%+v\n", selectedBooks)
}

func isBookAuthor(b book, author person) bool {
	return b.author == author
}
