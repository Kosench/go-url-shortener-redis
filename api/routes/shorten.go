package routes

import (
	"github.com/Kosench/go-url-shortener-redis/api/helpers"
	"github.com/gofiber/fiber/v3"
	"time"
)

type request struct {
	URL         string        `json:"url"`
	CustomShort string        `json:"short"`
	Expire      time.Duration `json:"expire"`
}

type response struct {
	URL                string        `json:"url"`
	CustomShort        string        `json:"short"`
	Expire             time.Duration `json:"expire"`
	RateLimitRemaining int           `json:"rate_limit"`
	RateLimitReset     time.Duration `json:"rate_limit_reset"`
}

func ShortenURL(c fiber.Ctx) error {
	body := new(request)

	if err := c.Bind().Body(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse JSON"})
	}

	//TODO implement rate limiting

	//TODO check if the input if an actual URL
	if !govalidator.IsURL(body.URL) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid URL"})
	}

	//TODO check for domain error
	if !helpers.RemoveDomainError(body.URL) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid URL"})
	}

	//TODO enforse htttps, SSL
	body.URL = helpers.EnforceHTTP(body.URL)
}
