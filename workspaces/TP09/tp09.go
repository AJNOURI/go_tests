package main

import "fmt"

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

var americanCountries = []string{
	"United States",
	"Mexico",
	"Canada",
	"Brazil",
	"Colombia",
	"Argentina",
	"Peru",
	"Venezuela",
	"Chile",
}

func main() {
    printSliceInfo(worldCountries, "worldCountries")

	europeanCountries := // Code this
	printSliceInfo(europeanCountries, "europeanCountries")

	africanCountries := // Code this
	printSliceInfo(africanCountries, "africanCountries")

	for _, c := range americanCountries {
		// Code this: Add c to worldcountries
	}
	printSliceInfo(worldCountries, "worldCountries")
}

func printSliceInfo(slice []string, name string) {
	fmt.Println(name)
	fmt.Printf("\t%v\n", slice)
	fmt.Printf("\tLen: %d, Cap: %d\n\n", len(slice), cap(slice))
}
