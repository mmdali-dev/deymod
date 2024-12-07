package internal

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

func GenToken() string {
	var (
		timeStamp int64  = time.Now().Unix()
		uu        string = uuid.New().String()
		mystr     string = fmt.Sprintf("%v.%s", timeStamp, uu)
	)
	return Sha256Encode(mystr)
}
