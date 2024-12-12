package service

import (
	"github.com/mmdali-dev/deymod/db"
	"github.com/mmdali-dev/deymod/db/model"
	"github.com/mmdali-dev/deymod/internal"
)

// init service
type userService service

var (
	User userService = userService{"UserService", "CRUD and auth for users", db.GetDB()}
)

func (s *userService) CreateUser(user *model.User) (err error) {
	err = s.DB.Create(&user).Error
	return err
}

func (s *userService) DeleteUser(user *model.User) (err error) {
	err = s.DB.Where(user).Delete(user).Error
	return err
}

func (s *userService) UpdateUser(user *model.User) (err error) {
	err = s.DB.Save(user).Error
	return err
}

func (s *userService) FindUser(user *model.User) (err error) {
	err = s.DB.Where(user).First(user).Error
	return err
}

func (s *userService) FindUsers(user *model.User, users *[]model.User) (err error) {
	err = s.DB.Where(user).Find(users).Error
	return err
}

func (s *userService) AllUsers(users *[]model.User) (err error) {
	err = s.DB.Find(&users).Error
	return err
}

func (s *userService) ChangeToken(user *model.User) error {
	var (
		token    string
		haveuser model.User
		err      error
	)
	ok := false
	for !ok {
		token = internal.GenToken()
		haveuser = model.User{Token: token}
		err = s.FindUser(&haveuser)
		if err != nil || haveuser.ID == 0 {
			ok = true
		}
	}

	user.Token = token
	err = s.UpdateUser(user)
	if err != nil || user.Token != token {
		return err
	}
	return nil

}

func (s *userService) RegisterUser(username, password string) (token string, err error) {
	user := model.User{Username: username, Password: password}
	err = s.CreateUser(&user)
	if err != nil || user.ID == 0 {
		token = ""
		return
	}
	s.ChangeToken(&user)
	token = user.Token
	return
}

func (s *userService) LoginUser(username, password string) (token string, err error) {
	user := model.User{Username: username, Password: password}
	err = s.FindUser(&user)
	if err != nil || user.ID == 0 {
		token = ""
		return
	}
	s.ChangeToken(&user)
	token = user.Token
	return
}

func (s *userService) GetUserByToken(token string) (user model.User, err error) {
	user = model.User{Token: token}
	err = s.FindUser(&user)
	return
}
