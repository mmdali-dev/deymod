package model

type PicturorComment struct {
	ID               uint   `gorm:"primary" json:"id"`
	Score            uint8  `json:"score"`
	Comment          string `json:"comment"`
	UserID           uint   `json:"user_id"`
	User             User   `json:"user" gorm:"foreignKey:UserID"`
	PublicPicturorID uint   `json:"public_picturor_id"`
}

type LocationComment struct {
	ID               uint   `gorm:"primary" json:"id"`
	Score            uint8  `json:"score"`
	Comment          string `json:"comment"`
	UserID           uint   `json:"user_id"`
	User             User   `json:"user" gorm:"foreignKey:UserID"`
	PublicLocationID uint   `json:"public_location_id"`
}

type ManComment struct {
	ID      uint   `gorm:"primary" json:"id"`
	Score   uint8  `json:"score"`
	Comment string `json:"comment"`
	UserID  uint   `json:"user_id"`
	User    User   `json:"user" gorm:"foreignKey:UserID"`
	ManID   uint   `json:"man_id"`
}

type WomanComment struct {
	ID      uint   `gorm:"primary" json:"id"`
	Score   uint8  `json:"score"`
	Comment string `json:"comment"`
	UserID  uint   `json:"user_id"`
	User    User   `json:"user" gorm:"foreignKey:UserID"`
	WomanID uint   `json:"woman_id"`
}
