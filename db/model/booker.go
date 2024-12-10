package model

type Booker struct {
	//gorm.Model
	ID        uint       `gorm:"primary" json:"id"`
	Username  string     `gorm:"unique" json:"username"`
	Password  string     `json:"password"`
	Email     string     `gorm:"unique" json:"email"`
	Token     string     `gorm:"unique" json:"token"`
	Name      string     `json:"name"`
	Family    string     `json:"family"`
	BirthDay  string     `json:"birth_day"`
	Phone     string     `json:"phone"`
	Address   string     `json:"address"`
	WorkType  string     `json:"work_type"`
	Page      string     `json:"page"`
	Models    []Model    `json:"models" gorm:"foreignKey:BookerID"`
	Picturors []Picturor `json:"picturors" gorm:"foreignKey:BookerID"`
	Locations []Location `json:"locations" gorm:"foreignKey:BookerID"`
}
