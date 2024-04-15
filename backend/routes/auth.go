package routes

import (
	"buzzer/auth"
	"time"

	"github.com/gofiber/fiber/v2"
)

func SetupAuthRoutes(app *fiber.App) {

	app.Get("/login", func(c *fiber.Ctx) error {
		return c.SendFile("login.html")
	})

	app.Post("/login", func(c *fiber.Ctx) error {
		username := c.FormValue("username")
		password := c.FormValue("password")

		token, err := auth.Create_JWT_Token(username, password)
		if err != nil {
			return c.SendString("Authentification failed, please retry")
		}

		c.Cookie(&fiber.Cookie{
			Name:  "jwt",
			Value: token,
			HTTPOnly: true,
			Secure: true,
		})

		return c.Redirect("/admin/dashboard")
	})

	app.Get("/logout", func(c *fiber.Ctx) error {
		c.Cookie(&fiber.Cookie{
			Name: "jwt",
			Value: "",
			Expires: time.Now().Add(-time.Hour),
		})

		return c.Redirect("/login")
	})
}