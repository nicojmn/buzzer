package routes

import (
	"buzzer/auth"
	"buzzer/database"
	"buzzer/observer"
	"buzzer/config"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func SetupBuzzerRoutes(app *fiber.App) {
	app.Route("/buzzer", func(buzzer fiber.Router) {
		buzzer.Get("/ws", websocket.New(func(c *websocket.Conn) {
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

		buzzer.Get("/create", func(c *fiber.Ctx) error {
			return c.SendFile("buzzer/create-team/index.html")
		})

		buzzer.Post("/create", func(c *fiber.Ctx) error {
			log.Println("Creating team")
			teamName := c.FormValue("teamName")
			if (len(teamName) > 24) {
				return c.SendString("Team name too long, please retry")
			}

			conf, err := config.LoadConfig("config.yaml")
			if err != nil {
				return c.SendString("Error adding team, please retry")
			}

			teamsNumber, err := database.GetTeams()
			if err != nil {
				return c.SendString("Error adding team, please retry")
			}

			if (len(teamsNumber) >= conf.MaxTeams) {
				return c.SendString("Error adding team, max teams reached")
			}

			err = database.AddTeam(teamName)
			if err != nil {
				return c.SendString("Error adding team, please retry")
			}

			token, err := auth.Create_Team_JWT_Token(teamName)
			if err != nil {
				return c.SendString("Error adding team, please retry")
			}

			c.Cookie(&fiber.Cookie{
				Name:  "jwt",
				Value: token,
				HTTPOnly: true,
				Secure: false,
			})

			return c.Redirect("/")
		})

		buzzer.Post("/press", func(c *fiber.Ctx) error {
			if (!auth.IsTeam(c)) {
				return c.Redirect("/buzzer/create-team")
			}

			teamID, err := strconv.Atoi(c.FormValue("team_id"))
			if err != nil {
				return c.SendStatus(400)
			}

			team, err := database.GetTeamID(teamID)
			if err != nil {
				return c.SendStatus(502)
			}

			database.UpdatePressedAt(team)
			return c.SendStatus(200)
		})
	})
}
