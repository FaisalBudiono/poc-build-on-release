package main

import (
	"building/internal/cred"
	"log"
)

func main() {
	res, err := cred.Something()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Username found: %s", res.Username)
}
