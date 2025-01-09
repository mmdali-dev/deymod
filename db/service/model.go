package service

import (
	"github.com/mmdali-dev/deymod/db"
	"github.com/mmdali-dev/deymod/db/model"
	"github.com/mmdali-dev/deymod/internal"
)

type modelService service

var (
	Model modelService = modelService{"modelService", "CRUD and auth for Models", db.GetDB()}
)

func (s *modelService) CreateModel(modell *model.Model) (err error) {
	err = s.DB.Create(&modell).Error
	return err
}

func (s *modelService) DeleteModel(modell *model.Model) (err error) {
	err = s.DB.Where(modell).Delete(modell).Error
	return err
}

func (s *modelService) UpdateModel(modell *model.Model) (err error) {
	err = s.DB.Save(modell).Error
	return err
}

func (s *modelService) FindModel(modell *model.Model) (err error) {
	err = s.DB.Where(modell).First(modell).Error
	return err
}

func (s *modelService) FindModels(modell *model.Model, modells *[]model.Model) (err error) {
	err = s.DB.Where(modell).Find(modells).Error
	return err
}

func (s *modelService) AllModels(modells *[]model.Model) (err error) {
	err = s.DB.Find(&modells).Error
	return err
}

func (s *modelService) ChangeToken(modell *model.Model) error {
	var (
		token    string
		haveuser model.Model
		err      error
	)
	ok := false
	for !ok {
		token = internal.GenToken()
		haveuser = model.Model{Token: token}
		err = s.FindModel(&haveuser)
		if err != nil || haveuser.ID == 0 {
			ok = true
		}
	}

	modell.Token = token
	err = s.UpdateModel(modell)
	if err != nil || modell.Token != token {
		return err
	}
	return nil

}

func (s *modelService) RegisterModel(username, password string, bookerid uint) (token string, err error) {
	modell := model.Model{Username: username, Password: password, BookerID: bookerid}
	err = s.CreateModel(&modell)
	if err != nil || modell.ID == 0 {
		token = ""
		return
	}
	s.ChangeToken(&modell)
	token = modell.Token
	return
}

func (s *modelService) LoginModel(username, password string) (token string, err error) {
	modell := model.Model{Username: username, Password: password}
	err = s.FindModel(&modell)
	if err != nil || modell.ID == 0 {
		token = ""
		return
	}
	s.ChangeToken(&modell)
	token = modell.Token
	return
}

func (s *modelService) GetModelByToken(token string) (modell model.Model, err error) {
	modell = model.Model{Token: token}
	err = s.FindModel(&modell)
	return
}

// public information for Mans service
func (s *modelService) CreatePublicMan(pubman *model.PublicManModel) (err error) {
	err = s.DB.Create(&pubman).Error
	return
}

func (s *modelService) UpdatePublicMan(pubman *model.PublicManModel) (err error) {
	err = s.DB.Save(pubman).Error
	return err
}
func (s *modelService) FindPublicMan(pubman *model.PublicManModel) (err error) {
	err = s.DB.Where(pubman).First(pubman).Error
	return err
}

func (s *modelService) FindPublicMans(pubman *model.PublicManModel, pubmans *[]model.PublicManModel) (err error) {
	err = s.DB.Where(pubman).Find(pubmans).Error
	return err
}
func (s *modelService) AllPublicMans(pubmans *[]model.PublicManModel) (err error) {
	err = s.DB.Find(&pubmans).Error
	return err
}
func (s *modelService) DeletePublicMan(pubman *model.PublicManModel) (err error) {
	err = s.DB.Where(pubman).Delete(pubman).Error
	return err
}

// womans .................................
func (s *modelService) CreatePublicWoman(pubwoman *model.PublicWomanModel) (err error) {
	err = s.DB.Create(&pubwoman).Error
	return
}

func (s *modelService) UpdatePublicWoman(pubwoman *model.PublicWomanModel) (err error) {
	err = s.DB.Save(pubwoman).Error
	return err
}
func (s *modelService) FindPublicWoman(pubwoman *model.PublicWomanModel) (err error) {
	err = s.DB.Where(pubwoman).First(pubwoman).Error
	return err
}

func (s *modelService) FindPublicWomans(pubwoman *model.PublicWomanModel, pubwomans *[]model.PublicWomanModel) (err error) {
	err = s.DB.Where(pubwoman).Find(pubwomans).Error
	return err
}
func (s *modelService) AllPublicWomans(pubwomans *[]model.PublicWomanModel) (err error) {
	err = s.DB.Find(&pubwomans).Error
	return err
}
func (s *modelService) DeletePublicWoman(pubwoman *model.PublicWomanModel) (err error) {
	err = s.DB.Where(pubwoman).Delete(pubwoman).Error
	return err
}

func (s *modelService) FullTreeMan(models *[]model.PublicManModel) (err error) {
	err = s.DB.Preload("Images").Preload("Comments").Find(models).Error
	return err
}

func (s *modelService) FullTreeWoman(models *[]model.PublicWomanModel) (err error) {
	err = s.DB.Preload("Images").Preload("Comments.User").Find(models).Error
	return err
}
