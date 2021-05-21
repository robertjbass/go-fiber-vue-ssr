package main

import (
	"fmt"
	"log"
	"github.com/gofiber/fiber/v2"
)

func main() {
	fmt.Println("Hello World")
	app := fiber.New()

	/* // Express Example
	app.get('/', (req, res) => {
		res.send("Hello World")
	})
	*/
	app.Static("/", "./public")
    // => http://localhost:3000/js/script.js
    // => http://localhost:3000/css/style.css

	app.Get("/", func(c *fiber.Ctx) error {
        msg := fmt.Sprintf("This is a Golang REST API! 🐿")
        return c.SendString(msg)
    })


		
		app.Get("/json", func(c *fiber.Ctx) error {
			json := map[string]string{
				"hello": "world",
				"welcome": "friends",
			}
			return c.JSON(json)
		})
		
		// c *fiber.Ctx will add Params, Query, etc (context) to var c
		// app.Get("/:name" ... ... c.Params("name")
		app.Get("/:name", func(c *fiber.Ctx) error {
			msg := fmt.Sprintf("Hello, %s 👋!", c.Params("name"))
			return c.SendString(msg)
    })

		// todo - Add POST
		
		app.Get("/api/*", func(c *fiber.Ctx) error {
				msg := fmt.Sprintf("✋ %s", c.Params("*"))
				return c.SendString(msg)
		})

    log.Fatal(app.Listen(":3000"))

}