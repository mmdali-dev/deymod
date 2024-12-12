package service

import "github.com/mmdali-dev/deymod/db"

type modelService service

var (
	Model modelService = modelService{"modelService", "CRUD and auth for Models", db.GetDB()}
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
