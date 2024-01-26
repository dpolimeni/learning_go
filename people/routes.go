package people

import (
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Group("/api/v1/people")
	app.Get("/api/v1/people", getPeople)
	app.Get("/api/v1/people/:id", getPerson)
	app.Post("/api/v1/people", addPerson)
	app.Delete("/api/v1/people/:id", deletePerson)
}

// getPeople godoc
// @Summary Get all people.
// @Description Get people from db.
// @Tags Root Base
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @BasePath /api/v1/people
func getPeople(c *fiber.Ctx) error {
	return c.SendString("All People")
}

func getPerson(c *fiber.Ctx) error {
	return c.SendString("Single Person")
}

func addPerson(c *fiber.Ctx) error {
	return c.SendString("Add Person")
}

func deletePerson(c *fiber.Ctx) error {
	return c.SendString("Delete Person")
}
