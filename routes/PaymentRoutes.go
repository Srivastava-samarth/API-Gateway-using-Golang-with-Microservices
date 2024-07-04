package routes

import (
	"bytes"
	"io"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func PaymentRoutes(app *fiber.App) {
	app.Post("/payment", func(c *fiber.Ctx) error {
		res, err := http.Post("http://localhost:6003/payment", "application/json", bytes.NewReader(c.Body()))
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
