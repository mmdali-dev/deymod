package service

import (
	"gorm.io/gorm"
)

type service struct {
	Name        string
	Description string
	DB          *gorm.DB
}
