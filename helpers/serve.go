package helpers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func HtmlServer() {
	app := fiber.New(fiber.Config{
		Prefork:               false,
		DisableStartupMessage: true,
	})

	app.Use(logger.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendFile("./build/index.html")
	})

	app.Get("/searching", func(c *fiber.Ctx) error {
		return c.SendFile("./build/searching/index.html")
	})

	app.Static("/", "./build")

	app.Get("/blog/:slug", func(c *fiber.Ctx) error {
		return c.SendFile("./build/blog/" + c.Params("slug") + "/index.html")
	})

	app.Listen(":3000")
}
