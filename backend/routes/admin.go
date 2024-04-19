package routes

import (
	"buzzer/auth"
	"buzzer/database"
	"buzzer/observer"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func SetupAdminRoutes(app *fiber.App) {
	app.Route("/admin", func(admin fiber.Router) {
		admin.Get("/dashboard", func(c *fiber.Ctx) error {
			if auth.IsLogged(c) {
				return c.SendFile("admin/dashboard/index.html")
			} else {
				return c.Redirect("/login")
			}
		})

		admin.Get("/teams", func(c *fiber.Ctx) error {
			teams, err := database.GetTeams()
			if err != nil {
				return c.SendStatus(502)
			}

			return c.JSON(teams)

		})

		admin.Get("/isAdmin", func(c *fiber.Ctx) error {
			if !auth.IsLogged(c) {
				return c.JSON(fiber.Map{"isAdmin": false})
			} else {
				return c.JSON(fiber.Map{"isAdmin": true})
			}
		})

		admin.Get("/ws", websocket.New(func(c *websocket.Conn) {
			observerWso := &observer.WebsocketObserver{Conn: c}
			subject := observer.SubjectInstance

			subject.Attach(observerWso)

			for {
				mtype, msg, err := c.ReadMessage()
				if err != nil {
					log.Printf("Error reading message: %s", err)
					subject.Detach(observerWso)
					break
				}
				log.Printf("Received message: %s", msg)

				err = c.WriteMessage(mtype, msg)
				if err != nil {
					log.Printf("Error sending message: %s", err)
					subject.Detach(observerWso)
					break
				}
			}
		}))

		admin.Get("/lock-state", func(c *fiber.Ctx) error {
			state, err := database.LockState()
			if err != nil {
				return c.SendStatus(502)
			}

			return c.JSON(fiber.Map{"locked": state})
		})

		admin.Get("/buzzer-number", func(c *fiber.Ctx) error {
			number, err := database.GetLockedTeams()
			if err != nil {
				return c.SendStatus(502)
			}
			return c.JSON(fiber.Map{"number": len(number)})
		})
		
		admin.Post("/teams/:id/increment", func(c *fiber.Ctx) error {
			if !auth.IsLogged(c) {
				return c.Redirect("/login")
			} else {
				id, err := strconv.Atoi(c.Params("id"))
				if err != nil {
					log.Println(err)
					return c.SendStatus(400)
				}

				team, err := database.GetTeamID(id)
				if err != nil {
					return c.SendStatus(502)
				}

				err = database.UpdateScore(team, 1)
				if err != nil {
					return c.SendStatus(502)
				}

				return c.SendStatus(200)
			}
		})

		admin.Post("/lock", func(c *fiber.Ctx) error {
			if (!auth.IsLogged(c)) {
				return c.Redirect("/login")
			}
			err := database.LockAllTeams()
			if err != nil {
				return c.SendString("Error locking teams, please retry")

			}
			return c.SendStatus(200)
		})

		admin.Post("/unlock", func(c *fiber.Ctx) error {
			if (!auth.IsLogged(c)) {
				return c.Redirect("/login")
			}
			err := database.UnlockAllTeams()
			if err != nil {
				return c.SendString("Error unlocking teams, please retry")

			}
			return c.SendStatus(200)
		})
	})
}
