package routes

import (
	"buzzer/auth"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupAuthRoutes(app *fiber.App, db *gorm.DB) {

	app.Get("/login", func(c *fiber.Ctx) error {
		return c.SendFile("login.html")
	})

	app.Post("/login", func(c *fiber.Ctx) error {
		username := c.FormValue("username")
		password := c.FormValue("password")

		token, err := auth.Create_JWT_Token(db, username, password)
		if err != nil {
			return c.SendString("Authentification failed, please retry")
		}

		c.Cookie(&fiber.Cookie{
			Name:  "jwt",
			Value: token,
			HTTPOnly: true,
			Secure: true,
		})

		return c.Redirect("/admin")
	})
}