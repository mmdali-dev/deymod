package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mmdali-dev/deymod/db/model"
	"github.com/mmdali-dev/deymod/db/service"
)

func CheckBooker(c *fiber.Ctx) error {
	token := c.Get("token")
	booker := model.Booker{Token: token}
	err := service.Booker.FindBooker(&booker)
	if err != nil || booker.ID == 0 {
		return c.Status(fiber.StatusUnauthorized).SendString("invalid token for booker")
	}
	c.Locals("booker", booker)
	return c.Next()
}

func CheckAdmin(c *fiber.Ctx) error {
	token := c.Get("token")
	admin := model.Admin{Token: token}
	err := service.Admin.FindAdmin(&admin)
	if err != nil || admin.ID == 0 {
		return c.Status(fiber.StatusUnauthorized).SendString("invalid token for admin")
	}
	c.Locals("admin", admin)
	return c.Next()
}

func CheckUser(c *fiber.Ctx) error {
	token := c.Get("token")
	user := model.User{Token: token}
	err := service.User.FindUser(&user)
	if err != nil || user.ID == 0 {
		return c.Status(fiber.StatusUnauthorized).SendString("invalid token for user")
	}
	c.Locals("user", user)
	return c.Next()
}
