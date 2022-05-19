package database

import "fmt"

// Person contains information about a person
type Person struct {
	Name string
	Age  int
}

var ken = Person{"Ken Thompson", 74}
var rob = Person{"Rob Pike", 62}

// GetPerson returns a person
func GetPerson(name string) *Person {
	switch name {
	case "Ken":
		return &ken
	case "Rob":
		return &rob
	}
	panic("unknow person")
}

// DisplayPersons displays all the persons off the database
func DisplayPersons() {
	fmt.Printf("%v\n", ken)
	fmt.Printf("%v\n", rob)
}
