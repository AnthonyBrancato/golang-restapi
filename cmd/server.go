package main

import (
	"github.com/AnthonyBrancato/golang-restapi/database"
	"github.com/gofiber/fiber/v3"
)

func main() {
	database.ConnectDb()
	app := fiber.New()

	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("live reloading")
	})

	app.Listen(":3000")
}
