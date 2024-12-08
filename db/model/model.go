package model

type Model struct {
	//gorm.Model
	ID       uint   `gorm:"primary"`
	Username string `gorm:"unique"`
	Password string
	Email    string `gorm:"unique"`
	Token    string `gorm:"unique"`

	Name     string
	Family   string
	BirthDay string
	Phone    string
	Address  string
	History  string
	WorkType string
}
