package main

import (
	"formation-go/TP04/database"
)

func main() {
	database.DisplayPersons()
	ken := database.GetPerson("Ken")
	ken.Age++
	// or (*ken).Age++
	// It would be necessary if ken were simple type variable
	database.DisplayPersons()
}
