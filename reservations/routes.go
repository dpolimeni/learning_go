package reservations

import (
	"context"
	"fmt"
	"os"

	"github.com/dpolimeni/fiber_app/common"
	"github.com/dpolimeni/fiber_app/ent/reservations"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

var DbClient = common.GetDB()

// getReservations godoc
// @Summary Get all reservations.
// @Description Get reservations from db.
// @Tags Reservations
// @BasePath /api/v1
// @Router /api/v1/reservations [get]
// @Produce json
// @Param Authorization header string true "Authorization" Default(Bearer)
func GetReservations(c *fiber.Ctx) error {
	// Make a query with the ent client
	reservations, err := DbClient.Reservations.Query().All(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println(reservations)
	// convert reservations to json to return on the api
	// return the JSON
	return c.JSON(reservations)
}

// getReservation godoc
// @Summary Get a single reservation.
// @Description Get reservation from db.
// @Tags Reservations
// @BasePath /api/v1
// @Router /api/v1/reservations/{id} [get]
// @Accept */*
// @Produce json
// @Param Authorization header string true "Authorization" Default(Bearer)
// @Param id path string true "Reservation ID"
func GetReservation(c *fiber.Ctx) error {
	reservation_id := c.Params("id")
	reservation, err := DbClient.Reservations.Query().Where(reservations.ID(reservation_id)).Only(context.Background())
	if err != nil {
		formattedString := fmt.Sprintf("No reservation with id %v", reservation_id)
		return c.Status(fiber.StatusBadRequest).JSON(formattedString) // , fiber.StatusBadRequest
	}
	return c.Status(fiber.StatusAccepted).JSON(reservation)
}

func AddReservationsRoutes(app *fiber.App) {
	v1 := app.Group("/api/v1")
	v1.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(os.Getenv("SECRET_KEY"))},
	}))
	v1.Get("/reservations", GetReservations)
	v1.Get("/reservations/:id", GetReservation)
}
