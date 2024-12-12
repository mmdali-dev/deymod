package model

type ImageLocation struct {
	//gorm.Model
	ID         uint     `gorm:"primary" json:"id"`
	FileName   string   `json:"file_name"`
	LocationID uint     `json:"location_id"`
	Location   Location `json:"location" gorm:"foreignKey:LocationID"`
}
