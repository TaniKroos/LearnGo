package main

import (
	"log"
	"os"

	"github.com/TaniKroos/url-shortner/routes"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/joho/godotenv"
)

func setUpRoutes(app *fiber.App) {
	app.Get("/:url", routes.ResolveURL)
	app.Post("/api/v1", routes.ShortenURL)

}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	app := fiber.New()

	app.Use(logger.New())
	setUpRoutes(app)

	log.Fatal(app.Listen(os.Getenv("APP_PORT")))
}
