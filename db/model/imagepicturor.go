package model

type ImagePicturor struct {
	//gorm.Model
	ID               uint           `gorm:"primary" json:"id"`
	FileName         string         `json:"file_name"`
	PublicPicturorID uint           `json:"public_picturor_id"`
	PublicPicturor   PublicPicturor `json:"public_picturor" gorm:"foreignKey:PublicPicturorID"`
}
