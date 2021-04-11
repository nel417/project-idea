package main

import (
	// "fmt"
	"log"

	// "github.com/gofiber/basicauth"
	"github.com/gofiber/fiber"
	"github.com/gofiber/template/html"
)

func main() {
	//moodboard images
	x := make([]string, 11)
	x[0] = "https://hips.hearstapps.com/hmg-prod.s3.amazonaws.com/images/540fe1b392f2f-03-tcx-guy-celebrating-retro-tennis-gear-0812-xl-lg-1529073447.jpg?crop=0.862xw:1.00xh;0.0374xw,0&resize=480:*"
	x[1] = "https://i.pinimg.com/564x/e1/a3/58/e1a35879d8fc23371df151d3850bce7c.jpg"
	x[2] = "https://upload.wikimedia.org/wikipedia/en/6/63/Right_Now%2C_Wrong_Then_%28poster%29.jpg"
	x[3] = "https://asianwiki.com/images/d/da/Right_Now%2C_Wrong_Then-p1.jpg"
	x[4] = "https://mir-s3-cdn-cf.behance.net/project_modules/1400/8146e165435111.5af41549b7840.jpg"
	x[5] = "https://shop.wilson.com/en-gb/media/catalog/product/cache/146/image/1800x/040ec09b1e35df139433887a97daa66f/5/c/5c260e027a31c428e87b24990e790838dbe88555_Triniti_alt_2.jpg"
	x[6] = "https://cdn11.bigcommerce.com/s-vq9b9x8mdr/products/154/images/756/calabasas_tennis_longsleeve__28797.1590614602.386.513.jpg?c=2"
	x[7] = "https://assets.vogue.com/photos/5b4771e676530c3d905c5491/master/w_1600%2Cc_limit/03-tennis.jpg"
	x[8] = "https://www.heddels.com/wp-content/uploads/2013/09/baldwin-denim-x-suit-supply-13-oz-blue-jeans-collaboration.jpg"
	x[9] = "https://cnet4.cbsistatic.com/img/d3F8xlb7Y2W-oyk5SpK5lfDuwMQ=/940x0/2019/11/27/1f62650e-4c18-47ab-aece-1b82824c69fc/stevejobs.jpg"
	x[10] = "https://globalassets.starbucks.com/assets/26495f2a8b644fe8960dd74e1891b7b7.jpg?impolicy=1by1_wide_1242"
	//end mooboard images

	engine := html.New("./views/", ".html")

	app := fiber.New(&fiber.Settings{
		Views: engine,
	})

	//Rendered assets
	app.Get("/", func(c *fiber.Ctx) {
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
			"Hello":       "Hello",
			"OS":          "Open Source Contributions",
			"Talk":        "Most Recent Talk On Open Source",
			"Reads":       "Favorite Books",
			"Software":    "Software / ML (click for demos)",
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
	Port :=8080
	log.Fatal(app.Listen(Port))

	// Get the PORT from heroku env
	port := os.Getenv("PORT")

	// Verify if heroku provided the port or not
	if os.Getenv("PORT") == "" {
		port = "8080"
	}

	// Start server on http://${heroku-url}:${port}
	log.Fatal(app.Listen(":" + port))
}
