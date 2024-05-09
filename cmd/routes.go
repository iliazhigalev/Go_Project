package main

import (
	"HTTP-REST-API/handlers"

	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {

	app.Get("/", handlers.ListFacts)

	app.Get("/fact", handlers.NewFactView)

	app.Post("/fact", handlers.CreateFact)

	app.Get("/fact/:id", handlers.ShowFact)
}
