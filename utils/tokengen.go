package utils

import (
	"math/rand"
	"os/exec"
	"strconv"
	"time"
)

func Get6DigitsRandomNumbers() string {

	rand.Seed(time.Now().Unix())
	rands := rand.Intn(999999-100000) + 100000
	return strconv.Itoa(rands)

}

func GenerateSecureSessionID() string {
	//We may need a better mechanism to generate sessionIDs in the future.
	sessionid, err := exec.Command("uuidgen").Output()
	if err != nil {
		return ""
	}
	return string(sessionid)
}
