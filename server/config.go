package server

import (
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func AppConfig() fiber.Config {
	engine := html.New("./views", ".html")
	return fiber.Config{
		Prefork:      false,
		ServerHeader: "Mohammad Ali",
		AppName:      "Deymod",
		JSONEncoder:  json.Marshal,
		JSONDecoder:  json.Unmarshal,
		Views:        engine,
	}
}
