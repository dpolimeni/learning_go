package people

import (
	"github.com/gofiber/fiber/v2"
)

// getPeople godoc
// @Summary Get all people.
// @Description Get people from db.
// @Tags People
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @BasePath /api/v1/people
// @Router /api/v1/people [get]
func GetPeople(c *fiber.Ctx) error {
	return c.SendString("All People")
}

// getPerson godoc
// @Summary Get a single person.
// @Description Get person from db.
// @Tags People
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @BasePath /api/v1/people
// @Router /api/v1/people [get]
func GetPerson(c *fiber.Ctx) error {
	return c.SendString("Single Person")
}

// addPerson godoc
// @Summary Add a person on DB.
// @Description Add a person on DB.
// @Tags People
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @BasePath /api/v1/people
// @Router /api/v1/people [post]
func AddPerson(c *fiber.Ctx) error {
	return c.SendString("Add Person")
}

// deletePerson godoc
// @Summary Delete a person.
// @Description Delete a person on DB.
// @Tags People
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @BasePath /api/v1/people
// @Router /api/v1/people [delete]
func DeletePerson(c *fiber.Ctx) error {
	return c.SendString("Delete Person")
}

func SetupRoutes(app *fiber.App) fiber.Router {
	v1 := app.Group("/api/v1")
	v1.Get("/people", GetPeople)
	v1.Get("/people/:id", GetPerson)
	v1.Post("/people", AddPerson)
	v1.Delete("/people/:id", DeletePerson)
	return v1
}
