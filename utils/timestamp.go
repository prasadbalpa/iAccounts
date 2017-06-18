package utils

import (
	"fmt"
	"strconv"
	"time"
)

func GenerateUnixTimeStamp() string {
	str := strconv.FormatInt(time.Now().UnixNano(), 10)
	fmt.Println(str)
	return str
}
