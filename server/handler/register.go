package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mmdali-dev/deymod/db/model"
	"github.com/mmdali-dev/deymod/db/service"
)

func RegisterUser(c *fiber.Ctx) error {
	data := new(singleform)
	if err := c.BodyParser(data); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	token, err := service.User.RegisterUser(data.Username, data.Password)
	if err != nil || token == "" {
		return c.Status(fiber.StatusBadRequest).SendString("username already exist")
	}
	return c.JSON(fiber.Map{
		"token": token,
		"role":  "user",
	})

}

func RegisterModelOrPicturor(c *fiber.Ctx) error {
	data := new(authform)
	var token string
	var err error
	if err := c.BodyParser(data); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	booker := c.Locals("booker").(model.Booker)
	switch data.UserType {
	case "model":
		token, err = service.Model.RegisterModel(data.Username, data.Password, booker.ID)
	case "picturor":
		token, err = service.Picturor.RegisterPicturor(data.Username, data.Password, booker.ID)

	default:
		return c.Status(fiber.StatusBadRequest).SendString("register type is wrong")
	}
	if err != nil || token == "" {
		return c.Status(fiber.StatusNotFound).SendString("username is already use")
	}
	return c.JSON(fiber.Map{
		"token": token,
		"role":  data.UserType,
	})

}

func RegisterBooker(c *fiber.Ctx) error {
	data := new(singleform)
	var token string
	var err error
	if err := c.BodyParser(data); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	token, err = service.Booker.RegisterBooker(data.Username, data.Password)
	if err != nil || token == "" {
		return c.Status(fiber.StatusNotFound).SendString("username already in use")
	}
	return c.JSON(fiber.Map{
		"token": token,
		"role":  "booker",
	})
}

func RegisterLocation(c *fiber.Ctx) error {
	data := new(locationform)
	if err := c.BodyParser(data); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	booker := c.Locals("booker").(model.Booker)
	loc := model.Location{NameCurator: data.NameCurator,
		FamilyCurator:  data.FamilyCurator,
		PhoneCurator:   data.PhoneCurator,
		AddressCurator: data.AddressCurator,
		DecorNumber:    data.DecorNumber,
		WorkType:       data.WorkType,
		City:           data.City,
		BookerID:       booker.ID}
	err := service.Location.CreateLocation(&loc)
	if err != nil || loc.ID == 0 {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	return c.SendStatus(200)
}
