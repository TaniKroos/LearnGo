package main

import (
	"log"
	"os"

	"github.com/TaniKroos/url-shortner/routes"
	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/joho/godotenv"
)

func setUpRoutes(app *fiber.App) {
	app.Get("/:url", routes.ResolveUrl)
	app.Post("/api/v1", routes.ShortenUrl)

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
