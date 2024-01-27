package people

import (
	"github.com/gofiber/fiber/v2"
)

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

// getPerson godoc
// @Summary Get a single person.
// @Description Get person from db.
// @Tags Root Base
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @BasePath /api/v1/people
func getPerson(c *fiber.Ctx) error {
	return c.SendString("Single Person")
}

// addPerson godoc
// @Summary Add a person on DB.
// @Description Add a person on DB.
// @Tags Root Base
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @BasePath /api/v1/people
func addPerson(c *fiber.Ctx) error {
	return c.SendString("Add Person")
}

// deletePerson godoc
// @Summary Delete a person.
// @Description Delete a person on DB.
// @Tags Root Base
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @BasePath /api/v1/people
func deletePerson(c *fiber.Ctx) error {
	return c.SendString("Delete Person")
}

func SetupRoutes(app *fiber.App) {
	v1 := app.Group("/api/v1")
	v1.Get("/people", getPeople)
	v1.Get("/people/:id", getPerson)
	v1.Post("/people", addPerson)
	v1.Delete("/people/:id", deletePerson)
}
