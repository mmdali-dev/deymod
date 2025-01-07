package model

type Admin struct {
	ID       uint   `gorm:"primary" json:"id"`
	Username string `gorm:"unique" json:"username"`
	Password string `json:"password"`
	Token    string `gorm:"unique" json:"token"`
}
