package model

type PublicManModel struct {
	//gorm.Model
	ID        uint   `gorm:"primary" json:"id"`
	ModelID   uint   `json:"model_id"`
	Height    uint   `json:"height"`
	Waist     uint   `json:"waist"`
	Chest     uint   `json:"chest"`
	HairColor string `json:"hair_color"`
	EyeColor  string `json:"eye_color"`
	ShoesSize uint   `json:"shoes_size"`
	//Model     Model           `json:"model" gorm:"foreignKey:ModelID"`
	Comments []ManComment    `json:"comments" gorm:"foreignKey:ManID"`
	Images   []ImageManModel `json:"images" gorm:"foreignKey:PublicManID"`
}
