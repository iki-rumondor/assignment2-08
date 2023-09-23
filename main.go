package main

import (
	"log"

	"github.com/iki-rumondor/assignment2-GLNG-KS-08-08/database"
)

func main() {
	_, err := database.InitPostgresDb()

	if err != nil {
		log.Fatal(err.Error())
	}

}
