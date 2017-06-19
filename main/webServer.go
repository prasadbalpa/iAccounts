package main

import (
	"fmt"
	"iAccounts/handlers"
	"log"
	"net/http"
	"iAccounts/cfg"
	"time"
)


func main() {

	fmt.Println("Initializing the web server....Please wait....")

	http.HandleFunc("/api/v1/ping", handlers.PingServer)
	http.HandleFunc("/api/v1/login", handlers.Login)
	http.HandleFunc("/api/v1/deliverylog", handlers.DeliveryLog)
	cfg.Starttime = time.Now()
	err := http.ListenAndServeTLS(cfg.HTTPS_SERVER_PORT, cfg.HTTPS_TLS_CERTIFICATE, cfg.HTTPS_TLS_KEY, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
