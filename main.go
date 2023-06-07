package main

import (
	"github.com/gofiber/fiber/v2"
)

func healthCheck(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusOK)
}

func jsondata(c *fiber.Ctx) error {
	type jsondata struct {
		Name  string
		Age   uint8
		Level uint8
	}
	data := jsondata{
		Name:  "John",
		Age:   22,
		Level: 99,
	}
	return c.JSON(data)
}

func main() {
	app := fiber.New()
	app.Get("/healthcheck", healthCheck)
	app.Get("/json", jsondata)
	app.Listen(":3000")
}
