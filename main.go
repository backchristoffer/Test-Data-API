package main

import (
	"github.com/gofiber/fiber/v2"
)

func healthCheck(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusOK)
}

func main() {
	app := fiber.New()
	app.Get("/healthcheck", healthCheck)
	app.Listen(":3000")
}
