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
	app.Use(cors.New(
		cors.Config{
			AllowOrigins:     "http://localhost:5173, http://0.0.0.0:5173, https://4f67-95-105-125-55.ngrok-free.app",
			AllowCredentials: true,
		},
	))

	routes.AllRoutes(app)

	log.Fatal(app.Listen(":3000"))

}

func main() {
	core()
}
