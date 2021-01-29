package main

import (
	"log"

	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware"
	"github.com/pascaloseko/go-rest-fibre-postgres/database"
	"github.com/pascaloseko/go-rest-fibre-postgres/router"

	_ "github.com/lib/pq"
)

func main() { // entry point to our program

	// Connect to database
	if err := database.Connect(); err != nil {
		log.Fatal(err)
	}

	app := fiber.New() // call the New() method - used to instantiate a new Fiber App

	app.Use(middleware.Logger())

	router.SetupRoutes(app)

	app.Listen(3000) // listen/Serve the new Fiber app on port 3000

}
