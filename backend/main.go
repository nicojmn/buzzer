package main

import (
	"buzzer/database"
	"buzzer/routes"
	"fmt"
	"log"
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
	
	db, err := database.InitDB("db.sqlite")
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Printf("Database initialized: %s", db.Name())

	database.Clear(db)

	for i := 0; i < 20; i++ {
		err = database.AddTeam(db, fmt.Sprintf("Team %d", i))
		if err != nil {
			log.Println(err)
		}
	}

	database.AddAdmin(db, "admin", "admin")
	database.AddAdmin(db, "nico", "nico")


	app.Get("/", func(c *fiber.Ctx) error {
        return c.SendFile("index.html")
    })

	routes.SetupAdminRoutes(app)
	routes.SetupAuthRoutes(app, db)

	app.Listen(":8080")

}
