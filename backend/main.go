package main

import (
	"buzzer/db"
	"log"
	"github.com/gofiber/fiber/v3"
)

type LogMsg struct {
	msg string
	log_type string
}

func (l LogMsg) Log() {
	switch l.log_type {
	case "INFO":
		log.Println("\033[34m" + "INFO: " + l.msg + "\033[0m")
	case "ERROR":
		log.Println("\033[31m" + "ERRROR:" + l.msg + "\033[0m")
	case "DEBUG":
		log.Println("\033[35m" + "DEBUG: " + l.msg + "\033[0m")
	case "WARNING":
		log.Println("\033[33m" + "WARNING: " + l.msg + "\033[0m")
	default:
		log.Println(l.msg)
	}
}

func main() {
	LogMsg{"Starting server", "INFO"}.Log()
	app := fiber.New()

	
	app.Static("/", "../frontend/dist")
	
	db.InitDB()
	LogMsg{"Database initialized", "WARNING"}.Log()
	


	app.Get("/", func(c fiber.Ctx) error {
        return c.SendFile("ndex.html")
    })

	app.Listen(":8080")
}
