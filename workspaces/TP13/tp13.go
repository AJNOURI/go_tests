package main

import (
	"fmt"
	"sort"
)

type Developer struct {
	FirstName, LastName string
	Age                 uint
	Language            string
}

type DevelopersByAge []Developer

// FIXME DevelopersByAge must implement sort.Interface https://pkg.go.dev/sort#Interface

func main() {
	developers := []Developer{
		{"Roy", "Trenneman", 39, "Python"},
		{"Jen", "Barber", 32, "Excel"},
		{"Maurice", "Moss", 34, "Rust"},
	}

	sort.Sort(DevelopersByAge(developers))

	fmt.Println(developers)
}
