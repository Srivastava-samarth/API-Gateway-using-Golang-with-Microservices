package main

import (
	"api/handlers"
	"api/middleware"

	"github.com/gofiber/fiber/v2"
)

func main() {
    app := fiber.New()
    api := app.Group("/api", middleware.Protected())
    api.Get("/payments", handlers.GetPayments)
    api.Post("/payments", handlers.CreatePayment)
    api.Put("/payments/:id", handlers.UpdatePayment)
    api.Delete("/payments/:id", handlers.DeletePayment)

    app.Listen(":6003")
}
