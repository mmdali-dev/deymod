package model

type PublicWomanModel struct {
	//gorm.Model
	ID        uint   `gorm:"primary" json:"id"`
	ModelID   uint   `json:"model_id"`
	Height    uint   `json:"height"`
	Waist     uint   `json:"waist"`
	Pants     uint   `json:"pants"`
	Clothes   uint   `json:"clothes"`
	HairColor string `json:"hair_color"`
	EyeColor  string `json:"eye_color"`
	ShoesSize uint   `json:"shoes_size"`
	//Model     Model             `json:"model" gorm:"foreignKey:ModelID"`
	Comments []WomanComment    `json:"comments" gorm:"foreignKey:WomanID"`
	Images   []ImageWomanModel `json:"images" gorm:"foreignKey:PublicWomanID"`
}
