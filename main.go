package main

import (
	"api/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New();

	routes.AuthRoutes(app);
	routes.OrdersRoutes(app);
	routes.PaymentRoutes(app);

	app.Listen(":6000");
}