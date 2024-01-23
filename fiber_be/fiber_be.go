package fiber_be

import "github.com/gofiber/fiber/v2"

func fiber_be() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(":3000")
}
