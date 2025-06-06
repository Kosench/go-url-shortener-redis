package main

import (
	"fmt"
	"github.com/Kosench/go-url-shortener-redis/api/routes"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func setuoRoutes(app *fiber.App) {
	app.Get("/:url", routes.ResolveURL)
	app.Post("/api/v1", routes.ShortenURL)
}

func main() {

	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}
	app := fiber.New()

	app.Use(logger.New())

	setuoRoutes(app)

	log.Fatal(app.Listen(os.Getenv("APP_PORT")))
}
