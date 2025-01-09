package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mmdali-dev/deymod/server/handler"
	"github.com/mmdali-dev/deymod/server/middleware"
)

func AddRoutes(app *fiber.App) {
	app.Post("/login", handler.Login)
	app.Post("/register/modelorpicturor", middleware.CheckBooker, handler.RegisterModelOrPicturor)
	app.Post("/register/location", middleware.CheckBooker, handler.RegisterLocation)
	app.Post("/register/booker", middleware.CheckAdmin, handler.RegisterBooker)
	app.Post("/register/user", handler.RegisterUser)

	app.Post("/comment", middleware.CheckUser, handler.AddComment)

	app.Get("/", handler.FullTree)
	app.Get("/test", handler.TestingData)
}
