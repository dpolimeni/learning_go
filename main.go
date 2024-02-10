package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/dpolimeni/fiber_app/auth"

	"github.com/dpolimeni/fiber_app/common"
	_ "github.com/dpolimeni/fiber_app/docs"
	"github.com/dpolimeni/fiber_app/ent"
	"github.com/dpolimeni/fiber_app/people"
	"github.com/dpolimeni/fiber_app/reservations"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"
)

var DbClient = common.GetDB()

// @title Fiber Example API
// @version 2.0
// @description This is a sample swagger for Fiber
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
// @name <YourTokenName>
func main() {

	// Create a new Fiber instance
	app := fiber.New()
	common.LoadEnv()
	app.Use(logger.New())
	app.Get("/swagger/*", swagger.HandlerDefault)
	people.SetupRoutes(app)
	auth.SetUpAuthRoutes(app)
	reservations.SetUpReservationsRoutes(app)
	app.Get("/", HealthCheck)
	password := os.Getenv("password")
	connection := fmt.Sprintf("host=localhost port=5432 user=postgres dbname=gotest password=%s sslmode=disable", password)
	client, err := ent.Open("postgres", connection)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	app.Listen(":8080")
}

// HealthCheck godoc
// @Summary Show the status of server.
// @Description Get test on base path.
// @Tags Root Base
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router / [get]
func HealthCheck(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}
