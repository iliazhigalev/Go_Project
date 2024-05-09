package handlers

import (
	"HTTP-REST-API/database"
	"HTTP-REST-API/models"

	"github.com/gofiber/fiber/v2"
)

func ListFacts(c *fiber.Ctx) error {
	facts := []models.Fact{}
	database.DB.Db.Find(&facts)

	return c.Render("index", fiber.Map{
		"Title":    "Div Rhino Trivia",
		"Subtitle": "Facts for funtimes with friends",
		"Facts":    facts,
	})
}

func NewFactView(c *fiber.Ctx) error {
	return c.Render("new", fiber.Map{
		"Title":    "New Facts",
		"Subtitle": "Add a cool facts!",
	})

}

func CreateFact(c *fiber.Ctx) error {
	fact := new(models.Fact)
	if err := c.BodyParser(fact); err != nil {
		return NewFactView(c)
	}
	// если на этом этапе будет возвращена ошибка возвратим форму для создания нового факта
	result := database.DB.Db.Create(&fact)
	if result.Error != nil {
		return NewFactView(c)
	}
	return ListFacts(c)

}

func ShowFact(c *fiber.Ctx) error {
	fact := models.Fact{}
	id := c.Params("id")

	result := database.DB.Db.Where("id = ?", id).First(&fact)
	if result.Error != nil {
		return NotFound(c)
	}

	return c.Render("show", fiber.Map{
		"Title": "Single Fact",
		"Fact":  fact,
	})
}

func NotFound(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotFound).SendFile("./public/404.html")
}
