package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mmdali-dev/deymod/db/model"
	"github.com/mmdali-dev/deymod/db/service"
)

func FullTree(c *fiber.Ctx) error {
	mans := []model.PublicManModel{}
	womans := []model.PublicWomanModel{}
	picturors := []model.PublicPicturor{}
	locs := []model.PublicLocation{}
	service.Model.FullTreeMan(&mans)
	service.Model.FullTreeWoman(&womans)
	service.Picturor.FullTree(&picturors)
	service.Location.FullTree(&locs)

	return c.JSON(fiber.Map{
		"mans":      mans,
		"womans":    womans,
		"picturors": picturors,
		"locs":      locs,
	})
}

func TestingData(c *fiber.Ctx) error {

	return c.SendString("ok")
}
