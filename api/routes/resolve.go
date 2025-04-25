package routes

import (
	"github.com/Kosench/go-url-shortener-redis/api/database"
	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
	"log"
)

func ResolveURL(c *fiber.Ctx) error {
	url := c.Params("url")
	if url == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "URL parameter is required",
		})
	}

	r := database.CreateClient(0)
	defer r.Close()

	value, err := r.Get(database.Ctx, url).Result()
	if err == redis.Nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "short URL not found in the database",
		})
	} else if err != nil {
		log.Printf("Redis error while fetching URL: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to fetch URL from the database",
		})
	}

	rInr := database.CreateClient(1)
	defer rInr.Close()

	if _, err := rInr.Incr(database.Ctx, "counter").Result(); err != nil {
		log.Printf("Redis error while incrementing counter: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to increment counter",
		})
	}

	return c.Redirect(value, 301)
}
