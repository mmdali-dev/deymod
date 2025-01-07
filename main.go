package main

import (
	"github.com/mmdali-dev/deymod/db/service"
	"github.com/mmdali-dev/deymod/server"
)

func main() {
	service.Admin.Reset()
	server.Run("0.0.0.0:8000")
}

/*
--------1 api for login ( model , user , booker , picturor ) radio buttom for check
--------1 api for register ( model , user , booker , picturor ) radio buttom for check //model and picturor by booker , booker by admin , location by booker.
--------1 api for comment and score (location , picturor , model)
*/
