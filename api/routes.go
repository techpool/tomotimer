package api

import (
	"github.com/gofiber/fiber/v2"
)

func Api(app *fiber.App) {

	api := app.Group("/api")
	api.Use("/todo", func(c *fiber.Ctx) error {
		return c.Send([]byte("/todo API"))
	})

}
