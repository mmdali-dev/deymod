package model

type PublicLocation struct {
	//gorm.Model
	ID            uint     `gorm:"primary" json:"id"`
	LocationID    uint     `json:"location_id"`
	DecorNumber   uint     `json:"decor_number"`
	StructureType string   `json:"structure_type"`
	VideoPhoto    bool     `json:"video_photo"`
	Location      Location `json:"location" gorm:"foreignKey:LocationID"`
}
