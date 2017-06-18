package main

import (
	"fmt"
	"iAccounts/handlers"
	"log"
	"net/http"
)

func main() {

	fmt.Println("Initializing the web server....Please wait....")

	http.HandleFunc("/api/v1/ping", handlers.PingServer)
	http.HandleFunc("/api/v1/login", handlers.Login)
	http.HandleFunc("/api/v1/deliverylog", handlers.DeliveryLog)
	err := http.ListenAndServeTLS(":8445", "/Users/prasadk/go/src/iAccounts/certs/localhost.crt", "/Users/prasadk/go/src/iAccounts/certs/localhost.key", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
