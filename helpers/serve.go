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
		return c.SendFile("./build/dev/index.html")
	})

	app.Get("/searching", func(c *fiber.Ctx) error {
		return c.SendFile("./build/dev/searching.html")
	})

	app.Static("/", "./build/dev")

	app.Get("/blog/:slug", func(c *fiber.Ctx) error {
		return c.SendFile("./build/dev/blog/" + c.Params("slug") + ".html")
	})

	app.Listen(":3000")
}
