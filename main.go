package main

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"

	"github.com/techpool/tomotimer/api"
	"github.com/techpool/tomotimer/utils"
)

func main() {
	_, err := utils.LoadConfig(".")
	if err != nil {
		panic("Cannot start application without loading configurations")
	}
	app := fiber.New()

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(http.StatusOK).
			JSON(map[string]interface{}{
				"health": "ok",
				"status": http.StatusOK,
			})
	})
	api.Api(app)

	app.Listen(":" + viper.GetString(utils.PORT))
}
