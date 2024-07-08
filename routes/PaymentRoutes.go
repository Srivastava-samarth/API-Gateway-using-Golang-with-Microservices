package routes

import (
	"bytes"
	"io"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func PaymentRoutes(app *fiber.App) {
	app.Get("/payments",func(c *fiber.Ctx) error {
		res,err := http.Get("http://localhost:6003/api/payments")
		if err!=nil{
			return c.Status(500).SendString(err.Error())
		}
		defer res.Body.Close();

		body,err := io.ReadAll(res.Body);
		if err!=nil{
			return c.Status(500).SendString(err.Error());
		}
		return c.Status(res.StatusCode).SendString(string(body));
	})


	app.Post("/payments", func(c *fiber.Ctx) error {
		res, err := http.Post("http://localhost:6003/api/payments", "application/json", bytes.NewReader(c.Body()))
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

	app.Put("/payment/:id",func(c *fiber.Ctx) error {
		req,err := http.NewRequest(http.MethodPut,"http://localhost:6003/api/payments/" + c.Params("id"),bytes.NewReader(c.Body()))
		if err!=nil{
			return c.Status(500).SendString(err.Error());
		}
		req.Header.Set("Content-Type","application/json")
		client := &http.Client{}
		res,err := client.Do(req);
		if err!=nil{
			return c.Status(500).SendString(err.Error());
		}
		defer res.Body.Close();

		body,err := io.ReadAll(res.Body);
		if err!=nil{
			return c.Status(500).SendString(err.Error())
		}
		return c.Status(res.StatusCode).SendString(string(body))
	})

	app.Delete("/payments/:id",func(c *fiber.Ctx) error {
		req,err := http.NewRequest(http.MethodPut,"http://localhost:6003/api/payments/" + c.Params("id"),nil)
		if err!=nil{
			return c.Status(500).SendString(err.Error());
		}
		client := &http.Client{}
		res,err := client.Do(req);
		if err!=nil{
			return c.Status(500).SendString(err.Error());
		}
		defer res.Body.Close();
		body,err := io.ReadAll(res.Body);
		if err!=nil{
			return c.Status(500).SendString(err.Error())
		}
		return c.Status(res.StatusCode).SendString(string(body))
	})
}
