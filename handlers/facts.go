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
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	database.DB.Db.Create(&fact)
	return ConfirmationView(c)

}

func ConfirmationView(c *fiber.Ctx) error {

	return c.Render("confirmation", fiber.Map{
		"Title":    "New Facts",
		"Subtitle": "Add more wonderful facts to the list!",
	})

}
