package events

import (
	"context"
	"fmt"
	"os"

	"github.com/dpolimeni/fiber_app/common"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

var DbClient = common.GetDB()

// getEvents godoc
// @Summary Get all events.
// @Description Get events from db.
// @Tags Events
// @Produce json
// @Param Authorization header string true "Authorization" Default(Bearer)
// @Success 200 {object} map[string]interface{}
// @BasePath /api/v1
// @Router /api/v1/events [get]
func GetEvents(c *fiber.Ctx) error {
	// Make a query with the ent client
	events, err := DbClient.Events.Query().All(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println(events)
	// convert events to json to return on the api
	// return the JSON
	return c.JSON(events)
}

// AddEvent godoc
// @Summary Add a new event to the db.
// @Description Add event to db.
// @Tags Events
// @Accept application/json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @BasePath /api/v1
// @Router /api/v1/events [post]
// @Param Authorization header string true "Authorization" Default(Bearer)
// @Param NewEvent body NewEvent true "New Event to add"
func AddEvent(c *fiber.Ctx) error {
	newEvent := new(NewEvent)
	if err := c.BodyParser(newEvent); err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}
	event, err := DbClient.Events.Create().
		SetName(newEvent.Name).
		SetCapacity(newEvent.Capacity).
		SetDescription(newEvent.Description).
		Save(context.Background())
	if err != nil {
		error_msg := fmt.Sprintf("Error creating event: %s", err)
		return c.Status(fiber.StatusBadRequest).JSON(error_msg)
	}
	event_response := EventResponse{
		Msg: fmt.Sprintf("Event with name %s, capacity %d and description %s created", event.Name, event.Capacity, event.Description),
	}
	return c.Status(fiber.StatusCreated).JSON(event_response)
}

func AddEventsRoutes(app *fiber.App) {
	v1 := app.Group("/api/v1")
	v1.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(os.Getenv("SECRET_KEY"))},
	}))
	v1.Get("/events", GetEvents)
	v1.Post("/events", AddEvent)
}
