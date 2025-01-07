package service

import (
	"github.com/mmdali-dev/deymod/db"
	"github.com/mmdali-dev/deymod/db/model"
	"github.com/mmdali-dev/deymod/internal"
)

type picturorService service

var (
	Picturor picturorService = picturorService{"picturorService", "CRUD and auth for Picturors", db.GetDB()}
)

func (s *picturorService) CreatePicturor(picturor *model.Picturor) (err error) {
	err = s.DB.Create(&picturor).Error
	return err
}

func (s *picturorService) DeletePicturor(picturor *model.Picturor) (err error) {
	err = s.DB.Where(picturor).Delete(picturor).Error
	return err
}

func (s *picturorService) UpdatePicturor(picturor *model.Picturor) (err error) {
	err = s.DB.Save(picturor).Error
	return err
}

func (s *picturorService) FindPicturor(picturor *model.Picturor) (err error) {
	err = s.DB.Where(picturor).First(picturor).Error
	return err
}

func (s *picturorService) FindPicturors(picturor *model.Picturor, picturors *[]model.Picturor) (err error) {
	err = s.DB.Where(picturor).Find(picturors).Error
	return err
}

func (s *picturorService) AllPicturors(picturors *[]model.Picturor) (err error) {
	err = s.DB.Find(&picturors).Error
	return err
}

func (s *picturorService) ChangeToken(picturor *model.Picturor) error {
	var (
		token    string
		haveuser model.Picturor
		err      error
	)
	ok := false
	for !ok {
		token = internal.GenToken()
		haveuser = model.Picturor{Token: token}
		err = s.FindPicturor(&haveuser)
		if err != nil || haveuser.ID == 0 {
			ok = true
		}
	}

	picturor.Token = token
	err = s.UpdatePicturor(picturor)
	if err != nil || picturor.Token != token {
		return err
	}
	return nil

}

func (s *picturorService) RegisterPicturor(username, password string, bookerid uint) (token string, err error) {
	picturor := model.Picturor{Username: username, Password: password, BookerID: bookerid}
	err = s.CreatePicturor(&picturor)
	if err != nil || picturor.ID == 0 {
		token = ""
		return
	}
	s.ChangeToken(&picturor)
	token = picturor.Token
	return
}

func (s *picturorService) LoginPicturor(username, password string) (token string, err error) {
	picturor := model.Picturor{Username: username, Password: password}
	err = s.FindPicturor(&picturor)
	if err != nil || picturor.ID == 0 {
		token = ""
		return
	}
	s.ChangeToken(&picturor)
	token = picturor.Token
	return
}

func (s *picturorService) GetPicturorByToken(token string) (picturor model.Picturor, err error) {
	picturor = model.Picturor{Token: token}
	err = s.FindPicturor(&picturor)
	return
}

// public information for Picturor service
func (s *picturorService) CreatePublicPicturor(pubpic *model.PublicPicturor) (err error) {
	err = s.DB.Create(&pubpic).Error
	return
}
func (s *picturorService) UpdatePublicPicturor(pubpic *model.PublicPicturor) (err error) {
	err = s.DB.Save(pubpic).Error
	return err
}
func (s *picturorService) FindPublicPicturor(pubpic *model.PublicPicturor) (err error) {
	err = s.DB.Where(pubpic).First(pubpic).Error
	return err
}

func (s *picturorService) FindPublicPicturors(pubpic *model.PublicPicturor, pubpics *[]model.PublicPicturor) (err error) {
	err = s.DB.Where(pubpic).Find(pubpics).Error
	return err
}
func (s *picturorService) AllPublicPicturors(pubpics *[]model.PublicPicturor) (err error) {
	err = s.DB.Find(&pubpics).Error
	return err
}
func (s *picturorService) DeletePublicPicturor(pubpic *model.PublicPicturor) (err error) {
	err = s.DB.Where(pubpic).Delete(pubpic).Error
	return err
}
