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

		return c.JSON(fiber.Map{"token": token})

	})
	app.Get("/protected", func(c *fiber.Ctx) error {
		token := c.Get("Authorization")
		if token == "" {
			return c.SendString(fiber.ErrUnauthorized.Message)
		}

		return c.SendStatus(500)
	})
}