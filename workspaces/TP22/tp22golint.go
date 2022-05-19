package main

import (
	"errors"
	. "fmt"
)

var SomeError = errors.New("Capitalised error message")

type unexported int
type Exported int

func (this unexported) Foo() {}

func (oneName Exported) Foo() {
}

func (anotherName Exported) Bar() {
	Println("Hi")
}
