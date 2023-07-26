package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nilu2039/tokiio-server-go/handlers"
	"github.com/nilu2039/tokiio-server-go/initializers"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectToDB()
}

func main() {
	app := fiber.New()

	api := app.Group("/api")

	v1 := api.Group("v1")

	v1.Get("/top-airing", handlers.TopAiring)

	app.Listen(":4000")
}
