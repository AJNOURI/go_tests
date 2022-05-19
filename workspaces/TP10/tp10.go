package main

import (
	"fmt"
)

var worldCountries = []string{
	"Germany",
	"France",
	"United Kingdom",
	"Italy",
	"Spain",
	"Poland",
	"Romania",
	"Netherlands",
	"Belgium",
	"Nigeria",
	"Ethiopia",
	"Egypt",
	"Democratic Republic of the Congo",
	"South Africa",
	"Tanzania",
	"Kenya",
	"Algeria",
	"Sudan",
	"Uganda",
	"Morocco",
}

var worldCities = []string{
	"Berlin",
	"Paris",
	"London",
	"Rome",
	"Madrid",
	"Varsaw",
	"Bucharest",
	"Amsterdam",
	"Brussels",
	"Abuja",
	"Addis Ababa",
	"Cairo",
	"Kinshasa",
	"Pretoria",
	"Dodoma",
	"Nairobi",
	"Algiers",
	"Khartoum",
	"Kampala",
	"Rabat",
}


func main() {
    worldCapitals := make(map[string]string, len(worldCountries))
    
    for i, country := range worldCountries {
        worldCapitals[country] = 
    }
    
    for country, city := range worldCapitals {
        fmt.Println(country, " - ", city)
    }
}
