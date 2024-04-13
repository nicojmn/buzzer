package routes

import (
	"buzzer/auth"
	"buzzer/database"

	"github.com/gofiber/fiber/v2"
)

func SetupAdminRoutes(app *fiber.App) {
	app.Route("/admin", func (admin fiber.Router) {
		admin.Get("/dashboard", func(c *fiber.Ctx) error {
			if (auth.IsLogged(c)) {
				return c.SendFile("admin/dashboard/index.html")
			} else {
				return c.Redirect("/login")
			}
		})

		admin.Get("/teams", func(c *fiber.Ctx) error {
			if (!auth.IsLogged(c)) {
				return c.Redirect("/login")
			} else {
				teams, err := database.GetTeams()
				if err != nil {
					return c.SendStatus(502)
				}

				return c.JSON(teams)
			}
		})

		admin.Get("/isAdmin" , func(c *fiber.Ctx) error {
			if (!auth.IsLogged(c)) {
				return c.JSON(fiber.Map{"isAdmin": false})
			} else {
				return c.JSON(fiber.Map{"isAdmin": true})
			}
		})

	})
}

