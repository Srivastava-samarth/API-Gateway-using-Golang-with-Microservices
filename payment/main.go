package main

import (
	"api/handlers"

	"github.com/gofiber/fiber/v2"
)

func main() {
    app := fiber.New()

    app.Post("/payment", handlers.ProcessPayment)

    app.Listen(":6003")
}
