package main

import (
	"log"
	"os"

	"formation-go/TPnet/server"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "0"
	}

	if err := server.ListenAndServe(":" + port); err != nil {
		log.Fatal(err)
	}
}
