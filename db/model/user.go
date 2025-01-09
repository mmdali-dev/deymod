package model

type User struct {
	//gorm.Model
	ID       uint   `gorm:"primary" json:"id"`
	Username string `gorm:"unique" json:"username"`
	Password string `json:"-"`
	Email    string `gorm:"unique" json:"email"`
	Token    string `gorm:"unique" json:"-"`
	Name     string `json:"name"`
	Family   string `json:"family"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
	WorkType string `json:"work_type"`
	Page     string `json:"page"`
}
