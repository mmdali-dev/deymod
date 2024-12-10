package model

type Location struct {
	//gorm.Model
	ID             uint   `gorm:"primary" json:"id"`
	BookerID       uint   `json:"booker_id"`
	NameCurator    string `json:"name_curator"`
	FamilyCurator  string `json:"family_curator"`
	PhoneCurator   string `json:"phone_curator"`
	AddressCurator string `json:"address_curator"`
	DecorNumber    uint8  `json:"decor_number"`
	WorkType       string `json:"work_type"`
	Booker         Booker `json:"booker" gorm:"foreignKey:BookerID"`
}
