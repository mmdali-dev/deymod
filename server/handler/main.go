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
	service.Booker.RegisterBooker("javad", "1234")
	service.Model.RegisterModel("maryam", "1234", 1)
	pubwom := model.PublicWomanModel{ModelID: 1, Waist: 343}
	service.Model.CreatePublicWoman(&pubwom)
	manimage := model.ImageWomanModel{FileName: "main.png", PublicWomanID: 1}
	service.Model.AddWomanImage(&manimage)
	service.User.RegisterUser("user1", "9876")
	com := model.WomanComment{Score: 3, Comment: "salam maryam", UserID: 1, WomanID: 1}
	service.Comment.CreateWomanComment(&com)
	return c.SendString("ok")
}
