package main

import (
	"log"
	"buzzer/database"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	
	app.Static("/", "../frontend/dist")
	
	db.InitDB()

	app.Get("/", func(c fiber.Ctx) error {
        return c.SendFile("index.html")
    })

	app.Listen(":8080")

}
