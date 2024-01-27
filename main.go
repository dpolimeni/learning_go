package main

import (
	"fmt"

	_ "github.com/dpolimeni/fiber_app/docs"
	"github.com/dpolimeni/fiber_app/people"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

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
func main() {

	fmt.Println("Hello World")

	app := fiber.New()
	people.SetupRoutes(app)

	app.Get("/swagger/*", swagger.HandlerDefault)
	app.Get("/", HealthCheck)

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
