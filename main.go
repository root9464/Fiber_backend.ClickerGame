package main

import (
	"log"
	database "root/src/database/controller"
	"root/src/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func core() {
	_, err := database.ConnectToDB()
	if err != nil {
		log.Fatal(err)
	}

	app := fiber.New()
	routes.AllRoutes(app)
	app.Use(cors.New())
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404)
	})

	log.Fatal(app.Listen(":3000"))

}

func main() {
	core()
}
