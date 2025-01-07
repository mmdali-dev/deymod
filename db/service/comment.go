package service

import (
	"github.com/mmdali-dev/deymod/db"
	"github.com/mmdali-dev/deymod/db/model"
)

type commentService service

var (
	Comment commentService = commentService{"modelService", "CRUD and auth for Models", db.GetDB()}
)

func (s *commentService) CreateManComment(com *model.ManComment) (err error) {
	err = s.DB.Create(&com).Error
	return err
}

func (s *commentService) CreateWomanComment(com *model.WomanComment) (err error) {
	err = s.DB.Create(&com).Error
	return err
}

func (s *commentService) CreateLocationComment(com *model.LocationComment) (err error) {
	err = s.DB.Create(&com).Error
	return err
}

func (s *commentService) CreatePicturorComment(com *model.PicturorComment) (err error) {
	err = s.DB.Create(&com).Error
	return err
}
