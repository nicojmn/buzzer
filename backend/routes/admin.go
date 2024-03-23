package routes

import (
	"github.com/gofiber/fiber/v2"
)

func SetupAdminRoutes(app *fiber.App) {
	app.Route("/admin", func (admin fiber.Router) {
		admin.Get("/", func(c *fiber.Ctx) error {
			return fiber.ErrUnauthorized;
		})

	})
}

