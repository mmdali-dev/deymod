package db

import (
	"fmt"

	//"github.com/glebarez/sqlite"
	"github.com/mmdali-dev/deymod/db/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func init() {

	//DB, err = gorm.Open(sqlite.Open("databse.db"), &gorm.Config{})
	dsn := "root:@tcp(127.0.0.1:3306)/deymod?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("error for setting db")
		panic(err)
	}
	err = DB.AutoMigrate(&model.User{}, &model.Booker{}, &model.Location{}, &model.Model{}, &model.Picturor{}, &model.PublicLocation{}, &model.PublicPicturor{}, &model.PublicManModel{}, &model.PublicWomanModel{})
}

func GetDB() *gorm.DB {
	return DB
}
