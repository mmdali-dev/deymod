package model

type ImageManModel struct {
	//gorm.Model
	ID             uint           `gorm:"primary" json:"id"`
	FileName       string         `json:"file_name"`
	PublicManID    uint           `json:"public_man_id"`
	PublicManModel PublicManModel `json:"public_man" gorm:"foreignKey:PublicManID"`
}

type ImageWomanModel struct {
	//gorm.Model
	ID             uint             `gorm:"primary" json:"id"`
	FileName       string           `json:"file_name"`
	PublicWomanID  uint             `json:"public_woman_id"`
	PublicManModel PublicWomanModel `json:"public_woman" gorm:"foreignKey:PublicWomanID"`
}
