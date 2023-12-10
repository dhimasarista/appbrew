package main

import (
	"appbrew/app/routes"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/mustache/v2"
)

func main() {
	engine := mustache.New("./views", ".mustache")
	app := fiber.New(
		fiber.Config{
			Views: engine,
		},
	)
	app.Static("/", "./public")

	// Routes
	routes.OrdersRoutes(app)

	log.Fatal(app.Listen(":8900"))
}
