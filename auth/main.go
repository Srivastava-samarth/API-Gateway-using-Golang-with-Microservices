package main

import (
	"api/handlers"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New();
	app.Post("/register",handlers.Register);
	app.Post("/login",handlers.Login);

	app.Listen(":6001");
}