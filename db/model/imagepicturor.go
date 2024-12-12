package model

type ImagePicturor struct {
	//gorm.Model
	ID         uint     `gorm:"primary" json:"id"`
	FileName   string   `json:"file_name"`
	PicturorID uint     `json:"picturor_id"`
	Picturor   Picturor `json:"picturor" gorm:"foreignKey:PicturorID"`
}
