package main

import (
	"log"

	"fiber.osaxon/database"
	"fiber.osaxon/router"
	"github.com/gofiber/fiber/v2"
	// "github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New(fiber.Config{
		Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Fiber",
		AppName:       "GoFiber",
	})
	// app.Use(cors.New())
	database.Connect()
	router.SetupRoutes(app)

	app.Get("/test/:name", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"name": c.Params("name")})
	})

	log.Fatal(app.Listen(":3000"))
}
