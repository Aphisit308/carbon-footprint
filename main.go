package main

import (
	"carbon-footprint/config"
	"carbon-footprint/controllers"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	config.SetupDB()
	controllers.SetupRoutes(app)
	log.Fatal(app.Listen(":8000"))
}
