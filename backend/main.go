package main

import (
	"buzzer/db"
	"github.com/gofiber/fiber/v3"
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
