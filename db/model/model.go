package model

type Model struct {
	//gorm.Model
	ID       uint   `gorm:"primary" json:"id"`
	BookerID uint   `json:"booker_id"`
	Username string `gorm:"unique" json:"username"`
	Password string `json:"password"`
	Email    string `gorm:"unique" json:"email"`
	Token    string `gorm:"unique" json:"token"`
	Name     string `json:"name"`
	Family   string `json:"family"`
	BirthDay string `json:"birth_day"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
	History  string `json:"history"`
	WorkType string `json:"work_type"`
	Gender   bool   `json:"gender"`
	Profile  uint   `json:"profile"`
	Booker   Booker `json:"booker" gorm:"foreignKey:BookerID"`
}
