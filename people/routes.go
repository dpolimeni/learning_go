package people

import (
	"fmt"

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
// @Router /api/v1/people/{id} [get]
// @Param id path int true "Person ID"
func GetPerson(c *fiber.Ctx) error {
	person_id := c.Params("id")
	return c.SendString(person_id)
}

type Person struct {
	Name string `json:"name" xml:"name" form:"name"`
	Pass string `json:"pass" xml:"pass" form:"pass"`
}

// addPerson godoc
// @Summary Add a person on DB.
// @Description Add a person on DB.
// @Tags People
// @Accept application/json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @BasePath /api/v1/people
// @Router /api/v1/people [post]
// @Param person body Person true "Person"
func AddPerson(c *fiber.Ctx) error {
	// Create an instance of Person to store the parsed data
	var p Person

	// Parse the request body and bind it to the Person struct
	if err := c.BodyParser(&p); err != nil {
		// Handle parsing error
		return c.Status(fiber.StatusBadRequest).SendString("Bad Request")
	}

	// Now, p will contain the values from the request body
	fmt.Println(p)

	return c.SendString("Person with name " + p.Name + " added to the database!")
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
	v1.Get("/people/:id<int>", GetPerson)
	v1.Post("/people", AddPerson)
	v1.Delete("/people/:id", DeletePerson)
	return v1
}
