package routes

import (
	"github.com/TaniKroos/url-shortner/database"
	"github.com/gofiber/fiber/v3"
	"github.com/redis/go-redis/v9"
)

func ResolveUrl(c *fiber.Ctx) error {
	url := c.Params("url")
	r := database.CreateClient(0)
	defer r.close()
	r.Get(database.Ctx, url)
	value, err := r.Get(database.Ctx, url).Result()
	if err == redis.Nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "short not found"})

	} else if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Cannot connect to db"})

	}
	rInr := database.CreateClient(1)
	defer rInr.Close()

	_ = rInr.rInr(database.Ctx, "counter")
	return c.Redirect(value, 301)

}
