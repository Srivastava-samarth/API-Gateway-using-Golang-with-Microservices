package main

import (
	"api/handlers"
	"api/middleware"
	"github.com/gofiber/fiber/v2"
)

func main() {
    app := fiber.New()

    api := app.Group("/api", middleware.Protected())
    api.Get("/orders", handlers.GetOrders)
    api.Post("/orders", handlers.CreateOrder)
    api.Put("/orders/:id", handlers.UpdateOrder)
    api.Delete("/orders/:id", handlers.DeleteOrder)
    

    app.Listen(":6002")
}