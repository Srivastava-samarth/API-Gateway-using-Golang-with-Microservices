package routes

import (
	"bytes"
	"io"
	"net/http"
	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(app *fiber.App) {
	app.Post("/auth/register", func(c *fiber.Ctx) error {
		res, err := http.Post("http://localhost:6001/register", "application/json", bytes.NewReader(c.Body()))
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

	app.Post("/auth/login", func(c *fiber.Ctx) error {
		res, err := http.Post("http://localhost:6001/login", "application/json", bytes.NewReader(c.Body()))
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
