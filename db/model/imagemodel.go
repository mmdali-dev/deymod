package model

type ImageModel struct {
	//gorm.Model
	ID       uint   `gorm:"primary" json:"id"`
	FileName string `json:"file_name"`
	ModelID  uint   `json:"model_id"`
	Model    Model  `json:"model" gorm:"foreignKey:ModelID"`
}
