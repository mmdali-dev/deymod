package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mmdali-dev/deymod/db/model"
	"github.com/mmdali-dev/deymod/db/service"
)

func AddComment(c *fiber.Ctx) error {
	data := new(commentdata)

	if err := c.BodyParser(data); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	if data.Score > 5 {
		return c.Status(fiber.StatusBadRequest).SendString("invalid score")
	}
	user := c.Locals("user").(model.User)
	switch data.UserType {
	case "man":
		com := model.ManComment{
			UserID:  user.ID,
			ManID:   data.TargetID,
			Score:   data.Score,
			Comment: data.Text,
		}
		err := service.Comment.CreateManComment(&com)
		if err != nil || com.ID == 0 {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}
		return c.SendStatus(fiber.StatusOK)
	case "woman":
		com := model.WomanComment{
			UserID:  user.ID,
			WomanID: data.TargetID,
			Score:   data.Score,
			Comment: data.Text,
		}
		err := service.Comment.CreateWomanComment(&com)
		if err != nil || com.ID == 0 {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}
		return c.SendStatus(fiber.StatusOK)
	case "picturor":
		com := model.PicturorComment{
			UserID:     user.ID,
			PicturorID: data.TargetID,
			Score:      data.Score,
			Comment:    data.Text,
		}
		err := service.Comment.CreatePicturorComment(&com)
		if err != nil || com.ID == 0 {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}
		return c.SendStatus(fiber.StatusOK)
	case "location":
		com := model.LocationComment{
			UserID:     user.ID,
			LocationID: data.TargetID,
			Score:      data.Score,
			Comment:    data.Text,
		}
		err := service.Comment.CreateLocationComment(&com)
		if err != nil || com.ID == 0 {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}
		return c.SendStatus(fiber.StatusOK)
	default:
		return c.Status(fiber.StatusBadRequest).SendString("user type is wrong")
	}

}
