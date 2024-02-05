package people

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/dpolimeni/fiber_app/common"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

var DbClient = common.GetDB()

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
	// Make a query with the ent client
	users, err := DbClient.User.Query().All(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println(users)
	// convert users to json to return on the api
	// return the JSON
	return c.JSON(users)
}

// getPerson godoc
// @Summary Get a single person.
// @Description Get person from db.
// @Tags People
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @BasePath /api/v1/people
// @Router /api/v1/person/{id} [get]
// @Param id path int true "Person ID"
func GetPerson(c *fiber.Ctx) error {
	person_id := c.Params("id")
	number, _ := strconv.Atoi(person_id)
	user, err := DbClient.User.Get(context.Background(), number)
	if err != nil {
		formattedString := fmt.Sprintf("No user with id %d", number)
		return c.Status(fiber.StatusBadRequest).JSON(formattedString) // , fiber.StatusBadRequest
	}
	return c.Status(fiber.StatusAccepted).JSON(user)
}

// addPerson godoc
// @Summary Add a person on DB.
// @Description Add a person on DB.
// @Tags People
// @Accept application/json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @BasePath /api/v1/people
// @Router /api/v1/person [post]
// @Param person body Person true "Person"
func AddPerson(c *fiber.Ctx) error {

	var p Person

	// Parse the request body and bind it to the Person struct
	if err := c.BodyParser(&p); err != nil {
		// Handle parsing error
		return c.Status(fiber.StatusInternalServerError).SendString("Error parsing the request body")
	}
	// Now, p will contain the values from the request body
	fmt.Println(p)

	// Create a new person entity with the values from the request body
	person, err := DbClient.User.Create().SetAge(p.Age).SetName(p.Name).Save(context.Background())
	if err != nil {
		// Handle error
		error_string := fmt.Sprintf("Error creating the person: %s", err)
		return c.Status(fiber.StatusInternalServerError).SendString(error_string)
	}
	return c.SendString("Person with name " + person.Name + " added to the database!")
}

// deletePerson godoc
// @Summary Delete a person.
// @Description Delete a person on DB.
// @Tags People
// @Accept application/json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @BasePath /api/v1/people
// @Router /api/v1/person/{id} [delete]
// @Param id path int true "Person ID"
func DeletePerson(c *fiber.Ctx) error {
	// Get the person id from the url
	person_id := c.Params("id")
	number, _ := strconv.Atoi(person_id)
	// Delete the person from the database
	err := DbClient.User.DeleteOneID(number).Exec(context.Background())
	if err != nil {
		// Handle error
		error_string := fmt.Sprintf("Error deleting the person: %s", err)
		return c.Status(fiber.StatusInternalServerError).SendString(error_string)
	}
	// Return a success message
	return c.SendString("Deleted Person with id " + person_id)
}

func SetupRoutes(app *fiber.App) fiber.Router {
	v1 := app.Group("/api/v1")
	v1.Get("/people", jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(os.Getenv("SECRET_KEY"))},
	}), GetPeople)
	v1.Get("/person/:id<int>", GetPerson)
	v1.Post("/person", AddPerson)
	v1.Delete("/person/:id", DeletePerson)
	return v1
}
