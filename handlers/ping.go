package handlers

import (
	"encoding/json"
	"fmt"
	"time"
	"iAccounts/cfg"
        "strconv"
)

type ping_Request struct {
	Health string `json:"health"`
	Uptime string `json:"uptime"`
}

func PingResponse() []byte {

	var dur time.Duration
	dur = time.Since(cfg.Starttime)
	var prequest = ping_Request{"OK", strconv.FormatInt(int64(dur), 10)}

	b, err := json.Marshal(prequest)
	if err != nil {
		return nil
	}
	fmt.Println(b)
	return b
}
