package handlers

import "github.com/gofiber/fiber/v2"

func ProcessPayment(c *fiber.Ctx) error {
    return c.SendString("Process payment endpoint")
}
