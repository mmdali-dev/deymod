package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mmdali-dev/deymod/db/service"
)

func Login(c *fiber.Ctx) error {
	data := new(authform)
	var token string
	var err error
	if err := c.BodyParser(data); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	switch data.UserType {
	case "user":
		token, err = service.User.LoginUser(data.Username, data.Password)

	case "model":
		token, err = service.Model.LoginModel(data.Username, data.Password)
	case "picturor":
		token, err = service.Picturor.LoginPicturor(data.Username, data.Password)
	case "booker":
		token, err = service.Booker.LoginBooker(data.Username, data.Password)
	case "admin":
		token, err = service.Admin.LoginAdmin(data.Username, data.Password)
	default:
		return c.Status(fiber.StatusBadRequest).SendString("user type is wrong")
	}
	if err != nil || token == "" {
		return c.Status(fiber.StatusNotFound).SendString("username or password wrong")
	}
	return c.JSON(fiber.Map{
		"token": token,
		"role":  data.UserType,
	})
}
