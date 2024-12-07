package db

import (
	"fmt"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Init() {

	DB, err = gorm.Open(sqlite.Open("databse.db"), &gorm.Config{
		//Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		fmt.Println("error for setting db")
		panic(err)
	}
	// err = DB.AutoMigrate(&model.User{}, &model.Admin{})
	// if err != nil {
	// 	fmt.Println("error for migration db")
	// 	panic(err)
	// }
}

func GetDB() *gorm.DB {
	return DB
}
