package model

type PublicPicturor struct {
	//gorm.Model
	ID               uint     `gorm:"primary" json:"id"`
	PicturorID       uint     `json:"picturor_id"`
	PhotographyStyle string   `json:"photography_style"`
	Video            bool     `json:"video"`
	Quadcopter       bool     `json:"quadcopter"`
	Photoshop        bool     `json:"photoshop"`
	Picturor         Picturor `json:"picturor" gorm:"foreignKey:PicturorID"`
}
