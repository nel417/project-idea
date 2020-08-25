package main

import (
	"fmt"
	"log"

	"github.com/gofiber/basicauth"
	"github.com/gofiber/fiber"
	"github.com/gofiber/template/html"
)

func main() {
	//moodboard images
	x := make([]string, 5)
	x[0] = "https://hips.hearstapps.com/hmg-prod.s3.amazonaws.com/images/540fe1b392f2f-03-tcx-guy-celebrating-retro-tennis-gear-0812-xl-lg-1529073447.jpg?crop=0.862xw:1.00xh;0.0374xw,0&resize=480:*"
	x[1] = "https://i.pinimg.com/564x/e1/a3/58/e1a35879d8fc23371df151d3850bce7c.jpg"
	x[2] = "https://upload.wikimedia.org/wikipedia/en/6/63/Right_Now%2C_Wrong_Then_%28poster%29.jpg"
	x[3] = "https://asianwiki.com/images/d/da/Right_Now%2C_Wrong_Then-p1.jpg"
	x[4] = "https://mir-s3-cdn-cf.behance.net/project_modules/1400/8146e165435111.5af41549b7840.jpg"
	//end mooboard images

	engine := html.New("./views/", ".html")

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
			"Title":       "Nick Landreville Studios",
			"MailingList": "Join Our Mailing List",
		}); err != nil {
			c.Next(err)
		}
	})

	app.Get("/contact", func(c *fiber.Ctx) {
		// Static fn
		// c.SendFile("./public/about.html")
		// Render fn
		if err := c.Render("contact", fiber.Map{
			"Title":       "Nick Landreville Studios",
			"ContactInfo": "n1ck@hey.com",
			"Location":    "Springfield, MO",
			"MailingList": "Join Our Mailing List",
		}); err != nil {
			c.Next(err)
		}
	})

	// 	THIS IS WHERE THE BLOG BEGINS!!!!!!!!
	app.Get("/blog", func(c *fiber.Ctx) {
		// Static fn
		// c.SendFile("./public/about.html")
		// Render fn
		if err := c.Render("blog", fiber.Map{
			"Title": "Nick Landreville Studios",
		}); err != nil {
			c.Next(err)
		}
	})

	app.Get("/a-post", func(c *fiber.Ctx) {
		// Static fn
		// c.SendFile("./public/about.html")
		// Render fn
		if err := c.Render("posts/dummy", fiber.Map{
			"Title": "Nick Landreville Studios",
		}); err != nil {
			c.Next(err)
		}
	})
	// THIS IS WHERE THE BLOG BEGINS!!!!!!!!!!!!

	// Static assets ( .css, .js etc )
	app.Static("/", "./public")
	log.Fatal(app.Listen(3000))
}
