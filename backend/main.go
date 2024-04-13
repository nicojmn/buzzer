package main

import (
	"buzzer/auth"
	"buzzer/database"
	"buzzer/routes"
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
		return
	}

	app := fiber.New()

	
	app.Static("/", "../frontend/build")

	// only for testing
	DebugPlayground()


	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendFile("index.html")
    })
	
	app.Get("/test", func(c *fiber.Ctx) error {
		if (auth.IsLogged(c)) {
			return c.SendString("You are logged")
		} else {
			return c.SendString("You are not logged")
		}
	})

	routes.SetupAdminRoutes(app)
	routes.SetupAuthRoutes(app)

	app.Listen(":8080")

}

func DebugPlayground() {

	database.Clear()

	for i := 0; i < 10; i++ {
		err := database.AddTeam(fmt.Sprintf("Team %d", i))
		if err != nil {
			log.Println(err)
		}
	}

	database.AddAdmin("admin", "admin")
	database.AddAdmin("nico", "nico")

	teams, err := database.GetTeams()
	if err != nil {
		log.Println(err)
	}
	for _, team := range teams {
		database.UpdatePressedAt(team)
		time.Sleep(1 * time.Second)
	}
}
