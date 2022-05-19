package main

import (
	"formation-go/TP04/database"
)

func main() {
	database.DisplayPersons()
	ken := database.GetPerson("Ken")
	ken.Age++
	database.DisplayPersons()
}
