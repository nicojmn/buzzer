package routes

import (
	"buzzer/auth"
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

	})
}

