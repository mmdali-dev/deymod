package server

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func Run(ipport string) {
	app := fiber.New(AppConfig())
	AddRoutes(app)
	log.Fatal(app.Listen(ipport))
}
