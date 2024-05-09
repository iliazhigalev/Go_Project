package main

import (
	"HTTP-REST-API/database"
	"HTTP-REST-API/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {
	database.ConnectDb()

	engine := html.New("./views", ".html") // указываем директорию для поиска шаблонов и расширение файлов.
	app := fiber.New(fiber.Config{
		Views:       engine,
		ViewsLayout: "layouts/main",
	})

	setupRoutes(app)

	app.Static("/", "./public")

	app.Use(handlers.NotFound) // обработчик, который будет вызываться для всех запросов, если не будет найден подходящий маршрут.

	app.Listen(":3000")
}
