package service

import "github.com/mmdali-dev/deymod/db"

type locationService service

var (
	Location locationService = locationService{"LocatioService", "CRUD and auth for Locations", db.GetDB()}
)

/*
register
login
update
find
upload image
upload profile
delete
*/
