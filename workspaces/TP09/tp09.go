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
	"Belgium", ///
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

var NewWorldCountries []string

func main() {
	printSliceInfo(worldCountries, "worldCountries")

	europeanCountries := worldCountries[:9]
	printSliceInfo(europeanCountries, "europeanCountries")

	africanCountries := worldCountries[9:]
	printSliceInfo(africanCountries, "africanCountries")

	worldCountries := append(worldCountries, americanCountries...)

	printSliceInfo(worldCountries, "worldCountries")
}

func printSliceInfo(slice []string, name string) {
	fmt.Println(name)
	fmt.Printf("\t%v\n", slice)
	fmt.Printf("\tLen: %d, Cap: %d\n\n", len(slice), cap(slice))
}
