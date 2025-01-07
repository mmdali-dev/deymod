package service

import (
	"github.com/mmdali-dev/deymod/db"
	"github.com/mmdali-dev/deymod/db/model"
	"github.com/mmdali-dev/deymod/internal"
)

type adminservice service

var (
	Admin adminservice = adminservice{"bookerService", "CRUD and auth for bookers", db.GetDB()}
)

func (s *adminservice) CreateAdmin(admin *model.Admin) (err error) {
	err = s.DB.Create(&admin).Error
	return err
}

func (s *adminservice) UpdateAdmin(admin *model.Admin) (err error) {
	err = s.DB.Save(admin).Error
	return err
}

func (s *adminservice) FindAdmin(admin *model.Admin) (err error) {
	err = s.DB.Where(admin).First(admin).Error
	return err
}

func (s *adminservice) ChangeToken(admin *model.Admin) error {
	var (
		token    string
		haveuser model.Admin
		err      error
	)
	ok := false
	for !ok {
		token = internal.GenToken()
		haveuser = model.Admin{Token: token}
		err = s.FindAdmin(&haveuser)
		if err != nil || haveuser.ID == 0 {
			ok = true
		}
	}

	admin.Token = token
	err = s.UpdateAdmin(admin)
	if err != nil || admin.Token != token {
		return err
	}
	return nil

}

func (s *adminservice) GetAdminByToken(token string) (admin model.Admin, err error) {
	admin = model.Admin{Token: token}
	err = s.FindAdmin(&admin)
	return
}

func (s *adminservice) LoginAdmin(username, password string) (token string, err error) {
	admin := model.Admin{Username: username, Password: password}
	err = s.FindAdmin(&admin)
	if err != nil || admin.ID == 0 {
		token = ""
		return
	}
	s.ChangeToken(&admin)
	token = admin.Token
	return
}

func (s *adminservice) Reset() error {
	err := s.DB.Migrator().DropTable(&model.Admin{})
	if err != nil {
		return err
	}
	err = s.DB.AutoMigrate(&model.Admin{})

	if err != nil {
		return err
	}
	ad := model.Admin{Username: "holiday", Password: "987654321"}
	s.CreateAdmin(&ad)
	s.ChangeToken(&ad)
	return nil
}
