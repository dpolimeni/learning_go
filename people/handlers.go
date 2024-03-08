package people

import (
	"context"
	"fmt"
	"strconv"

	"github.com/dpolimeni/fiber_app/ent"
	"github.com/dpolimeni/fiber_app/ent/user"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/mongo"
)

type PeopleHandler struct {
	DbClient    *ent.Client
	MongoClient *mongo.Client
}

// getPeople godoc
// @Summary Get all people.
// @Description Get people from db.
// @Tags People
// @Produce json
// @Param Authorization header string true "Authorization" Default(Bearer )
// @Success 200 {object} map[string]interface{}
// @BasePath /api/v1/people
// @Router /api/v1/people [get]
func (handler PeopleHandler) GetPeople(c *fiber.Ctx) error {
	// Make a query with the ent client
	users, err := handler.DbClient.User.Query().All(context.Background())
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
// @Param Authorization header string true "Authorization" Default(Bearer )
func (handler PeopleHandler) GetPerson(c *fiber.Ctx) error {
	person_id := c.Params("id")
	number, _ := strconv.Atoi(person_id)
	user, err := handler.DbClient.User.Get(context.Background(), number)
	if err != nil {
		formattedString := fmt.Sprintf("No user with id %d", number)
		return c.Status(fiber.StatusBadRequest).JSON(formattedString) // , fiber.StatusBadRequest
	}
	return c.Status(fiber.StatusAccepted).JSON(user)
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
// @Param Authorization header string true "Authorization" Default(Bearer )
func (handler PeopleHandler) DeletePerson(c *fiber.Ctx) error {
	// Get the person id from the url
	person_id := c.Params("id")
	number, _ := strconv.Atoi(person_id)
	// Delete the person from the database
	err := handler.DbClient.User.DeleteOneID(number).Exec(context.Background())
	if err != nil {
		// Handle error
		error_string := fmt.Sprintf("Error deleting the person: %s", err)
		return c.Status(fiber.StatusInternalServerError).SendString(error_string)
	}
	// Return a success message
	return c.SendString("Deleted Person with id " + person_id)
}

// Update person Doc
// @Summary Update a person data
// @Description Update a person data (age/email)
// @Tags People
// @Accept application/json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @BasePath /api/v1
// @Router /api/v1/person/{username} [put]
// @Param username path string true "Person username"
// @Param person body Person true "Person data to update"
// @Param Authorization header string true "Authorization" Default(Bearer )
func (handler PeopleHandler) UpdatePerson(c *fiber.Ctx) error {
	// Get the username from the person username
	username := c.Params("username")

	// Get user data from DB
	user, err := handler.DbClient.User.Query().Where(user.UsernameEQ(username)).First(context.Background())
	if err != nil {
		// Handle error
		error_string := fmt.Sprintf("The username provided does not exists: %s", err)
		return c.Status(fiber.StatusInternalServerError).JSON(error_string)
	}
	// Verify if the username in the token is the same to the username in the url
	// if not return an error
	userToken := c.Locals("user").(*jwt.Token)
	fmt.Println(userToken)
	fmt.Println(c.Locals("user"))
	token_username := userToken.Claims.(jwt.MapClaims)["username"]
	if token_username != username {
		return c.Status(fiber.StatusUnauthorized).JSON("You are not authorized to update this user")
	}
	// Get the new age and email from the request
	var p Person
	if err := c.BodyParser(&p); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}
	// Update the user data
	_, err = handler.DbClient.User.UpdateOne(user).SetAge(p.Age).SetName(p.Name).Save(context.Background())
	if err != nil {
		// Handle error
		error_string := fmt.Sprintf("Error updating the person: %s", err)
		return c.Status(fiber.StatusInternalServerError).JSON(error_string)
	}
	// Return a success message
	response := map[string]interface{}{
		"message": "User updated successfully",
		"person":  p,
	}
	return c.JSON(response)

}

// Insert person Doc
// @Summary Insert a new person
// @Description Create Person record on MONGO
// @Tags People
// @Accept application/json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @BasePath /api/v1
// @Router /api/v1/person/new [post]
func (handler PeopleHandler) AddPerson(c *fiber.Ctx) error {
	// Try fake mongo creation of a document
	// Generate random big integer id
	//id := uuid.New()
	collection := handler.MongoClient.Database("test").Collection("people")
	person := Person{
		Name: "Diego",
		Age:  25,
	}

	_, err := collection.InsertOne(context.TODO(), person)
	if err != nil {
		fmt.Println("ERroR")
		fmt.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}
	return c.SendString("Person added successfully")
}
