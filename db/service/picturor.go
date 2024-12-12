package service

import "github.com/mmdali-dev/deymod/db"

type picturorService service

var (
	Picturor picturorService = picturorService{"picturorService", "CRUD and auth for Picturors", db.GetDB()}
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
