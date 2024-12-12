package service

import (
	"github.com/mmdali-dev/deymod/db"
	"github.com/mmdali-dev/deymod/db/model"
	"github.com/mmdali-dev/deymod/internal"
)

// init service
type bookerService service

var (
	Booker bookerService = bookerService{"bookerService", "CRUD and auth for bookers", db.GetDB()}
)

func (s *bookerService) CreateBooker(booker *model.Booker) (err error) {
	err = s.DB.Create(&booker).Error
	return err
}

func (s *bookerService) DeleteBooker(booker *model.Booker) (err error) {
	err = s.DB.Where(booker).Delete(booker).Error
	return err
}

func (s *bookerService) UpdateBooker(booker *model.Booker) (err error) {
	err = s.DB.Save(booker).Error
	return err
}

func (s *bookerService) FindBooker(booker *model.Booker) (err error) {
	err = s.DB.Where(booker).First(booker).Error
	return err
}

func (s *bookerService) FindBookers(booker *model.Booker, bookers *[]model.Booker) (err error) {
	err = s.DB.Where(booker).Find(bookers).Error
	return err
}

func (s *bookerService) AllBookers(bookers *[]model.Booker) (err error) {
	err = s.DB.Find(&bookers).Error
	return err
}

func (s *bookerService) ChangeToken(booker *model.Booker) error {
	var (
		token    string
		haveuser model.Booker
		err      error
	)
	ok := false
	for !ok {
		token = internal.GenToken()
		haveuser = model.Booker{Token: token}
		err = s.FindBooker(&haveuser)
		if err != nil || haveuser.ID == 0 {
			ok = true
		}
	}

	booker.Token = token
	err = s.UpdateBooker(booker)
	if err != nil || booker.Token != token {
		return err
	}
	return nil

}

func (s *bookerService) RegisterBooker(username, password string) (token string, err error) {
	booker := model.Booker{Username: username, Password: password}
	err = s.CreateBooker(&booker)
	if err != nil || booker.ID == 0 {
		token = ""
		return
	}
	s.ChangeToken(&booker)
	token = booker.Token
	return
}

func (s *bookerService) LoginBooker(username, password string) (token string, err error) {
	booker := model.Booker{Username: username, Password: password}
	err = s.FindBooker(&booker)
	if err != nil || booker.ID == 0 {
		token = ""
		return
	}
	s.ChangeToken(&booker)
	token = booker.Token
	return
}

func (s *bookerService) GetBookerByToken(token string) (booker model.Booker, err error) {
	booker = model.Booker{Token: token}
	err = s.FindBooker(&booker)
	return
}
