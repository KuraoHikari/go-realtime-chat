package main

import (
	"fmt"
	"log"
	"server/db"
)

func main() {
	_, err := db.NewDatabase()
	if err != nil {
		log.Fatalf("could not initialize connection: %s", err)
	}

	fmt.Println("maybe connected")
}