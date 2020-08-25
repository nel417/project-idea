package main

import (
	"fmt"
	"log"

	"github.com/gofiber/basicauth"
	"github.com/gofiber/fiber"
	"github.com/gofiber/template/html"
)

func main() {
	// moodboard images
	x := make([]string, 3)
	x[0] = "https://hips.hearstapps.com/hmg-prod.s3.amazonaws.com/images/540fe1b392f2f-03-tcx-guy-celebrating-retro-tennis-gear-0812-xl-lg-1529073447.jpg?crop=0.862xw:1.00xh;0.0374xw,0&resize=480:*"
	x[1] = "https://hips.hearstapps.com/hmg-prod.s3.amazonaws.com/images/540fe1b392f2f-03-tcx-guy-celebrating-retro-tennis-gear-0812-xl-lg-1529073447.jpg?crop=0.862xw:1.00xh;0.0374xw,0&resize=480:*"
	x[2] = "https://hips.hearstapps.com/hmg-prod.s3.amazonaws.com/images/540fe1b392f2f-03-tcx-guy-celebrating-retro-tennis-gear-0812-xl-lg-1529073447.jpg?crop=0.862xw:1.00xh;0.0374xw,0&resize=480:*"
	//end mooboard images
	engine := html.New("./views", ".html")
	app := fiber.New(&fiber.Settings{
		Views: engine,
	})

	cfg := basicauth.Config{
		Users: map[string]string{
			"Nick":  "Sharks12!",
			"admin": "123456",
		},
	}
	app.Use(basicauth.New(cfg))
	//Rendered assets
	app.Get("/", func(c *fiber.Ctx) {
		username := c.Locals("username").(string)
		password := c.Locals("password").(string)
		fmt.Println(username, password)
		if err := c.Render("index", fiber.Map{
			"Title": "Nick Landreville Studios",
			"loopy": x,
		}); err != nil {
			c.Next(err)
		}
	})
	app.Get("/about", func(c *fiber.Ctx) {
		// Static fn
		// c.SendFile("./public/about.html")
		// Render fn
		if err := c.Render("about", fiber.Map{
			"AboutTitle": "About Nick Landreville Studios",
		}); err != nil {
			c.Next(err)
		}
	})

	app.Get("/contact", func(c *fiber.Ctx) {
		// Static fn
		// c.SendFile("./public/about.html")
		// Render fn
		if err := c.Render("contact", fiber.Map{
			"ContactInfo": "n1ck@hey.com",
		}); err != nil {
			c.Next(err)
		}
	})
	// Static assets ( .css, .js etc )
	app.Static("/", "./public")
	log.Fatal(app.Listen(3000))
}
