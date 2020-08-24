package main

import "github.com/gofiber/fiber"

func main() {

	app := fiber.New()
	app.Static("/", "./public")

	app.Get("/", func(c *fiber.Ctx) {
		if err := c.Render("index", fiber.Map{}); err != nil {
			c.Next(err)
		}
	})

	app.Get("/about", func(c *fiber.Ctx) {
		c.Send("about")
	})

	app.Listen(3000)
}
