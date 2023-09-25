package main

import (
	"log"

	"github.com/iki-rumondor/assignment2-08/applications/services"
	"github.com/iki-rumondor/assignment2-08/database"
	"github.com/iki-rumondor/assignment2-08/handlers"
	"github.com/iki-rumondor/assignment2-08/routers"
)

var PORT string = ":8080"

func main() {
	db, err := database.InitPostgresDb()

	if err != nil {
		log.Fatal(err.Error())
		return
	}

	orderServices := services.NewOrderServicve(db)
	orderHandler := handlers.OrderHandler{Services: orderServices}

	handlers := handlers.Handlers{
		OrderHandler: &orderHandler,
	}

	routers.StartServer(&handlers).Run(PORT)
}
