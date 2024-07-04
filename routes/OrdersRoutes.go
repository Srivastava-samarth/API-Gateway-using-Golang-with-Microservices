package routes

import (
	"bytes"
	"io"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func OrdersRoutes(app *fiber.App) {
	app.Get("/orders", func(c *fiber.Ctx) error {
		res, err := http.Get("http://localhost:6002/orders")
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		defer res.Body.Close()

		body, err := io.ReadAll(res.Body)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		return c.Status(res.StatusCode).SendString(string(body))
	})

	app.Post("/orders", func(c *fiber.Ctx) error {
		res, err := http.Post("http://localhost:6002/orders", "application/json", bytes.NewReader(c.Body()))
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		defer res.Body.Close()

		body, err := io.ReadAll(res.Body)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		return c.Status(res.StatusCode).SendString(string(body))
	})
}