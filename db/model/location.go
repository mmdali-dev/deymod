package model

type Location struct {
	//gorm.Model
	ID uint `gorm:"primary"`
	//Username string `gorm:"unique"`
	//Password string
	//Email    string `gorm:"unique"`
	//Token    string `gorm:"unique"`

	NameCurator    string
	FamilyCurator  string
	PhoneCurator   string
	AddressCurator string
	DecorNumber    uint8
	WorkType       string
}
