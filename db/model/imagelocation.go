package model

type ImageLocation struct {
	//gorm.Model
	ID               uint           `gorm:"primary" json:"id"`
	FileName         string         `json:"file_name"`
	PublicLocationID uint           `json:"public_location_id"`
	PublicLocation   PublicLocation `json:"public_location" gorm:"foreignKey:PublicLocationID"`
}
