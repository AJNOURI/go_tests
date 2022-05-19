package main

import (
	"fmt"
	"log"

	owm "github.com/briandowns/openweathermap"
)

func main() {
	w, err := owm.NewCurrent("C", "fr", "bfcc1da7cd06d87153e372c619ab13fb")
	if err != nil {
		log.Fatalln(err)
	}

	for _, city := range []string{"Paris", "London", "Rome"} {
		if err := w.CurrentByName(city); err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("%s: %s, %.1fC\n", w.Name, w.Weather[0].Description, w.Main.Temp)
	}
}
