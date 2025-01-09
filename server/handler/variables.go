package handler

type authform struct {
	UserType string `form:"user_type"`
	Username string `form:"username"`
	Password string `form:"password"`
}

type locationform struct {
	NameCurator    string `form:"name_curator"`
	FamilyCurator  string `form:"family_curator"`
	PhoneCurator   string `form:"phone_curator"`
	AddressCurator string `form:"address_curator"`
	DecorNumber    uint8  `form:"decor_number"`
	WorkType       string `form:"work_type"`
	City           string `form:"city"`
}

type commentdata struct {
	Text     string `form:"text"`
	Score    uint8  `form:"score"`
	UserType string `form:"user_type"`
	TargetID uint   `form:"tragte_id"`
}

type singleform struct {
	Username string `form:"username"`
	Password string `form:"password"`
}
