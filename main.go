package main

import (
	"fmt"

	"github.com/mmdali-dev/deymod/db/service"
)

func main() {

	user, _ := service.User.GetUserByToken("439d7d77ee5b165672107ea89646055e459473169f7278c2978fc0c20c7f8995")
	fmt.Println(user)

}
