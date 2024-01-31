package main

import (
	"log"

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

	router.SetupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
