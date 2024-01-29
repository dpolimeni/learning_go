package people

import (
	"fmt"

	"database/sql"

	"os"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
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
	const (
		host   = "localhost"
		port   = 5432
		user   = "postgres"
		dbname = "gotest"
	)
	var password = os.Getenv("password")
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT name, surname FROM pippo")
	if err != nil {
		fmt.Println("Error executing query:", err)
	}
	defer rows.Close()

	for rows.Next() {
		var name, surname string
		err := rows.Scan(&name, &surname)
		if err != nil {
			fmt.Println("Error scanning row:", err)
			return c.SendString("Error scanning rows:")
		}
		fmt.Printf("name: %s, surname: %s\n", name, surname)
	}

	if err := rows.Err(); err != nil {
		fmt.Println("Error retrieving rows:", err)
		return c.SendString("Error retrieving rows:")
	}

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
